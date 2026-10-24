package authv4

import (
	"log"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/appenv"
)

// getRedirectURLFromQuery returns redirect url retrieved from query.
// If it is not found or it is invalid, default redirect url will be returned.
func getRedirectURLFromQuery(c echo.Context) string {
	return getRedirectURL(c.QueryParam("redirect_url"))
}

// getRedirectURLFromCookie returns redirect url retrieved from cookie.
// If it is not found or it is invalid, default redirect url will be returned.
func getRedirectURLFromCookie(c echo.Context) string {
	if cookie, err := c.Cookie(appenv.COOKIE_AUTH_REDIRECT_URL_NAME); err == nil {
		if redirectURL := cookie.Value; isValidRedirectURL(redirectURL) {
			return redirectURL
		}
	}
	return appenv.AUTH_DEFAULT_REDIRECT_URL
}

func isValidRedirectURL(redirectURL string) bool {
	return lo.Contains(appenv.AUTH_ALLOWED_REDIRECT_URLS, redirectURL)
}

func getRedirectURL(redirectURL string) string {
	return lo.Ternary(isValidRedirectURL(redirectURL), redirectURL, appenv.AUTH_DEFAULT_REDIRECT_URL)
}

// getConnectResultRedirectURL returns the url to redirect to after connecting an authentication successfully.
func getConnectResultRedirectURL(result oauth2Result) string {
	return addQueryToURL(appenv.AUTH_CONNECT_REDIRECT_URL, "connect_result", string(result))
}

// getConnectErrorRedirectURL returns the url to redirect to after failing to connect an authentication.
func getConnectErrorRedirectURL(code oauth2ErrorCode) string {
	return addQueryToURL(appenv.AUTH_CONNECT_REDIRECT_URL, "connect_error", string(code))
}

// getLoginErrorRedirectURL returns the url to redirect to after failing to log in.
func getLoginErrorRedirectURL(code oauth2ErrorCode) string {
	return addQueryToURL(appenv.AUTH_LOGIN_ERROR_REDIRECT_URL, "login_error", string(code))
}

func addQueryToURL(rawURL, key, value string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Printf("failed to parse url %s, %+v", rawURL, err)
		return rawURL
	}
	query := u.Query()
	query.Set(key, value)
	u.RawQuery = query.Encode()
	return u.String()
}
