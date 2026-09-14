package calendardomain_test

import (
	"testing"
	"time"

	"cloud.google.com/go/civil"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
)

func newModule() *calendardomain.SchoolCalendarModule {
	return &calendardomain.SchoolCalendarModule{
		Module:     schoolcalendardomain.ModuleSpringA,
		Start:      civil.Date{Year: 2026, Month: 4, Day: 1},
		End:        civil.Date{Year: 2026, Month: 5, Day: 31},
		Exceptions: map[time.Weekday][]civil.Date{},
		Additions:  map[time.Weekday][]civil.Date{},
	}
}

func TestSchoolCalendarModule_AddException(t *testing.T) {
	m := newModule()
	date := civil.Date{Year: 2026, Month: 4, Day: 6}

	m.AddException(time.Monday, date)
	if len(m.Exceptions[time.Monday]) != 1 {
		t.Fatalf("len(Exceptions[Monday]) = %v, want 1", len(m.Exceptions[time.Monday]))
	}
	if m.Exceptions[time.Monday][0] != date {
		t.Errorf("Exceptions[Monday][0] = %v, want %v", m.Exceptions[time.Monday][0], date)
	}

	// Adding the same date again should be a no-op.
	m.AddException(time.Monday, date)
	if len(m.Exceptions[time.Monday]) != 1 {
		t.Errorf("duplicate not filtered: len = %v, want 1", len(m.Exceptions[time.Monday]))
	}

	// Adding a different date should append.
	date2 := civil.Date{Year: 2026, Month: 4, Day: 13}
	m.AddException(time.Monday, date2)
	if len(m.Exceptions[time.Monday]) != 2 {
		t.Errorf("len(Exceptions[Monday]) = %v, want 2", len(m.Exceptions[time.Monday]))
	}
}

func TestSchoolCalendarModule_AddAddition(t *testing.T) {
	m := newModule()
	date := civil.Date{Year: 2026, Month: 4, Day: 7}

	m.AddAddition(time.Tuesday, date)
	if len(m.Additions[time.Tuesday]) != 1 {
		t.Fatalf("len(Additions[Tuesday]) = %v, want 1", len(m.Additions[time.Tuesday]))
	}
	if m.Additions[time.Tuesday][0] != date {
		t.Errorf("Additions[Tuesday][0] = %v, want %v", m.Additions[time.Tuesday][0], date)
	}

	// Adding the same date again should be a no-op.
	m.AddAddition(time.Tuesday, date)
	if len(m.Additions[time.Tuesday]) != 1 {
		t.Errorf("duplicate not filtered: len = %v, want 1", len(m.Additions[time.Tuesday]))
	}

	// Adding a different date should append.
	date2 := civil.Date{Year: 2026, Month: 4, Day: 14}
	m.AddAddition(time.Tuesday, date2)
	if len(m.Additions[time.Tuesday]) != 2 {
		t.Errorf("len(Additions[Tuesday]) = %v, want 2", len(m.Additions[time.Tuesday]))
	}
}
