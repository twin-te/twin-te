package schoolcalendardomain_test

import (
	"errors"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/samber/mo"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestEventType_IsZero(t *testing.T) {
	if !schoolcalendardomain.EventType(0).IsZero() {
		t.Error("zero value should be zero")
	}
	if schoolcalendardomain.EventTypeHoliday.IsZero() {
		t.Error("EventTypeHoliday should not be zero")
	}
}

func TestEventType_IsSubstituteDay(t *testing.T) {
	if !schoolcalendardomain.EventTypeSubstituteDay.IsSubstituteDay() {
		t.Error("EventTypeSubstituteDay should be substitute day")
	}
	if schoolcalendardomain.EventTypeHoliday.IsSubstituteDay() {
		t.Error("EventTypeHoliday should not be substitute day")
	}
}

func TestEventType_String(t *testing.T) {
	tests := []struct {
		et   schoolcalendardomain.EventType
		want string
	}{
		{schoolcalendardomain.EventTypeHoliday, "Holiday"},
		{schoolcalendardomain.EventTypePublicHoliday, "PublicHoliday"},
		{schoolcalendardomain.EventTypeExam, "Exam"},
		{schoolcalendardomain.EventTypeSubstituteDay, "SubstituteDay"},
		{schoolcalendardomain.EventTypeOther, "Other"},
	}
	for _, tt := range tests {
		if got := tt.et.String(); got != tt.want {
			t.Errorf("String() = %v, want %v", got, tt.want)
		}
	}
}

func TestParseEventType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    schoolcalendardomain.EventType
		wantErr bool
	}{
		{"holiday", "Holiday", schoolcalendardomain.EventTypeHoliday, false},
		{"public holiday", "PublicHoliday", schoolcalendardomain.EventTypePublicHoliday, false},
		{"exam", "Exam", schoolcalendardomain.EventTypeExam, false},
		{"substitute day", "SubstituteDay", schoolcalendardomain.EventTypeSubstituteDay, false},
		{"other", "Other", schoolcalendardomain.EventTypeOther, false},
		{"invalid", "invalid", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := schoolcalendardomain.ParseEventType(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_IsExamDescriptionHelpers(t *testing.T) {
	tests := []struct {
		name  string
		desc  string
		check func(e *schoolcalendardomain.Event) bool
	}{
		{"spring A", "春A 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsSpringAExam() }},
		{"spring AB", "春AB 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsSpringABExam() }},
		{"spring ABC", "春ABC 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsSpringABCExam() }},
		{"spring C", "春C 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsSpringCExam() }},
		{"fall A", "秋A 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsFallAExam() }},
		{"fall AB", "秋AB 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsFallABExam() }},
		{"fall ABC", "秋ABC 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsFallABCExam() }},
		{"fall C", "秋C 期末試験", func(e *schoolcalendardomain.Event) bool { return e.IsFallCExam() }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &schoolcalendardomain.Event{Description: tt.desc}
			if !tt.check(e) {
				t.Errorf("expected true for description %q", tt.desc)
			}
			other := &schoolcalendardomain.Event{Description: "other"}
			if tt.check(other) {
				t.Errorf("expected false for non-matching description")
			}
		})
	}
}

func TestEvent_Clone(t *testing.T) {
	id, err := idtype.ParseEventID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	e := &schoolcalendardomain.Event{
		ID:          id,
		Type:        schoolcalendardomain.EventTypeHoliday,
		Date:        civil.DateOf(time.Now()),
		Description: "desc",
	}

	clone := e.Clone()
	if *clone != *e {
		t.Fatalf("clone = %+v, want equal to %+v", clone, e)
	}
	if clone == e {
		t.Error("Clone should return a different pointer")
	}
}

func TestConstructEvent(t *testing.T) {
	id, err := idtype.ParseEventID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	date := civil.DateOf(time.Now())

	t.Run("success", func(t *testing.T) {
		e, err := schoolcalendardomain.ConstructEvent(func(e *schoolcalendardomain.Event) error {
			e.ID = id
			e.Type = schoolcalendardomain.EventTypeHoliday
			e.Date = date
			e.Description = "desc"
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if e.ID != id || e.Type != schoolcalendardomain.EventTypeHoliday || e.Date != date {
			t.Errorf("got %+v", e)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := schoolcalendardomain.ConstructEvent(func(e *schoolcalendardomain.Event) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := schoolcalendardomain.ConstructEvent(func(e *schoolcalendardomain.Event) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("substitute day without changeTo", func(t *testing.T) {
		_, err := schoolcalendardomain.ConstructEvent(func(e *schoolcalendardomain.Event) error {
			e.ID = id
			e.Type = schoolcalendardomain.EventTypeSubstituteDay
			e.Date = date
			e.Description = "desc"
			return nil
		})
		if err == nil {
			t.Error("expected error for missing ChangeTo, got nil")
		}
	})

	t.Run("substitute day with changeTo", func(t *testing.T) {
		e, err := schoolcalendardomain.ConstructEvent(func(e *schoolcalendardomain.Event) error {
			e.ID = id
			e.Type = schoolcalendardomain.EventTypeSubstituteDay
			e.Date = date
			e.Description = "desc"
			e.ChangeTo = mo.Some(time.Monday)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if v, ok := e.ChangeTo.Get(); !ok || v != time.Monday {
			t.Errorf("ChangeTo = %+v, want Monday", e.ChangeTo)
		}
	})
}

func TestParseEvent(t *testing.T) {
	t.Run("success without changeTo", func(t *testing.T) {
		e, err := schoolcalendardomain.ParseEvent(1, "Holiday", "2024-01-01", "desc", mo.None[string]())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		wantDate, _ := civil.ParseDate("2024-01-01")
		if e.Type != schoolcalendardomain.EventTypeHoliday || e.Date != wantDate || e.Description != "desc" {
			t.Errorf("got %+v", e)
		}
	})

	t.Run("success substitute day with changeTo", func(t *testing.T) {
		e, err := schoolcalendardomain.ParseEvent(1, "SubstituteDay", "2024-01-01", "desc", mo.Some("Monday"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		v, ok := e.ChangeTo.Get()
		if !ok || v != time.Monday {
			t.Errorf("ChangeTo = %+v, want Monday", e.ChangeTo)
		}
	})

	errTests := []struct {
		name        string
		id          int
		eventType   string
		date        string
		description string
		changeTo    mo.Option[string]
	}{
		{"invalid id", 0, "Holiday", "2024-01-01", "desc", mo.None[string]()},
		{"invalid event type", 1, "invalid", "2024-01-01", "desc", mo.None[string]()},
		{"invalid date", 1, "Holiday", "invalid-date", "desc", mo.None[string]()},
		{"substitute day missing changeTo", 1, "SubstituteDay", "2024-01-01", "desc", mo.None[string]()},
		{"substitute day invalid changeTo", 1, "SubstituteDay", "2024-01-01", "desc", mo.Some("invalid")},
	}
	for _, tt := range errTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := schoolcalendardomain.ParseEvent(tt.id, tt.eventType, tt.date, tt.description, tt.changeTo)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
