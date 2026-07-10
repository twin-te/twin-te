package timetabledomain_test

import (
	"testing"
	"time"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestParseModule(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    timetabledomain.Module
		wantErr bool
	}{
		{"spring a", "SpringA", timetabledomain.ModuleSpringA, false},
		{"spring b", "SpringB", timetabledomain.ModuleSpringB, false},
		{"spring c", "SpringC", timetabledomain.ModuleSpringC, false},
		{"fall a", "FallA", timetabledomain.ModuleFallA, false},
		{"fall b", "FallB", timetabledomain.ModuleFallB, false},
		{"fall c", "FallC", timetabledomain.ModuleFallC, false},
		{"summer vacation", "SummerVacation", timetabledomain.ModuleSummerVacation, false},
		{"spring vacation", "SpringVacation", timetabledomain.ModuleSpringVacation, false},
		{"invalid", "invalid", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParseModule(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if !tt.wantErr && got.String() != tt.input {
				t.Errorf("String() = %v, want %v", got.String(), tt.input)
			}
		})
	}
}

func TestParseDay(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    timetabledomain.Day
		wantErr bool
	}{
		{"sun", "Sun", timetabledomain.DaySun, false},
		{"mon", "Mon", timetabledomain.DayMon, false},
		{"tue", "Tue", timetabledomain.DayTue, false},
		{"wed", "Wed", timetabledomain.DayWed, false},
		{"thu", "Thu", timetabledomain.DayThu, false},
		{"fri", "Fri", timetabledomain.DayFri, false},
		{"sat", "Sat", timetabledomain.DaySat, false},
		{"intensive", "Intensive", timetabledomain.DayIntensive, false},
		{"appointment", "Appointment", timetabledomain.DayAppointment, false},
		{"anytime", "AnyTime", timetabledomain.DayAnyTime, false},
		{"nt", "NT", timetabledomain.DayNT, false},
		{"invalid", "invalid", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParseDay(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if !tt.wantErr && got.String() != tt.input {
				t.Errorf("String() = %v, want %v", got.String(), tt.input)
			}
		})
	}
}

