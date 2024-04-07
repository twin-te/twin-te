package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/appenv"
	"github.com/twin-te/twin-te/back/apperr"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func NewEchoErrorHandler() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err == nil {
				return nil
			}

			if aerr, ok := apperr.As(err); ok {
				if httpStatusCode, ok := AppErrorCodeToHttpStatusCode[aerr.Code]; ok {
					err = echo.NewHTTPError(httpStatusCode, aerr.Message)
				}
			}

			httpError, ok := err.(*echo.HTTPError)
			if !ok {
				httpError = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if httpError.Code >= 500 {
				log.Printf("unexpected error: %+v", httpError)
			}

			return httpError
		}
	}
}

func NewEchoWithActor(accessController authmodule.AccessController) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionID, ok := getSessionIDFromEchoContext(c)
			ctx, err := accessController.WithActor(c.Request().Context(), lo.Ternary(ok, &sessionID, nil))
			if err != nil {
				return err
			}

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func getSessionIDFromEchoContext(c echo.Context) (sessionID idtype.SessionID, ok bool) {
	cookie, err := c.Cookie(appenv.COOKIE_SESSION_NAME)
	if err != nil {
		return
	}

	sessionID, err = idtype.ParseSessionID(cookie.Value)
	if err != nil {
		return
	}

	return sessionID, true
}
