package calendardomain

import (
	"testing"
	"time"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestGetSchedulesNighttimePeriods(t *testing.T) {
	modules := []SchoolCalendarModule{
		{
			Module:     schoolcalendardomain.ModuleSpringA,
			Start:      civil.Date{Year: 2026, Month: time.April, Day: 6},
			End:        civil.Date{Year: 2026, Month: time.May, Day: 25},
			Exceptions: map[time.Weekday][]civil.Date{},
			Additions:  map[time.Weekday][]civil.Date{},
		},
	}
	date := civil.Date{Year: 2026, Month: time.April, Day: 6}

	tests := []struct {
		name    string
		periods []timetabledomain.Period
		start   civil.Time
		end     civil.Time
	}{
		{name: "period 7", periods: []timetabledomain.Period{7}, start: civil.Time{Hour: 18, Minute: 0}, end: civil.Time{Hour: 19, Minute: 15}},
		{name: "period 8", periods: []timetabledomain.Period{8}, start: civil.Time{Hour: 19, Minute: 20}, end: civil.Time{Hour: 20, Minute: 35}},
		{name: "periods 6 and 7", periods: []timetabledomain.Period{6, 7}, start: civil.Time{Hour: 16, Minute: 45}, end: civil.Time{Hour: 19, Minute: 15}},
		{name: "periods 7 and 8", periods: []timetabledomain.Period{7, 8}, start: civil.Time{Hour: 18, Minute: 0}, end: civil.Time{Hour: 20, Minute: 35}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := make([]timetabledomain.Schedule, len(tt.periods))
			for i, p := range tt.periods {
				ss[i] = timetabledomain.Schedule{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: p}
			}

			got := GetSchedules(modules, ss)
			if len(got) != 1 {
				t.Fatalf("expected 1 schedule, got %d", len(got))
			}

			wantStart := civil.DateTime{Date: date, Time: tt.start}
			wantEnd := civil.DateTime{Date: date, Time: tt.end}
			if got[0].StartTime != wantStart || got[0].EndTime != wantEnd {
				t.Errorf("expected %v - %v, got %v - %v", wantStart, wantEnd, got[0].StartTime, got[0].EndTime)
			}
		})
	}
}
