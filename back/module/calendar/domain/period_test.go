package calendardomain_test

import (
	"testing"

	"cloud.google.com/go/civil"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestGetPeriodStart(t *testing.T) {
	tests := []struct {
		name   string
		period timetabledomain.Period
		want   civil.Time
	}{
		{"period 1", 1, civil.Time{Hour: 8, Minute: 40}},
		{"period 2", 2, civil.Time{Hour: 10, Minute: 10}},
		{"period 3", 3, civil.Time{Hour: 12, Minute: 15}},
		{"period 4", 4, civil.Time{Hour: 13, Minute: 45}},
		{"period 5", 5, civil.Time{Hour: 15, Minute: 15}},
		{"period 6", 6, civil.Time{Hour: 16, Minute: 45}},
		{"unknown period", 0, civil.Time{Hour: 0, Minute: 0}},
		{"out of range period", 99, civil.Time{Hour: 0, Minute: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calendardomain.GetPeriodStart(tt.period)
			if got != tt.want {
				t.Errorf("GetPeriodStart(%v) = %v, want %v", tt.period, got, tt.want)
			}
		})
	}
}

func TestGetPeriodEnd(t *testing.T) {
	tests := []struct {
		name   string
		period timetabledomain.Period
		want   civil.Time
	}{
		{"period 1", 1, civil.Time{Hour: 9, Minute: 55}},
		{"period 2", 2, civil.Time{Hour: 11, Minute: 25}},
		{"period 3", 3, civil.Time{Hour: 13, Minute: 30}},
		{"period 4", 4, civil.Time{Hour: 15, Minute: 0}},
		{"period 5", 5, civil.Time{Hour: 16, Minute: 30}},
		{"period 6", 6, civil.Time{Hour: 18, Minute: 0}},
		{"unknown period", 0, civil.Time{Hour: 23, Minute: 59}},
		{"out of range period", 99, civil.Time{Hour: 23, Minute: 59}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calendardomain.GetPeriodEnd(tt.period)
			if got != tt.want {
				t.Errorf("GetPeriodEnd(%v) = %v, want %v", tt.period, got, tt.want)
			}
		})
	}
}
