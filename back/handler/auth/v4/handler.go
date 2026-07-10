package authv4

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/twin-te/twin-te/back/appenv"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func (h *impl) handleOAuth2(c echo.Context) error {
	state := generateState()

	var url string
	switch c.Param("provider") {
	case "google":
		url = googleOAuth2Config.AuthCodeURL(state)
	case "apple":
		url = appleOAuth2Config.AuthCodeURL(state)
	case "twitter":
		url = twitterOAuth2Config.AuthCodeURL(state, s256ChallengeOption)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "invalid provider")
	}

	redirectURL := getRedirectURLFromQuery(c)
	if redirectURL != appenv.AUTH_DEFAULT_REDIRECT_URL {
		setAuthRedirectURLInCookie(c, redirectURL)
	}

	setAuthStateInCookie(c, state)

	return c.Redirect(http.StatusFound, url)
}

func (h *impl) handleOAuth2Callback(c echo.Context) (err error) {
	defer func() {
		clearAuthStateFromCookie(c)
		clearAuthRedirectURLFromCookie(c)

		if err != nil {
			return
		}

		c.Redirect(http.StatusFound, getRedirectURLFromCookie(c))
	}()

	if err := validateState(c); err != nil {
		return err
	}

	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "authorization code is required")
	}

	var (
		provider authdomain.Provider
		socialID authdomain.SocialID
	)

	switch c.Param("provider") {
	case "google":
		provider = authdomain.ProviderGoogle
		socialID, err = getGoogleSocialID(c.Request().Context(), code)
	case "apple":
		provider = authdomain.ProviderApple
		socialID, err = getAppleSocialID(c.Request().Context(), code)
	case "twitter":
		provider = authdomain.ProviderTwitter
		socialID, err = getTwitterSocialID(c.Request().Context(), code)
	default:
		err = echo.NewHTTPError(http.StatusBadRequest, "invalid provider")
	}
	if err != nil {
		return
	}

	userAuthentication := authdomain.NewUserAuthentication(provider, socialID)

	if _, err := h.accessController.Authenticate(c.Request().Context()); err == nil {
		err = h.authUseCase.AddUserAuthentication(c.Request().Context(), userAuthentication)
		if err != nil {
			return err
		}
	}

	session, err := h.authUseCase.SignUpOrLogin(c.Request().Context(), userAuthentication)
	if err != nil {
		return
	}

	setSessionInCookie(c, session)

	return nil
}

func (h *impl) handleIDTokenGoogle(c echo.Context) error {
	credential, err := getIDTokenCredential(c)
	if err != nil {
		return err
	}

	expectedNonce := ""
	if credential.ChallengeID == "" {
		markDeprecatedAuthRequest(c)
	} else {
		challenge, err := h.consumeAuthChallenge(c, authdomain.ProviderGoogle, credential.ChallengeID)
		if err != nil {
			return err
		}
		expectedNonce = challenge.Nonce
	}

	idToken := credential.Token
	socialID, err := verifyGoogleIDToken(c.Request().Context(), idToken, expectedNonce)
	if err != nil {
		return err
	}

	_, err = h.signUpOrLoginWithAuthentication(
		c,
		authdomain.NewUserAuthentication(authdomain.ProviderGoogle, socialID),
	)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, getRedirectURL(credential.RedirectURL))
}

