package timetableport

import (
	"time"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type CourseWithoutID struct {
	Year              shareddomain.AcademicYear
	Code              timetabledomain.Code
	Name              shareddomain.RequiredString
	Instructors       string
	Credit            timetabledomain.Credit
	Overview          string
	Remarks           string
	LastUpdatedAt     time.Time
	HasParseError     bool
	IsAnnual          bool
	RecommendedGrades []timetabledomain.RecommendedGrade
	Methods           []timetabledomain.CourseMethod
	Schedules         []timetabledomain.Schedule
}