func TestDay_Weekday(t *testing.T) {
	tests := []struct {
		day  timetabledomain.Day
		want time.Weekday
	}{
		{timetabledomain.DaySun, time.Sunday},
		{timetabledomain.DayMon, time.Monday},
		{timetabledomain.DayTue, time.Tuesday},
		{timetabledomain.DayWed, time.Wednesday},
		{timetabledomain.DayThu, time.Thursday},
		{timetabledomain.DayFri, time.Friday},
		{timetabledomain.DaySat, time.Saturday},
	}
	for _, tt := range tests {
		t.Run(tt.day.String(), func(t *testing.T) {
			if got := tt.day.Weekday(); got != tt.want {
				t.Errorf("Weekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay_Weekday_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for special day, got none")
		}
	}()
	timetabledomain.DayIntensive.Weekday()
}

func TestDay_IsNormal(t *testing.T) {
	tests := []struct {
		day  timetabledomain.Day
		want bool
	}{
		{timetabledomain.DaySun, true},
		{timetabledomain.DaySat, true},
		{timetabledomain.DayIntensive, false},
		{timetabledomain.Day(0), false},
	}
	for _, tt := range tests {
		if got := tt.day.IsNormal(); got != tt.want {
			t.Errorf("IsNormal(%v) = %v, want %v", tt.day, got, tt.want)
		}
	}
}

func TestDay_IsSpecial(t *testing.T) {
	tests := []struct {
		day  timetabledomain.Day
		want bool
	}{
		{timetabledomain.DayIntensive, true},
		{timetabledomain.DayNT, true},
		{timetabledomain.DaySun, false},
		{timetabledomain.Day(0), false},
	}
	for _, tt := range tests {
		if got := tt.day.IsSpecial(); got != tt.want {
			t.Errorf("IsSpecial(%v) = %v, want %v", tt.day, got, tt.want)
		}
	}
}

func TestPeriod_Int(t *testing.T) {
	p := timetabledomain.Period(3)
	if p.Int() != 3 {
		t.Errorf("Int() = %v, want 3", p.Int())
	}
}

func TestPeriod_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		period timetabledomain.Period
		want   bool
	}{
		{"zero", timetabledomain.Period(0), true},
		{"non-zero", timetabledomain.Period(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.period.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePeriod(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    timetabledomain.Period
		wantErr bool
	}{
		{"min", 1, 1, false},
		{"max", 8, 8, false},
		{"too small", 0, 0, true},
		{"too large", 9, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParsePeriod(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule_IsNormal(t *testing.T) {
	tests := []struct {
		name     string
		schedule timetabledomain.Schedule
		want     bool
	}{
		{"normal with period", timetabledomain.Schedule{Day: timetabledomain.DayMon, Period: 1}, true},
		{"normal without period", timetabledomain.Schedule{Day: timetabledomain.DayMon, Period: 0}, false},
		{"special day", timetabledomain.Schedule{Day: timetabledomain.DayIntensive, Period: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.schedule.IsNormal(); got != tt.want {
				t.Errorf("IsNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule_IsSpecial(t *testing.T) {
	tests := []struct {
		name     string
		schedule timetabledomain.Schedule
		want     bool
	}{
		{"special day", timetabledomain.Schedule{Day: timetabledomain.DayIntensive}, true},
		{"normal day", timetabledomain.Schedule{Day: timetabledomain.DayMon}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.schedule.IsSpecial(); got != tt.want {
				t.Errorf("IsSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructSchedule(t *testing.T) {
	t.Run("valid normal schedule", func(t *testing.T) {
		s, err := timetabledomain.ConstructSchedule(func() (timetabledomain.Schedule, error) {
			return timetabledomain.Schedule{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1}, nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Day != timetabledomain.DayMon || s.Period != 1 {
			t.Errorf("got %+v", s)
		}
	})

	t.Run("valid special schedule", func(t *testing.T) {
		s, err := timetabledomain.ConstructSchedule(func() (timetabledomain.Schedule, error) {
			return timetabledomain.Schedule{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayIntensive}, nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Day != timetabledomain.DayIntensive {
			t.Errorf("got %+v", s)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		_, err := timetabledomain.ConstructSchedule(func() (timetabledomain.Schedule, error) {
			return timetabledomain.Schedule{}, errBoom
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("normal day without period", func(t *testing.T) {
		_, err := timetabledomain.ConstructSchedule(func() (timetabledomain.Schedule, error) {
			return timetabledomain.Schedule{Day: timetabledomain.DayMon, Period: 0}, nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("special day with period", func(t *testing.T) {
		_, err := timetabledomain.ConstructSchedule(func() (timetabledomain.Schedule, error) {
			return timetabledomain.Schedule{Day: timetabledomain.DayIntensive, Period: 1}, nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestParseSchedule(t *testing.T) {
	t.Run("valid normal", func(t *testing.T) {
		s, err := timetabledomain.ParseSchedule("SpringA", "Mon", 1, "room1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Module != timetabledomain.ModuleSpringA || s.Day != timetabledomain.DayMon || s.Period != 1 || s.Locations != "room1" {
			t.Errorf("got %+v", s)
		}
	})

	t.Run("valid special", func(t *testing.T) {
		s, err := timetabledomain.ParseSchedule("SpringA", "Intensive", 0, "room1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Day != timetabledomain.DayIntensive || !s.Period.IsZero() {
			t.Errorf("got %+v", s)
		}
	})

	t.Run("invalid module", func(t *testing.T) {
		_, err := timetabledomain.ParseSchedule("invalid", "Mon", 1, "room1")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("invalid day", func(t *testing.T) {
		_, err := timetabledomain.ParseSchedule("SpringA", "invalid", 1, "room1")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("invalid period for normal day", func(t *testing.T) {
		_, err := timetabledomain.ParseSchedule("SpringA", "Mon", 0, "room1")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

var errBoom = &boomErr{}

type boomErr struct{}

func (e *boomErr) Error() string { return "boom" }
