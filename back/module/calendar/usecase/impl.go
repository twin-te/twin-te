package calendarusecase

import (
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
)

var _ calendarmodule.UseCase = (*impl)(nil)

type impl struct {
	a              authmodule.AccessController
	r              calendarport.Repository
	schoolcalendar schoolcalendarmodule.UseCase
	timetable      timetablemodule.UseCase
}

func New(
	a authmodule.AccessController,
	r calendarport.Repository,
	schoolcalendar schoolcalendarmodule.UseCase,
	timetable timetablemodule.UseCase,
) *impl {
	return &impl{a: a, r: r, schoolcalendar: schoolcalendar, timetable: timetable}
}
