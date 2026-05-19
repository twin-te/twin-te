package authv4

import (
	"net/http"

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

	idToken := credential.Token
	socialID, err := verifyGoogleIDToken(c.Request().Context(), idToken)
	if err != nil {
		return err
	}

	return h.signUpOrLoginWithAuthentication(
		c,
		authdomain.NewUserAuthentication(authdomain.ProviderGoogle, socialID),
		credential.RedirectURL,
	)
}

func (h *impl) handleIDTokenApple(c echo.Context) error {
	credential, err := getIDTokenCredential(c)
	if err != nil {
		return err
	}

	socialID, err := verifyAppleIDToken(c.Request().Context(), credential.Token)
	if err != nil {
		return err
	}

	return h.signUpOrLoginWithAuthentication(
		c,
		authdomain.NewUserAuthentication(authdomain.ProviderApple, socialID),
		credential.RedirectURL,
	)
}

func (h *impl) signUpOrLoginWithAuthentication(
	c echo.Context,
	userAuthentication authdomain.UserAuthentication,
	redirectURL string,
) error {
	if _, err := h.accessController.Authenticate(c.Request().Context()); err == nil {
		err = h.authUseCase.AddUserAuthentication(c.Request().Context(), userAuthentication)
		if err != nil {
			return err
		}
	}

	session, err := h.authUseCase.SignUpOrLogin(c.Request().Context(), userAuthentication)
	if err != nil {
		return err
	}

	setSessionInCookie(c, session)

	return c.Redirect(http.StatusFound, getRedirectURL(redirectURL))
}

type idTokenCredential struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

func getIDTokenCredential(c echo.Context) (*idTokenCredential, error) {
	if c.Request().Method == http.MethodGet {
		credential := &idTokenCredential{
			Token:       c.QueryParam("token"),
			RedirectURL: c.QueryParam("redirect_url"),
		}
		if credential.Token == "" {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "token is required")
		}
		return credential, nil
	}

	credential := &idTokenCredential{
		Token:       c.FormValue("token"),
		RedirectURL: c.FormValue("redirect_url"),
	}
	if credential.Token == "" || credential.RedirectURL == "" {
		jsonCredential := &idTokenCredential{}
		if err := c.Bind(jsonCredential); err != nil && credential.Token == "" {
			return nil, err
		}
		if credential.Token == "" {
			credential.Token = jsonCredential.Token
		}
		if credential.RedirectURL == "" {
			credential.RedirectURL = jsonCredential.RedirectURL
		}
	}
	if credential.Token == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "token is required")
	}
	return credential, nil
}

func (h *impl) handleLogout(c echo.Context) error {
	if err := h.authUseCase.Logout(c.Request().Context()); err != nil {
		return err
	}

	clearSessionFromCookie(c)

	return c.Redirect(http.StatusFound, getRedirectURLFromQuery(c))
}
