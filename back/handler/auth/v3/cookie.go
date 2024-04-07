package authv3

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/twin-te/twin-te/back/appenv"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func setAuthStateInCookie(c echo.Context, state string) {
	c.SetCookie(&http.Cookie{
		Name:     appenv.COOKIE_AUTH_STATE_NAME,
		Value:    state,
		Path:     "/",
		MaxAge:   3 * 60,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearAuthStateFromCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:   appenv.COOKIE_AUTH_STATE_NAME,
		Path:   "/",
		MaxAge: -1,
	})
}

func setAuthRedirectURLInCookie(c echo.Context, redirectURL string) {
	c.SetCookie(&http.Cookie{
		Name:     appenv.COOKIE_AUTH_REDIRECT_URL_NAME,
		Value:    redirectURL,
		Path:     "/",
		MaxAge:   3 * 60,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearAuthRedirectURLFromCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:   appenv.COOKIE_AUTH_REDIRECT_URL_NAME,
		Path:   "/",
		MaxAge: -1,
	})
}

func setSessionInCookie(c echo.Context, session *authdomain.Session) {
	c.SetCookie(&http.Cookie{
		Name:     appenv.COOKIE_SESSION_NAME,
		Value:    session.ID.String(),
		Path:     "/",
		Expires:  session.ExpiredAt,
		Secure:   appenv.COOKIE_SECURE,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearSessionFromCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:   appenv.COOKIE_SESSION_NAME,
		Path:   "/",
		MaxAge: -1,
	})
}
