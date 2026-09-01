package authdomain_test

import (
	"errors"
	"testing"
	"time"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestConstructSession(t *testing.T) {
	id := idtype.NewSessionID()
	userID := idtype.NewUserID()
	expiredAt := time.Now().Add(24 * time.Hour)

	t.Run("success", func(t *testing.T) {
		s, err := authdomain.ConstructSession(func(s *authdomain.Session) error {
			s.ID = id
			s.UserID = userID
			s.ExpiredAt = expiredAt
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.ID != id || s.UserID != userID || !s.ExpiredAt.Equal(expiredAt) {
			t.Errorf("got %+v", s)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := authdomain.ConstructSession(func(s *authdomain.Session) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := authdomain.ConstructSession(func(s *authdomain.Session) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("missing expired at", func(t *testing.T) {
		_, err := authdomain.ConstructSession(func(s *authdomain.Session) error {
			s.ID = id
			s.UserID = userID
			return nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestSessionLifeTime(t *testing.T) {
	if authdomain.SessionLifeTime <= 0 {
		t.Errorf("SessionLifeTime = %v, want > 0", authdomain.SessionLifeTime)
	}
}
