package calendardomain_test

import (
	"testing"
	"time"

	"cloud.google.com/go/civil"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func springAModule() calendardomain.SchoolCalendarModule {
	return calendardomain.SchoolCalendarModule{
		Module:     schoolcalendardomain.ModuleSpringA,
		Start:      civil.Date{Year: 2026, Month: 4, Day: 1},
		End:        civil.Date{Year: 2026, Month: 5, Day: 31},
		Exceptions: map[time.Weekday][]civil.Date{},
		Additions:  map[time.Weekday][]civil.Date{},
	}
}

func springBModule() calendardomain.SchoolCalendarModule {
	return calendardomain.SchoolCalendarModule{
		Module:     schoolcalendardomain.ModuleSpringB,
		Start:      civil.Date{Year: 2026, Month: 6, Day: 1},
		End:        civil.Date{Year: 2026, Month: 7, Day: 31},
		Exceptions: map[time.Weekday][]civil.Date{},
		Additions:  map[time.Weekday][]civil.Date{},
	}
}

func TestGetSchedules_Basic(t *testing.T) {
	modules := []calendardomain.SchoolCalendarModule{springAModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
	}

	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 1 {
		t.Fatalf("len(got) = %v, want 1", len(got))
	}
	s := got[0]
	if s.Weekday != time.Monday {
		t.Errorf("Weekday = %v, want Monday", s.Weekday)
	}
	if s.Location != "Room A" {
		t.Errorf("Location = %v, want Room A", s.Location)
	}
	// first Monday on/after 2026-04-01
	wantDate := civil.Date{Year: 2026, Month: 4, Day: 6}
	if s.StartTime.Date != wantDate {
		t.Errorf("StartTime.Date = %v, want %v", s.StartTime.Date, wantDate)
	}
	if s.StartTime.Time != (civil.Time{Hour: 8, Minute: 40}) {
		t.Errorf("StartTime.Time = %v", s.StartTime.Time)
	}
	if s.EndTime.Time != (civil.Time{Hour: 9, Minute: 55}) {
		t.Errorf("EndTime.Time = %v", s.EndTime.Time)
	}
	wantUntil := civil.DateTime{Date: civil.Date{Year: 2026, Month: 5, Day: 31}, Time: civil.Time{Hour: 23, Minute: 59}}
	if s.Until != wantUntil {
		t.Errorf("Until = %v, want %v", s.Until, wantUntil)
	}
}

func TestGetSchedules_SkipsSpecialDay(t *testing.T) {
	modules := []calendardomain.SchoolCalendarModule{springAModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayIntensive, Period: 1, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 0 {
		t.Errorf("len(got) = %v, want 0", len(got))
	}
}

func TestGetSchedules_SkipsOutOfRangePeriod(t *testing.T) {
	modules := []calendardomain.SchoolCalendarModule{springAModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 0, Locations: "Room A"},
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 7, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 0 {
		t.Errorf("len(got) = %v, want 0", len(got))
	}
}

func TestGetSchedules_MergesConsecutivePeriods(t *testing.T) {
	modules := []calendardomain.SchoolCalendarModule{springAModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 2, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 1 {
		t.Fatalf("len(got) = %v, want 1", len(got))
	}
	s := got[0]
	if s.StartTime.Time != (civil.Time{Hour: 8, Minute: 40}) {
		t.Errorf("StartTime.Time = %v, want period1 start", s.StartTime.Time)
	}
	if s.EndTime.Time != (civil.Time{Hour: 11, Minute: 25}) {
		t.Errorf("EndTime.Time = %v, want period2 end", s.EndTime.Time)
	}
}

func TestGetSchedules_DoesNotMergeDifferentLocation(t *testing.T) {
	modules := []calendardomain.SchoolCalendarModule{springAModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 2, Locations: "Room B"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 2 {
		t.Fatalf("len(got) = %v, want 2", len(got))
	}
}

func TestGetSchedules_MergesConsecutiveModules(t *testing.T) {
	modules := []calendardomain.SchoolCalendarModule{springAModule(), springBModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
		{Module: timetabledomain.ModuleSpringB, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 1 {
		t.Fatalf("len(got) = %v, want 1", len(got))
	}
	s := got[0]
	wantUntil := civil.DateTime{Date: civil.Date{Year: 2026, Month: 7, Day: 31}, Time: civil.Time{Hour: 23, Minute: 59}}
	if s.Until != wantUntil {
		t.Errorf("Until = %v, want %v", s.Until, wantUntil)
	}
}

func TestGetSchedules_ExceptionsBetweenModules(t *testing.T) {
	// Spring A ends 2026-05-31, Spring B starts 2026-06-01, no gap so no exceptions from gap.
	// Use modules with a gap to test exception generation for skipped weeks (e.g. vacation between modules).
	modA := springAModule()
	modA.End = civil.Date{Year: 2026, Month: 4, Day: 30} // ends before continuing to next module with gap
	modB := springBModule()
	modB.Start = civil.Date{Year: 2026, Month: 5, Day: 15} // gap between modA.End+1 and modB.Start

	modules := []calendardomain.SchoolCalendarModule{modA, modB}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
		{Module: timetabledomain.ModuleSpringB, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 1 {
		t.Fatalf("len(got) = %v, want 1", len(got))
	}
	wantExceptionDates := []civil.Date{
		{Year: 2026, Month: 5, Day: 4},
		{Year: 2026, Month: 5, Day: 11},
	}
	if len(got[0].Exceptions) != len(wantExceptionDates) {
		t.Fatalf("len(Exceptions) = %v, want %v", len(got[0].Exceptions), len(wantExceptionDates))
	}
	for i, want := range wantExceptionDates {
		if got[0].Exceptions[i].Date != want {
			t.Errorf("Exceptions[%d].Date = %v, want %v", i, got[0].Exceptions[i].Date, want)
		}
	}
}

func TestGetSchedules_ModuleExceptionsAndAdditions(t *testing.T) {
	mod := springAModule()
	exDate := civil.Date{Year: 2026, Month: 4, Day: 13}
	addDate := civil.Date{Year: 2026, Month: 4, Day: 20}
	mod.Exceptions[time.Monday] = []civil.Date{exDate}
	mod.Additions[time.Monday] = []civil.Date{addDate}

	modules := []calendardomain.SchoolCalendarModule{mod}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 1 {
		t.Fatalf("len(got) = %v, want 1", len(got))
	}
	s := got[0]
	if len(s.Exceptions) != 1 || s.Exceptions[0].Date != exDate {
		t.Errorf("Exceptions = %v, want [%v]", s.Exceptions, exDate)
	}
	if len(s.Additions) != 1 || s.Additions[0].Date != addDate {
		t.Errorf("Additions = %v, want [%v]", s.Additions, addDate)
	}
}

func TestGetSchedules_NoMatchingModule(t *testing.T) {
	// module list does not contain the schedule's module, so it should be skipped.
	modules := []calendardomain.SchoolCalendarModule{springBModule()}
	ss := []timetabledomain.Schedule{
		{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "Room A"},
	}
	got := calendardomain.GetSchedules(modules, ss)
	if len(got) != 0 {
		t.Errorf("len(got) = %v, want 0", len(got))
	}
}

func TestGetSchedules_Empty(t *testing.T) {
	got := calendardomain.GetSchedules(nil, nil)
	if len(got) != 0 {
		t.Errorf("len(got) = %v, want 0", len(got))
	}
}
