package timetabledomain_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"

	"github.com/samber/mo"
)

func TestCourse_Clone(t *testing.T) {
	c := newValidCourse(t)
	clone := c.Clone()

	if clone.ID != c.ID || !reflect.DeepEqual(clone.RecommendedGrades, c.RecommendedGrades) ||
		!reflect.DeepEqual(clone.Methods, c.Methods) || !reflect.DeepEqual(clone.Schedules, c.Schedules) {
		t.Fatalf("clone = %+v, want equal to %+v", clone, c)
	}

	clone.RecommendedGrades[0] = 6
	if c.RecommendedGrades[0] == clone.RecommendedGrades[0] {
		t.Error("Clone should copy RecommendedGrades slice, not share it")
	}

	clone.Methods[0] = timetabledomain.CourseMethodOthers
	if c.Methods[0] == clone.Methods[0] {
		t.Error("Clone should copy Methods slice, not share it")
	}

	clone.Schedules[0].Locations = "changed"
	if c.Schedules[0].Locations == clone.Schedules[0].Locations {
		t.Error("Clone should copy Schedules slice, not share it")
	}
}

func TestCourse_BeforeUpdateHook(t *testing.T) {
	c := newValidCourse(t)
	c.BeforeUpdateHook()

	before, ok := c.BeforeUpdated.Get()
	if !ok {
		t.Fatal("expected BeforeUpdated to be present")
	}
	if before.ID != c.ID {
		t.Errorf("before.ID = %v, want %v", before.ID, c.ID)
	}
}

func TestCourse_Update(t *testing.T) {
	c := newValidCourse(t)

	newName := mustRequiredString(t, "new name")
	newCredit := mustCredit(t, "3.0")
	newTime := time.Now().Add(time.Hour)
	newGrades := []timetabledomain.RecommendedGrade{3}
	newMethods := []timetabledomain.CourseMethod{timetabledomain.CourseMethodOthers}
	newSchedules := []timetabledomain.Schedule{{Module: timetabledomain.ModuleFallA, Day: timetabledomain.DayTue, Period: 2}}

	data := timetabledomain.CourseDataToUpdate{
		Name:              mo.Some(newName),
		Instructors:       mo.Some("new instructor"),
		Credit:            mo.Some(newCredit),
		Overview:          mo.Some("new overview"),
		Remarks:           mo.Some("new remarks"),
		LastUpdatedAt:     mo.Some(newTime),
		HasParseError:     mo.Some(true),
		IsAnnual:          mo.Some(true),
		RecommendedGrades: mo.Some(newGrades),
		Methods:           mo.Some(newMethods),
		Schedules:         mo.Some(newSchedules),
	}

	c.Update(data)

	if c.Name != newName || c.Instructors != "new instructor" || c.Credit != newCredit ||
		c.Overview != "new overview" || c.Remarks != "new remarks" || !c.LastUpdatedAt.Equal(newTime) ||
		!c.HasParseError || !c.IsAnnual ||
		!reflect.DeepEqual(c.RecommendedGrades, newGrades) ||
		!reflect.DeepEqual(c.Methods, newMethods) ||
		!reflect.DeepEqual(c.Schedules, newSchedules) {
		t.Errorf("got %+v", c)
	}
}

func TestCourse_Update_NoChange(t *testing.T) {
	c := newValidCourse(t)
	original := c.Clone()

	c.Update(timetabledomain.CourseDataToUpdate{})

	if !reflect.DeepEqual(c, original) {
		t.Errorf("expected no change, got %+v, want %+v", c, original)
	}
}

func TestConstructCourse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := newValidCourse(t)
		if c.ID.IsZero() {
			t.Error("expected non-zero ID")
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := timetabledomain.ConstructCourse(func(c *timetabledomain.Course) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := timetabledomain.ConstructCourse(func(c *timetabledomain.Course) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
