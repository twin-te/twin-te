package timetabledomain_test

import (
	"testing"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestCode_String(t *testing.T) {
	c := timetabledomain.Code("AB12345")
	if c.String() != "AB12345" {
		t.Errorf("String() = %v, want AB12345", c.String())
	}
}

func TestCode_IsZero(t *testing.T) {
	tests := []struct {
		name string
		code timetabledomain.Code
		want bool
	}{
		{"zero", timetabledomain.Code(""), true},
		{"non-zero", timetabledomain.Code("AB12345"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.code.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid", "AB12345", false},
		{"valid all digits", "1234567", false},
		{"too short", "AB1234", true},
		{"too long", "AB123456", true},
		{"lowercase", "ab12345", true},
		{"empty", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParseCode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got.String() != tt.input {
				t.Errorf("got %v, want %v", got, tt.input)
			}
		})
	}
}
