package calendarv1

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/twin-te/twin-te/back/handler/common/middleware"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
)

const PathPrefix = "/calendar/v1"

// impl handles the requests beginning with the following paths.
//   - "/timetable.ics"
type impl struct {
	calendar calendarmodule.UseCase
}

func New(
	accessController authmodule.AccessController,
	calendar calendarmodule.UseCase,
) *echo.Echo {
	h := &impl{
		calendar: calendar,
	}

	e := echo.New()

	e.Use(
		echomiddleware.Recover(),
		echomiddleware.Logger(),
		middleware.NewEchoErrorHandler(),
		middleware.NewEchoWithActor(accessController),
	)

	e.GET("/timetable.ics", h.ICSHandler)

	return e
}
