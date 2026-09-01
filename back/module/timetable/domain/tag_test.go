package timetabledomain_test

import (
	"errors"
	"testing"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"

	"github.com/samber/mo"
)

func newValidTag(t *testing.T) *timetabledomain.Tag {
	tag, err := timetabledomain.ConstructTag(func(tag *timetabledomain.Tag) error {
		tag.ID = idtype.NewTagID()
		tag.UserID = idtype.NewUserID()
		tag.Name = mustRequiredString(t, "tag name")
		tag.Order = 0
		return nil
	})
	if err != nil {
		t.Fatalf("failed to construct tag: %v", err)
	}
	return tag
}

func TestTag_Clone(t *testing.T) {
	tag := newValidTag(t)
	clone := tag.Clone()

	if clone.ID != tag.ID || clone.UserID != tag.UserID || clone.Name != tag.Name || clone.Order != tag.Order {
		t.Fatalf("clone = %+v, want equal to %+v", clone, tag)
	}
}

func TestTag_BeforeUpdateHook(t *testing.T) {
	tag := newValidTag(t)
	tag.BeforeUpdateHook()

	before, ok := tag.BeforeUpdated.Get()
	if !ok {
		t.Fatal("expected BeforeUpdated to be present")
	}
	if before.ID != tag.ID {
		t.Errorf("before.ID = %v, want %v", before.ID, tag.ID)
	}
}

func TestTag_Update(t *testing.T) {
	tag := newValidTag(t)
	newName := mustRequiredString(t, "new name")

	tag.Update(timetabledomain.TagDataToUpdate{Name: mo.Some(newName)})

	if tag.Name != newName {
		t.Errorf("Name = %v, want %v", tag.Name, newName)
	}
}

func TestTag_Update_NoChange(t *testing.T) {
	tag := newValidTag(t)
	original := tag.Name

	tag.Update(timetabledomain.TagDataToUpdate{})

	if tag.Name != original {
		t.Errorf("expected no change, got %v, want %v", tag.Name, original)
	}
}

func TestConstructTag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tag := newValidTag(t)
		if tag.ID.IsZero() {
			t.Error("expected non-zero ID")
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := timetabledomain.ConstructTag(func(tag *timetabledomain.Tag) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := timetabledomain.ConstructTag(func(tag *timetabledomain.Tag) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestRearrangeTags(t *testing.T) {
	id1 := idtype.NewTagID()
	id2 := idtype.NewTagID()
	id3 := idtype.NewTagID()

	tags := []*timetabledomain.Tag{
		{ID: id1, Order: 0},
		{ID: id2, Order: 1},
		{ID: id3, Order: 2},
	}

	timetabledomain.RearrangeTags(tags, []idtype.TagID{id3, id1, id2})

	wantOrder := map[idtype.TagID]int{id3: 0, id1: 1, id2: 2}
	for _, tag := range tags {
		if tag.Order.Int() != wantOrder[tag.ID] {
			t.Errorf("tag %v order = %v, want %v", tag.ID, tag.Order.Int(), wantOrder[tag.ID])
		}
	}
}
