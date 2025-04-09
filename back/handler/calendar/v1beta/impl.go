package calendarv1beta

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/twin-te/twin-te/back/handler/common/middleware"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
)

// impl handles the requests beginning with the following paths.
//   - "/timetable.ics"
type impl struct {
	schoolcalendar schoolcalendarmodule.UseCase
	timetable      timetablemodule.UseCase
}

func New(
	accessController authmodule.AccessController,

	schoolcalendar schoolcalendarmodule.UseCase,
	timetable timetablemodule.UseCase,
) *echo.Echo {
	h := &impl{
		schoolcalendar: schoolcalendar,
		timetable:      timetable,
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
