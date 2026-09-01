package timetabledomain_test

import (
	"errors"
	"reflect"
	"testing"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"

	"github.com/samber/mo"
)

func newValidManualRegisteredCourse(t *testing.T) *timetabledomain.RegisteredCourse {
	rc, err := timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error {
		rc.ID = idtype.NewRegisteredCourseID()
		rc.UserID = idtype.NewUserID()
		rc.Year = mustAcademicYear(t, 2024)
		rc.Name = mo.Some(mustRequiredString(t, "manual course"))
		rc.Instructors = mo.Some("instructor")
		rc.Credit = mo.Some(mustCredit(t, "2.0"))
		rc.Methods = mo.Some([]timetabledomain.CourseMethod{timetabledomain.CourseMethodFaceToFace})
		rc.Schedules = mo.Some([]timetabledomain.Schedule{
			{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1},
		})
		rc.Memo = "memo"
		rc.Attendance = 0
		rc.Absence = 0
		rc.Late = 0
		rc.TagIDs = []idtype.TagID{idtype.NewTagID()}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to construct registered course: %v", err)
	}
	return rc
}

func newValidBasedRegisteredCourse(t *testing.T) (*timetabledomain.RegisteredCourse, *timetabledomain.Course) {
	course := newValidCourse(t)
	rc, err := timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error {
		rc.ID = idtype.NewRegisteredCourseID()
		rc.UserID = idtype.NewUserID()
		rc.Year = mustAcademicYear(t, 2024)
		rc.CourseID = mo.Some(course.ID)
		return nil
	})
	if err != nil {
		t.Fatalf("failed to construct registered course: %v", err)
	}
	return rc, course
}

func TestRegisteredCourse_HasBasedCourse(t *testing.T) {
	manual := newValidManualRegisteredCourse(t)
	if manual.HasBasedCourse() {
		t.Error("manual registered course should not have based course")
	}

	based, _ := newValidBasedRegisteredCourse(t)
	if !based.HasBasedCourse() {
		t.Error("based registered course should have based course")
	}
}

func TestRegisteredCourse_Clone(t *testing.T) {
	rc := newValidManualRegisteredCourse(t)
	clone := rc.Clone()

	if clone.ID != rc.ID || !reflect.DeepEqual(clone.TagIDs, rc.TagIDs) {
		t.Fatalf("clone = %+v, want equal to %+v", clone, rc)
	}

	clone.TagIDs[0] = idtype.NewTagID()
	if rc.TagIDs[0] == clone.TagIDs[0] {
		t.Error("Clone should copy TagIDs slice, not share it")
	}

	cloneMethods := clone.Methods.MustGet()
	cloneMethods[0] = timetabledomain.CourseMethodOthers
	if rc.Methods.MustGet()[0] == cloneMethods[0] {
		t.Error("Clone should copy Methods slice, not share it")
	}
}

func TestRegisteredCourse_BeforeUpdateHook(t *testing.T) {
	rc := newValidManualRegisteredCourse(t)
	rc.BeforeUpdateHook()

	before, ok := rc.BeforeUpdated.Get()
	if !ok {
		t.Fatal("expected BeforeUpdated to be present")
	}
	if before.ID != rc.ID {
		t.Errorf("before.ID = %v, want %v", before.ID, rc.ID)
	}
}

func TestRegisteredCourse_DetachTag(t *testing.T) {
	id1 := idtype.NewTagID()
	id2 := idtype.NewTagID()
	rc := newValidManualRegisteredCourse(t)
	rc.TagIDs = []idtype.TagID{id1, id2}

	rc.DetachTag(id1)

	if len(rc.TagIDs) != 1 || rc.TagIDs[0] != id2 {
		t.Errorf("TagIDs = %v, want [%v]", rc.TagIDs, id2)
	}
}

