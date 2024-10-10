package timetableappdto

import (
	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type RegisteredCourse struct {
	ID          idtype.RegisteredCourseID
	UserID      idtype.UserID
	Year        shareddomain.AcademicYear
	CourseID    mo.Option[idtype.CourseID]
	Code        mo.Option[timetabledomain.Code]
	Name        shareddomain.RequiredString
	Instructors string
	Credit      timetabledomain.Credit
	Methods     []timetabledomain.CourseMethod
	Schedules   []timetabledomain.Schedule
	Memo        string
	Attendance  shareddomain.NonNegativeInt
	Absence     shareddomain.NonNegativeInt
	Late        shareddomain.NonNegativeInt
	TagIDs      []idtype.TagID
}
