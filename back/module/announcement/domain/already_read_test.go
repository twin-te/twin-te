package announcementdomain_test

import (
	"errors"
	"testing"
	"time"

	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestConstructAlreadyRead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := idtype.NewAlreadyReadID()
		userID := idtype.NewUserID()
		announcementID := idtype.NewAnnouncementID()
		now := time.Now()

		ar, err := announcementdomain.ConstructAlreadyRead(func(ar *announcementdomain.AlreadyRead) error {
			ar.ID = id
			ar.UserID = userID
			ar.AnnouncementID = announcementID
			ar.ReadAt = now
			return nil
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ar.ID != id {
			t.Errorf("ID = %v, want %v", ar.ID, id)
		}
		if ar.UserID != userID {
			t.Errorf("UserID = %v, want %v", ar.UserID, userID)
		}
		if ar.AnnouncementID != announcementID {
			t.Errorf("AnnouncementID = %v, want %v", ar.AnnouncementID, announcementID)
		}
		if !ar.ReadAt.Equal(now) {
			t.Errorf("ReadAt = %v, want %v", ar.ReadAt, now)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := announcementdomain.ConstructAlreadyRead(func(ar *announcementdomain.AlreadyRead) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := announcementdomain.ConstructAlreadyRead(func(ar *announcementdomain.AlreadyRead) error {
			return nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