func TestRegisteredCourse_DetachFromCourse(t *testing.T) {
	rc, course := newValidBasedRegisteredCourse(t)

	rc.DetachFromCourse(course)

	if rc.CourseID.IsPresent() {
		t.Error("CourseID should be absent after detach")
	}
	if rc.Name.MustGet() != course.Name {
		t.Errorf("Name = %v, want %v", rc.Name.MustGet(), course.Name)
	}
	if rc.Instructors.MustGet() != course.Instructors {
		t.Errorf("Instructors = %v, want %v", rc.Instructors.MustGet(), course.Instructors)
	}
	if rc.Credit.MustGet() != course.Credit {
		t.Errorf("Credit = %v, want %v", rc.Credit.MustGet(), course.Credit)
	}
	if !reflect.DeepEqual(rc.Methods.MustGet(), course.Methods) {
		t.Errorf("Methods = %v, want %v", rc.Methods.MustGet(), course.Methods)
	}
	if !reflect.DeepEqual(rc.Schedules.MustGet(), course.Schedules) {
		t.Errorf("Schedules = %v, want %v", rc.Schedules.MustGet(), course.Schedules)
	}
}

func TestRegisteredCourse_DetachFromCourse_KeepsOverwritten(t *testing.T) {
	rc, course := newValidBasedRegisteredCourse(t)
	overwrittenName := mustRequiredString(t, "overwritten")
	rc.Name = mo.Some(overwrittenName)

	rc.DetachFromCourse(course)

	if rc.Name.MustGet() != overwrittenName {
		t.Errorf("Name = %v, want %v (should keep overwritten value)", rc.Name.MustGet(), overwrittenName)
	}
}

