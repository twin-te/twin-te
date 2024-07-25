package timetableport

import (
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetabledto "github.com/twin-te/twin-te/back/module/timetable/dto"
)

type Factory interface {
	NewCourse(courseWithoutID timetabledto.CourseWithoutID) (*timetabledomain.Course, error)

	NewRegisteredCourseFromCourse(userID idtype.UserID, course *timetabledomain.Course) (*timetabledomain.RegisteredCourse, error)

	NewRegisteredCourseMannualy(
		userID idtype.UserID,
		year shareddomain.AcademicYear,
		name shareddomain.RequiredString,
		instructors string,
		credit timetabledomain.Credit,
		methods []timetabledomain.CourseMethod,
		schedules []timetabledomain.Schedule,
	) (*timetabledomain.RegisteredCourse, error)

	NewTag(
		userID idtype.UserID,
		name shareddomain.RequiredString,
	) (*timetabledomain.Tag, error)
}
