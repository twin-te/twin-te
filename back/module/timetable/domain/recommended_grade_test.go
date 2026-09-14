package timetabledomain_test

import (
	"testing"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestRecommendedGrade_Int(t *testing.T) {
	g := timetabledomain.RecommendedGrade(3)
	if g.Int() != 3 {
		t.Errorf("Int() = %v, want 3", g.Int())
	}
}

func TestRecommendedGrade_IsZero(t *testing.T) {
	tests := []struct {
		name  string
		grade timetabledomain.RecommendedGrade
		want  bool
	}{
		{"zero", timetabledomain.RecommendedGrade(0), true},
		{"non-zero", timetabledomain.RecommendedGrade(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grade.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRecommendedGrade(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    timetabledomain.RecommendedGrade
		wantErr bool
	}{
		{"min", 1, 1, false},
		{"max", 6, 6, false},
		{"mid", 3, 3, false},
		{"too small", 0, 0, true},
		{"too large", 7, 0, true},
		{"negative", -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParseRecommendedGrade(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
