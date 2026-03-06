package calendarv1beta

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/appctx"
	"github.com/twin-te/twin-te/back/appenv"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func newEchoWithActorOrIcalToken(accessController authmodule.AccessController, calendar calendarmodule.UseCase) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			if tokenStr := c.QueryParam("token"); tokenStr != "" {
				if c.Path() != "/timetable.ics" {
					return echo.NewHTTPError(http.StatusForbidden, "token auth is not allowed for this endpoint")
				}
				id, err := idtype.ParseIcalSubscriptionID(tokenStr)
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
				}
				userID, ok, err := calendar.ResolveUserIDByIcalSubscriptionID(ctx, id)
				if err != nil {
					return err
				}
				if !ok {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
				}
				ctx = appctx.SetActor(ctx, authdomain.NewAuthNUser(userID))
				c.SetRequest(c.Request().WithContext(ctx))
				return next(c)
			}
			sessionID, ok := extractSessionIDFromEchoContext(c)
			ctx, err := accessController.WithActor(c.Request().Context(), mo.TupleToOption(sessionID, ok))
			if err != nil {
				return err
			}
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func extractSessionIDFromEchoContext(c echo.Context) (sessionID idtype.SessionID, ok bool) {
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
