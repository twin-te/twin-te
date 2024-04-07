package authv3

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/twin-te/twin-te/back/appenv"
)

func generateState() string {
	return uuid.NewString()
}

func validateState(c echo.Context) error {
	cookie, err := c.Cookie(appenv.COOKIE_AUTH_STATE_NAME)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "state is not found in cookie")
	}

	stateInCookie, stateInQuery := cookie.Value, c.QueryParam("state")
	if stateInCookie == "" || stateInQuery == "" || stateInCookie != stateInQuery {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid state")
	}

	return nil
}
