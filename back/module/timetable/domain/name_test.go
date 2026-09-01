package timetabledomain_test

import (
	"testing"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestParseName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := timetabledomain.ParseName("course name")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.String() != "course name" {
			t.Errorf("got %v, want course name", got.String())
		}
	})

	t.Run("empty", func(t *testing.T) {
		_, err := timetabledomain.ParseName("")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
