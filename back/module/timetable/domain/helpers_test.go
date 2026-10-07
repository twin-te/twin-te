package timetabledomain_test

import (
	"time"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func mustAcademicYear(t interface{ Fatalf(string, ...any) }, year int) shareddomain.AcademicYear {
	y, err := shareddomain.ParseAcademicYear(year)
	if err != nil {
		t.Fatalf("failed to parse academic year: %v", err)
	}
	return y
}

func mustRequiredString(t interface{ Fatalf(string, ...any) }, s string) shareddomain.RequiredString {
	rs, err := timetabledomain.ParseName(s)
	if err != nil {
		t.Fatalf("failed to parse required string: %v", err)
	}
	return rs
}

func mustCode(t interface{ Fatalf(string, ...any) }, s string) timetabledomain.Code {
	c, err := timetabledomain.ParseCode(s)
	if err != nil {
		t.Fatalf("failed to parse code: %v", err)
	}
	return c
}

func mustCredit(t interface{ Fatalf(string, ...any) }, s string) timetabledomain.Credit {
	c, err := timetabledomain.ParseCredit(s)
	if err != nil {
		t.Fatalf("failed to parse credit: %v", err)
	}
	return c
}

func newValidCourse(t interface{ Fatalf(string, ...any) }) *timetabledomain.Course {
	c, err := timetabledomain.ConstructCourse(func(c *timetabledomain.Course) error {
		c.ID = idtype.NewCourseID()
		c.Year = mustAcademicYear(t, 2024)
		c.Code = mustCode(t, "AB12345")
		c.Name = mustRequiredString(t, "course name")
		c.Instructors = "instructor"
		c.Credit = mustCredit(t, "2.0")
		c.Overview = "overview"
		c.Remarks = "remarks"
		c.LastUpdatedAt = time.Now()
		c.HasParseError = false
		c.IsAnnual = false
		c.RecommendedGrades = []timetabledomain.RecommendedGrade{1, 2}
		c.Methods = []timetabledomain.CourseMethod{timetabledomain.CourseMethodFaceToFace}
		c.Schedules = []timetabledomain.Schedule{
			{Module: timetabledomain.ModuleSpringA, Day: timetabledomain.DayMon, Period: 1, Locations: "room1"},
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to construct course: %v", err)
	}
	return c
}
