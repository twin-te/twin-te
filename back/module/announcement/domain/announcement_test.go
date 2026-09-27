package announcementdomain_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func TestParseAnnouncementTag(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    announcementdomain.AnnouncementTag
		wantErr bool
	}{
		{"information", "Information", announcementdomain.AnnouncementTagInformation, false},
		{"notification", "Notification", announcementdomain.AnnouncementTagNotification, false},
		{"invalid", "invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := announcementdomain.ParseAnnouncementTag(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("tag = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnouncement_Clone(t *testing.T) {
	a := &announcementdomain.Announcement{
		ID:          idtype.NewAnnouncementID(),
		Tags:        []announcementdomain.AnnouncementTag{announcementdomain.AnnouncementTagInformation},
		Title:       "title",
		Content:     "content",
		PublishedAt: time.Now(),
	}

	clone := a.Clone()

	if clone.ID != a.ID || !reflect.DeepEqual(a.Tags, clone.Tags) {
		t.Fatalf("clone = %+v, want equal to %+v", clone, a)
	}

	clone.Tags[0] = announcementdomain.AnnouncementTagNotification
	if a.Tags[0] == clone.Tags[0] {
		t.Error("Clone should copy the Tags slice, not share it")
	}
}

func TestConstructAnnouncement(t *testing.T) {
	id := idtype.NewAnnouncementID()
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		a, err := announcementdomain.ConstructAnnouncement(func(a *announcementdomain.Announcement) error {
			a.ID = id
			a.Tags = []announcementdomain.AnnouncementTag{announcementdomain.AnnouncementTagInformation}
			a.Title = "title"
			a.Content = "content"
			a.PublishedAt = now
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if a.ID != id || a.Title.String() != "title" || a.Content.String() != "content" || !a.PublishedAt.Equal(now) {
			t.Errorf("got %+v", a)
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := announcementdomain.ConstructAnnouncement(func(a *announcementdomain.Announcement) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := announcementdomain.ConstructAnnouncement(func(a *announcementdomain.Announcement) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestParseAnnouncement(t *testing.T) {
	id := idtype.NewAnnouncementID()
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		a, err := announcementdomain.ParseAnnouncement(id.String(), []string{"Information", "Notification"}, "title", "content", now)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		wantTags := []announcementdomain.AnnouncementTag{
			announcementdomain.AnnouncementTagInformation,
			announcementdomain.AnnouncementTagNotification,
		}
		if a.ID != id || !reflect.DeepEqual(a.Tags, wantTags) || a.Title.String() != "title" ||
			a.Content.String() != "content" || !a.PublishedAt.Equal(now) {
			t.Errorf("got %+v", a)
		}
	})

	errTests := []struct {
		name    string
		id      string
		tags    []string
		title   string
		content string
	}{
		{"invalid id", "invalid-id", []string{"Information"}, "title", "content"},
		{"invalid tag", id.String(), []string{"invalid"}, "title", "content"},
		{"empty title", id.String(), []string{"Information"}, "", "content"},
		{"empty content", id.String(), []string{"Information"}, "title", ""},
	}
	for _, tt := range errTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := announcementdomain.ParseAnnouncement(tt.id, tt.tags, tt.title, tt.content, now)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
