package calendarusecase

import (
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

var _ calendarmodule.UseCase = (*impl)(nil)

type impl struct {
	a              authmodule.AccessController
	r              calendarport.Repository
	schoolcalendar schoolcalendarmodule.UseCase
	timetable      timetableport.Query
}

func New(
	a authmodule.AccessController,
	r calendarport.Repository,
	schoolcalendar schoolcalendarmodule.UseCase,
	timetable timetableport.Query,
) *impl {
	return &impl{a: a, r: r, schoolcalendar: schoolcalendar, timetable: timetable}
}
