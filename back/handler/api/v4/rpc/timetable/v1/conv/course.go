package timetablev1conv

import (
	sharedconv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/shared/conv"
	timetablev1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/timetable/v1"

	"github.com/twin-te/twin-te/back/base"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func ToPBCourse(course *timetabledomain.Course) (pbCourse *timetablev1.Course, err error) {
	pbCourse = &timetablev1.Course{
		Id:                sharedconv.ToPBUUID(course.ID),
		Year:              sharedconv.ToPBAcademicYear(course.Year),
		Code:              course.Code.String(),
		Name:              course.Name.String(),
		Instructors:       course.Instructors,
		Credit:            course.Credit.String(),
		Overview:          course.Overview,
		Remarks:           course.Remarks,
		LastUpdatedAt:     sharedconv.ToPBRFC3339DateTime(course.LastUpdatedAt),
		HasParseError:     course.HasParseError,
		IsAnnual:          course.IsAnnual,
		RecommendedGrades: base.Map(course.RecommendedGrades, ToPBRecommendedGrade),
	}

	pbCourse.Methods, err = base.MapWithErr(course.Methods, ToPBCourseMethod)
	if err != nil {
		return
	}

	pbCourse.Schedules, err = base.MapWithErr(course.Schedules, ToPBSchedule)
	if err != nil {
		return
	}

	return
}
