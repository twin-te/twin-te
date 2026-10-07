package calendardomain

import (
	"slices"
	"time"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
)

type SchoolCalendarModule struct {
	Module     schoolcalendardomain.Module
	Start      civil.Date
	End        civil.Date
	Exceptions map[time.Weekday][]civil.Date
	Additions  map[time.Weekday][]civil.Date
}

func (m *SchoolCalendarModule) AddException(weekday time.Weekday, date civil.Date) {
	if slices.Contains(m.Exceptions[weekday], date) {
		return
	}
	m.Exceptions[weekday] = append(m.Exceptions[weekday], date)
}

func (m *SchoolCalendarModule) AddAddition(weekday time.Weekday, date civil.Date) {
	if slices.Contains(m.Additions[weekday], date) {
		return
	}
	m.Additions[weekday] = append(m.Additions[weekday], date)
}
