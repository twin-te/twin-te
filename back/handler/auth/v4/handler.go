package authv4

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/appenv"
	"github.com/twin-te/twin-te/back/apperr"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	autherr "github.com/twin-te/twin-te/back/module/auth/err"
)

func (h *impl) handleOAuth2(c echo.Context) error {
	state := generateState()

	url, ok := getOAuth2AuthCodeURL(c.Param("provider"), state)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid provider")
	}

	redirectURL := getRedirectURLFromQuery(c)
	if redirectURL != appenv.AUTH_DEFAULT_REDIRECT_URL {
		setAuthRedirectURLInCookie(c, redirectURL)
	}

	// Discard the connect flow which may have been started before, so that the callback is handled as login.
	clearAuthConnectFromCookie(c)
	setAuthStateInCookie(c, state)

	return c.Redirect(http.StatusFound, url)
}

func (h *impl) handleOAuth2Connect(c echo.Context) error {
	if _, err := h.accessController.Authenticate(c.Request().Context()); err != nil {
		return c.Redirect(http.StatusFound, getConnectErrorRedirectURL(oauth2ErrorCodeUnauthenticated))
	}

	state := generateState()

	url, ok := getOAuth2AuthCodeURL(c.Param("provider"), state)
	if !ok {
		return c.Redirect(http.StatusFound, getConnectErrorRedirectURL(oauth2ErrorCodeInvalidProvider))
	}

	setAuthConnectInCookie(c)
	setAuthStateInCookie(c, state)

	return c.Redirect(http.StatusFound, url)
}

func getOAuth2AuthCodeURL(provider string, state string) (url string, ok bool) {
	switch provider {
	case "google":
		return googleOAuth2Config.AuthCodeURL(state), true
	case "apple":
		return appleOAuth2Config.AuthCodeURL(state), true
	case "twitter":
		return twitterOAuth2Config.AuthCodeURL(state, s256ChallengeOption), true
	default:
		return "", false
	}
}

// handleOAuth2Callback handles the callback of both login and connect.
// It always redirects to the front end, even if an error occurs.
func (h *impl) handleOAuth2Callback(c echo.Context) error {
	isConnecting := isConnectingFromCookie(c)
	redirectURL := getRedirectURLFromCookie(c)

	clearAuthStateFromCookie(c)
	clearAuthRedirectURLFromCookie(c)
	clearAuthConnectFromCookie(c)

	if isConnecting {
		return c.Redirect(http.StatusFound, h.connectWithOAuth2Callback(c))
	}
	return c.Redirect(http.StatusFound, h.loginWithOAuth2Callback(c, redirectURL))
}

// connectWithOAuth2Callback adds the authentication to the logged-in user, and returns the url to redirect to.
func (h *impl) connectWithOAuth2Callback(c echo.Context) string {
	userAuthentication, err := getUserAuthenticationFromOAuth2Callback(c)
	if err != nil {
		return getConnectErrorRedirectURL(getOAuth2ErrorCode(err))
	}

	err = h.authUseCase.AddUserAuthentication(c.Request().Context(), userAuthentication)
	if err == nil {
		return getConnectResultRedirectURL(oauth2ResultConnected)
	}

	if apperr.Is(err, autherr.CodeUserAuthenticationAlreadyExists) {
		if user, err := h.authUseCase.GetMe(c.Request().Context()); err == nil && lo.Contains(user.Authentications, userAuthentication) {
			return getConnectResultRedirectURL(oauth2ResultAlreadyConnected)
		}
	}

	return getConnectErrorRedirectURL(getOAuth2ErrorCode(err))
}

// loginWithOAuth2Callback signs up or logs in, and returns the url to redirect to.
// Unlike the id token endpoints, the authentication is not added to the logged-in user.
func (h *impl) loginWithOAuth2Callback(c echo.Context, redirectURL string) string {
	userAuthentication, err := getUserAuthenticationFromOAuth2Callback(c)
	if err != nil {
		return getLoginErrorRedirectURL(getOAuth2ErrorCode(err))
	}

	session, err := h.authUseCase.SignUpOrLogin(c.Request().Context(), userAuthentication)
	if err != nil {
		return getLoginErrorRedirectURL(getOAuth2ErrorCode(err))
	}

	setSessionInCookie(c, session)

	return redirectURL
}

func getUserAuthenticationFromOAuth2Callback(c echo.Context) (userAuthentication authdomain.UserAuthentication, err error) {
	if err = validateState(c); err != nil {
		return userAuthentication, newOAuth2Error(oauth2ErrorCodeInvalidState, err)
	}

	// The provider returns error instead of code, e.g. when the user cancels the authorization.
	if providerError := c.QueryParam("error"); providerError != "" {
		return userAuthentication, newOAuth2Error(
			getOAuth2ErrorCodeFromProviderError(providerError),
			fmt.Errorf("the provider returned error, %s", providerError),
		)
	}

	code := c.QueryParam("code")
	if code == "" {
		return userAuthentication, echo.NewHTTPError(http.StatusBadRequest, "authorization code is required")
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
		err = newOAuth2Error(oauth2ErrorCodeInvalidProvider, errors.New("invalid provider"))
	}
	if err != nil {
		return
	}

	return authdomain.NewUserAuthentication(provider, socialID), nil
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
	Token       string `query:"token" form:"token" json:"token"`
	RedirectURL string `query:"redirect_url" form:"redirect_url" json:"redirect_url"`
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

func (h *impl) handleLogout(c echo.Context) error {
	if err := h.authUseCase.Logout(c.Request().Context()); err != nil {
		return err
	}

	clearSessionFromCookie(c)

	return c.Redirect(http.StatusFound, getRedirectURLFromQuery(c))
}