func (h *impl) handleIDTokenApple(c echo.Context) error {
	credential, err := getIDTokenCredential(c)
	if err != nil {
		return err
	}

	expectedNonce := ""
	if credential.ChallengeID == "" {
		markDeprecatedAuthRequest(c)
	} else {
		challenge, err := h.consumeAuthChallenge(c, authdomain.ProviderApple, credential.ChallengeID)
		if err != nil {
			return err
		}
		expectedNonce = challenge.Nonce
	}

	if credential.AuthorizationCode == "" && credential.ChallengeID != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "authorization_code is required")
	}
	identity, err := verifyAppleIDToken(c.Request().Context(), credential.Token, expectedNonce)
	if err != nil {
		return err
	}
	refreshToken := ""
	// Keep the challenge-less endpoint compatible with already released clients.
	// Authorization-code exchange is required only for the new challenge-bound flow.
	if credential.ChallengeID != "" {
		refreshToken, err = exchangeAppleAuthorizationCode(
			c.Request().Context(),
			credential.AuthorizationCode,
			identity,
			expectedNonce,
		)
		if err != nil {
			return err
		}
	}

	session, err := h.signUpOrLoginWithAuthentication(
		c,
		authdomain.NewUserAuthentication(authdomain.ProviderApple, identity.SocialID),
	)
	if err != nil {
		return err
	}
	if refreshToken != "" {
		if err := h.authUseCase.SaveAppleCredential(c.Request().Context(), &authdomain.AppleCredential{
			UserID:       session.UserID,
			ClientID:     identity.Audience,
			RefreshToken: refreshToken,
		}); err != nil {
			return err
		}
	}
	return c.Redirect(http.StatusFound, getRedirectURL(credential.RedirectURL))
}

func (h *impl) signUpOrLoginWithAuthentication(
	c echo.Context,
	userAuthentication authdomain.UserAuthentication,
) (*authdomain.Session, error) {
	if _, err := h.accessController.Authenticate(c.Request().Context()); err == nil {
		err = h.authUseCase.AddUserAuthentication(c.Request().Context(), userAuthentication)
		if err != nil {
			return nil, err
		}
	}

	session, err := h.authUseCase.SignUpOrLogin(c.Request().Context(), userAuthentication)
	if err != nil {
		return nil, err
	}

	setSessionInCookie(c, session)

	return session, nil
}

type idTokenCredential struct {
	Token             string `query:"token" form:"token" json:"token"`
	RedirectURL       string `query:"redirect_url" form:"redirect_url" json:"redirect_url"`
	ChallengeID       string `query:"challenge_id" form:"challenge_id" json:"challenge_id"`
	AuthorizationCode string `query:"authorization_code" form:"authorization_code" json:"authorization_code"`
}

func getIDTokenCredential(c echo.Context) (*idTokenCredential, error) {
	credential := &idTokenCredential{}
	if err := c.Bind(credential); err != nil {
		return nil, err
	}
	if credential.Token == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "token is required")
	}
	return credential, nil
}

func (h *impl) handleCreateAuthChallenge(c echo.Context) error {
	provider, err := providerFromPath(c.Param("provider"))
	if err != nil || provider == authdomain.ProviderTwitter {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid authentication provider")
	}
	challenge, err := h.authUseCase.CreateAuthChallenge(c.Request().Context(), provider)
	if err != nil {
		return err
	}
	response := url.Values{
		"challenge_id": []string{challenge.ID},
		"nonce":        []string{challenge.Nonce},
	}.Encode()
	return c.String(http.StatusOK, response)
}

func (h *impl) consumeAuthChallenge(
	c echo.Context,
	provider authdomain.Provider,
	challengeID string,
) (*authdomain.AuthChallenge, error) {
	if challengeID == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "challenge_id is required")
	}
	challenge, err := h.authUseCase.ConsumeAuthChallenge(c.Request().Context(), challengeID, provider)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid authentication challenge").SetInternal(err)
	}
	return challenge, nil
}

func providerFromPath(value string) (authdomain.Provider, error) {
	switch value {
	case "google":
		return authdomain.ProviderGoogle, nil
	case "apple":
		return authdomain.ProviderApple, nil
	default:
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid authentication provider")
	}
}

func markDeprecatedAuthRequest(c echo.Context) {
	c.Response().Header().Set("Deprecation", "true")
}

func (h *impl) handleLogout(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		markDeprecatedAuthRequest(c)
	}
	if err := h.authUseCase.Logout(c.Request().Context()); err != nil {
		return err
	}

	clearSessionFromCookie(c)

	return c.Redirect(http.StatusFound, getRedirectURLFromQuery(c))
}
