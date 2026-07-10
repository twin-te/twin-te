package shareddomain_test

import (
	"testing"
	"time"

	"cloud.google.com/go/civil"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

func TestAcademicYear_Int(t *testing.T) {
	year := shareddomain.AcademicYear(2024)
	if year.Int() != 2024 {
		t.Errorf("Int() = %v, want 2024", year.Int())
	}
}

func TestAcademicYear_IsZero(t *testing.T) {
	tests := []struct {
		name string
		year shareddomain.AcademicYear
		want bool
	}{
		{"zero value", shareddomain.AcademicYear(0), true},
		{"non-zero value", shareddomain.AcademicYear(2024), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.year.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAcademicYear(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    shareddomain.AcademicYear
		wantErr bool
	}{
		{"positive", 2024, shareddomain.AcademicYear(2024), false},
		{"zero", 0, 0, true},
		{"negative", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shareddomain.ParseAcademicYear(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAcademicYear(t *testing.T) {
	tests := []struct {
		name    string
		year    int
		month   time.Month
		want    shareddomain.AcademicYear
		wantErr bool
	}{
		{"april is start of academic year", 2024, time.April, shareddomain.AcademicYear(2024), false},
		{"december stays in same academic year", 2024, time.December, shareddomain.AcademicYear(2024), false},
		{"january belongs to previous academic year", 2024, time.January, shareddomain.AcademicYear(2023), false},
		{"march belongs to previous academic year", 2024, time.March, shareddomain.AcademicYear(2023), false},
		{"invalid result year", 1, time.January, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shareddomain.NewAcademicYear(tt.year, tt.month)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAcademicYearFromDate(t *testing.T) {
	tests := []struct {
		name string
		date civil.Date
		want shareddomain.AcademicYear
	}{
		{"april", civil.Date{Year: 2024, Month: time.April, Day: 1}, shareddomain.AcademicYear(2024)},
		{"march", civil.Date{Year: 2024, Month: time.March, Day: 31}, shareddomain.AcademicYear(2023)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shareddomain.NewAcademicYearFromDate(tt.date)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
