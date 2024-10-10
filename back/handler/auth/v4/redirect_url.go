package authv4

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/appenv"
)

// getRedirectURLFromQuery returns redirect url retrieved from query.
// If it is not found or it is invalid, default redirect url will be returned.
func getRedirectURLFromQuery(c echo.Context) string {
	redirectURL := c.QueryParam("redirect_url")
	return lo.Ternary(isValidRedirectURL(redirectURL), redirectURL, appenv.AUTH_DEFAULT_REDIRECT_URL)
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