func TestRegisteredCourse_Update_Manual(t *testing.T) {
	rc := newValidManualRegisteredCourse(t)
	newName := mustRequiredString(t, "updated name")
	newCredit := mustCredit(t, "3.0")

	data := timetabledomain.RegisteredCourseDataToUpdate{
		Name:       mo.Some(newName),
		Credit:     mo.Some(newCredit),
		Memo:       mo.Some("updated memo"),
		Attendance: mo.Some(shareddomain.NonNegativeInt(5)),
		Absence:    mo.Some(shareddomain.NonNegativeInt(1)),
		Late:       mo.Some(shareddomain.NonNegativeInt(2)),
	}

	if err := rc.Update(data, mo.None[*timetabledomain.Course]()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rc.Name.MustGet() != newName || rc.Credit.MustGet() != newCredit || rc.Memo != "updated memo" ||
		rc.Attendance.Int() != 5 || rc.Absence.Int() != 1 || rc.Late.Int() != 2 {
		t.Errorf("got %+v", rc)
	}
}

func TestRegisteredCourse_Update_NoChange(t *testing.T) {
	rc := newValidManualRegisteredCourse(t)
	original := rc.Clone()

	if err := rc.Update(timetabledomain.RegisteredCourseDataToUpdate{}, mo.None[*timetabledomain.Course]()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(rc, original) {
		t.Errorf("expected no change, got %+v, want %+v", rc, original)
	}
}

func TestRegisteredCourse_Update_Based_SameAsBaseKeepsAbsent(t *testing.T) {
	rc, course := newValidBasedRegisteredCourse(t)

	data := timetabledomain.RegisteredCourseDataToUpdate{
		Name:        mo.Some(course.Name),
		Instructors: mo.Some(course.Instructors),
		Credit:      mo.Some(course.Credit),
		Methods:     mo.Some(course.Methods),
		Schedules:   mo.Some(course.Schedules),
	}

	if err := rc.Update(data, mo.Some(course)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rc.Name.IsPresent() || rc.Instructors.IsPresent() || rc.Credit.IsPresent() ||
		rc.Methods.IsPresent() || rc.Schedules.IsPresent() {
		t.Errorf("expected fields to stay absent when equal to base course, got %+v", rc)
	}
}

func TestRegisteredCourse_Update_Based_DifferentOverwrites(t *testing.T) {
	rc, course := newValidBasedRegisteredCourse(t)

	newName := mustRequiredString(t, "different name")
	newInstructors := "different instructors"
	newMethods := []timetabledomain.CourseMethod{timetabledomain.CourseMethodOthers}
	newSchedules := []timetabledomain.Schedule{{Module: timetabledomain.ModuleFallA, Day: timetabledomain.DayTue, Period: 3}}

	data := timetabledomain.RegisteredCourseDataToUpdate{
		Name:        mo.Some(newName),
		Instructors: mo.Some(newInstructors),
		Methods:     mo.Some(newMethods),
		Schedules:   mo.Some(newSchedules),
	}

	if err := rc.Update(data, mo.Some(course)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rc.Name.MustGet() != newName {
		t.Errorf("Name = %v, want %v", rc.Name.MustGet(), newName)
	}
	if rc.Instructors.MustGet() != newInstructors {
		t.Errorf("Instructors = %v, want %v", rc.Instructors.MustGet(), newInstructors)
	}
	if !reflect.DeepEqual(rc.Methods.MustGet(), newMethods) {
		t.Errorf("Methods = %v, want %v", rc.Methods.MustGet(), newMethods)
	}
	if !reflect.DeepEqual(rc.Schedules.MustGet(), newSchedules) {
		t.Errorf("Schedules = %v, want %v", rc.Schedules.MustGet(), newSchedules)
	}
}

func TestRegisteredCourse_Update_Manual_InstructorsOverwrite(t *testing.T) {
	rc := newValidManualRegisteredCourse(t)
	newInstructors := "different instructors"

	data := timetabledomain.RegisteredCourseDataToUpdate{
		Instructors: mo.Some(newInstructors),
	}

	if err := rc.Update(data, mo.None[*timetabledomain.Course]()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rc.Instructors.MustGet() != newInstructors {
		t.Errorf("Instructors = %v, want %v", rc.Instructors.MustGet(), newInstructors)
	}
}

func TestRegisteredCourse_Update_TagIDs(t *testing.T) {
	rc := newValidManualRegisteredCourse(t)
	newTagIDs := []idtype.TagID{idtype.NewTagID(), idtype.NewTagID()}

	if err := rc.Update(timetabledomain.RegisteredCourseDataToUpdate{TagIDs: mo.Some(newTagIDs)}, mo.None[*timetabledomain.Course]()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(rc.TagIDs, newTagIDs) {
		t.Errorf("TagIDs = %v, want %v", rc.TagIDs, newTagIDs)
	}
}

func TestConstructRegisteredCourse(t *testing.T) {
	t.Run("success manual", func(t *testing.T) {
		rc := newValidManualRegisteredCourse(t)
		if rc.ID.IsZero() {
			t.Error("expected non-zero ID")
		}
	})

	t.Run("success based", func(t *testing.T) {
		rc, _ := newValidBasedRegisteredCourse(t)
		if rc.ID.IsZero() {
			t.Error("expected non-zero ID")
		}
	})

	t.Run("fn returns error", func(t *testing.T) {
		wantErr := errors.New("boom")
		_, err := timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error {
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Errorf("err = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error { return nil })
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("manual without name missing", func(t *testing.T) {
		_, err := timetabledomain.ConstructRegisteredCourse(func(rc *timetabledomain.RegisteredCourse) error {
			rc.ID = idtype.NewRegisteredCourseID()
			rc.UserID = idtype.NewUserID()
			rc.Year = mustAcademicYear(t, 2024)
			rc.Instructors = mo.Some("instructor")
			rc.Credit = mo.Some(mustCredit(t, "2.0"))
			rc.Methods = mo.Some([]timetabledomain.CourseMethod{timetabledomain.CourseMethodFaceToFace})
			rc.Schedules = mo.Some([]timetabledomain.Schedule{{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1}})
			return nil
		})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
