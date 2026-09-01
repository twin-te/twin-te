package calendardomain_test

import (
	"testing"
	"time"

	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
)

func TestJST(t *testing.T) {
	if calendardomain.JST == nil {
		t.Fatal("JST is nil")
	}
	if calendardomain.JST.String() != "Asia/Tokyo" {
		t.Errorf("JST = %v, want Asia/Tokyo", calendardomain.JST.String())
	}

	// Verify offset is UTC+9.
	tm := time.Date(2026, 7, 10, 0, 0, 0, 0, calendardomain.JST)
	_, offset := tm.Zone()
	if offset != 9*60*60 {
		t.Errorf("offset = %v, want %v", offset, 9*60*60)
	}
}
