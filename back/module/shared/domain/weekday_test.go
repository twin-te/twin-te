package shareddomain_test

import (
	"testing"
	"time"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

func TestAllWeekdays(t *testing.T) {
	if len(shareddomain.AllWeekdays) != 7 {
		t.Fatalf("len(AllWeekdays) = %v, want 7", len(shareddomain.AllWeekdays))
	}
}

func TestParseWeekday(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Weekday
		wantErr bool
	}{
		{"sunday", "Sunday", time.Sunday, false},
		{"monday", "Monday", time.Monday, false},
		{"saturday", "Saturday", time.Saturday, false},
		{"invalid", "Funday", 0, true},
		{"empty", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shareddomain.ParseWeekday(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
