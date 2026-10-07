package schoolcalendardomain_test

import (
	"errors"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestModule_IsZero(t *testing.T) {
	if !schoolcalendardomain.Module(0).IsZero() {
		t.Error("zero value should be zero")
	}
	if schoolcalendardomain.ModuleSpringA.IsZero() {
		t.Error("ModuleSpringA should not be zero")
	}
}

func TestModule_String(t *testing.T) {
	tests := []struct {
		m    schoolcalendardomain.Module
		want string
	}{
		{schoolcalendardomain.ModuleSpringA, "SpringA"},
		{schoolcalendardomain.ModuleSpringB, "SpringB"},
		{schoolcalendardomain.ModuleSpringC, "SpringC"},
		{schoolcalendardomain.ModuleSummerVacation, "SummerVacation"},
		{schoolcalendardomain.ModuleFallA, "FallA"},
		{schoolcalendardomain.ModuleFallB, "FallB"},
		{schoolcalendardomain.ModuleWinterVacation, "WinterVacation"},
		{schoolcalendardomain.ModuleFallC, "FallC"},
		{schoolcalendardomain.ModuleSpringVacation, "SpringVacation"},
	}
	for _, tt := range tests {
		if got := tt.m.String(); got != tt.want {
			t.Errorf("String() = %v, want %v", got, tt.want)
		}
	}
}

func TestParseModule(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    schoolcalendardomain.Module
		wantErr bool
	}{
		{"spring a", "SpringA", schoolcalendardomain.ModuleSpringA, false},
		{"winter vacation", "WinterVacation", schoolcalendardomain.ModuleWinterVacation, false},
		{"invalid", "invalid", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := schoolcalendardomain.ParseModule(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModuleDetail_Clone(t *testing.T) {
	id, err := idtype.ParseModuleDetailID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	year, err := shareddomain.ParseAcademicYear(2024)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	md := &schoolcalendardomain.ModuleDetail{
		ID:     id,
		Year:   year,
		Module: schoolcalendardomain.ModuleSpringA,
		Start:  civil.DateOf(time.Now()),
		End:    civil.DateOf(time.Now().AddDate(0, 1, 0)),
	}

	clone := md.Clone()
	if *clone != *md {
		t.Fatalf("clone = %+v, want equal to %+v", clone, md)
	}
	if clone == md {
		t.Error("Clone should return a different pointer")
	}
}

func TestConstructModuleDetail(t *testing.T) {
	id, err := idtype.ParseModuleDetailID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	year, err := shareddomain.ParseAcademicYear(2024)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	start := civil.DateOf(time.Now())
	end := civil.DateOf(time.Now().AddDate(0, 1, 0))

	t.Run("success", func(t *testing.T) {
		md, err := schoolcalendardomain.ConstructModuleDetail(func(md *schoolcalendardomain.ModuleDetail) error {
			md.ID = id
			md.Year = year
			md.Module = schoolcalendardomain.ModuleSpringA
			md.Start = start
			md.End = end
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if md.ID != id || md.Year != year || md.Module != schoolcalendardomain.ModuleSpringA {
			t.Errorf("got %+v", md)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := schoolcalendardomain.ConstructModuleDetail(func(md *schoolcalendardomain.ModuleDetail) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := schoolcalendardomain.ConstructModuleDetail(func(md *schoolcalendardomain.ModuleDetail) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestParseModuleDetail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		md, err := schoolcalendardomain.ParseModuleDetail(1, 2024, "SpringA", "2024-04-01", "2024-05-31")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		wantStart, _ := civil.ParseDate("2024-04-01")
		wantEnd, _ := civil.ParseDate("2024-05-31")
		if md.Module != schoolcalendardomain.ModuleSpringA || md.Start != wantStart || md.End != wantEnd {
			t.Errorf("got %+v", md)
		}
	})

	errTests := []struct {
		name   string
		id     int
		year   int
		module string
		start  string
		end    string
	}{
		{"invalid id", 0, 2024, "SpringA", "2024-04-01", "2024-05-31"},
		{"invalid year", 1, 0, "SpringA", "2024-04-01", "2024-05-31"},
		{"invalid module", 1, 2024, "invalid", "2024-04-01", "2024-05-31"},
		{"invalid start", 1, 2024, "SpringA", "invalid-date", "2024-05-31"},
		{"invalid end", 1, 2024, "SpringA", "2024-04-01", "invalid-date"},
	}
	for _, tt := range errTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := schoolcalendardomain.ParseModuleDetail(tt.id, tt.year, tt.module, tt.start, tt.end)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
