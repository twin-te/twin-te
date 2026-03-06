package calendardomain

import (
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
	for _, d := range m.Exceptions[weekday] {
		if d == date {
			return
		}
	}
	m.Exceptions[weekday] = append(m.Exceptions[weekday], date)
}
