package timetabledomain_test

import (
	"testing"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestCredit_String(t *testing.T) {
	c := timetabledomain.Credit("2.0")
	if c.String() != "2.0" {
		t.Errorf("String() = %v, want 2.0", c.String())
	}
}

func TestCredit_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		credit timetabledomain.Credit
		want   bool
	}{
		{"zero", timetabledomain.Credit(""), true},
		{"non-zero", timetabledomain.Credit("2.0"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.credit.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCredit(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid 2.0", "2.0", false},
		{"valid 1.5", "1.5", false},
		{"valid two digit", "12.0", false},
		{"invalid decimal", "2.3", true},
		{"no decimal", "2", true},
		{"empty", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParseCredit(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got.String() != tt.input {
				t.Errorf("got %v, want %v", got, tt.input)
			}
		})
	}
}
