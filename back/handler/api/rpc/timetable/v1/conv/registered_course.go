package timetablev1conv

import (
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	sharedconv "github.com/twin-te/twinte-back/handler/api/rpc/shared/conv"
	timetablev1 "github.com/twin-te/twinte-back/handler/api/rpcgen/timetable/v1"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func ToPBRegisteredCourse(registeredCourse *timetabledomain.RegisteredCourse) (pbRegisteredCourse *timetablev1.RegisteredCourse, err error) {
	pbRegisteredCourse = &timetablev1.RegisteredCourse{
		Id:          sharedconv.ToPBUUID(registeredCourse.ID),
		UserId:      sharedconv.ToPBUUID(registeredCourse.UserID),
		Year:        sharedconv.ToPBAcademicYear(registeredCourse.Year),
		Name:        registeredCourse.GetName().String(),
		Instructors: registeredCourse.GetInstructors(),
		Credit:      registeredCourse.GetCredit().String(),
		Memo:        registeredCourse.Memo,
		Attendance:  int32(registeredCourse.Attendance),
		Absence:     int32(registeredCourse.Absence),
		Late:        int32(registeredCourse.Late),
		TagIds:      base.Map(registeredCourse.TagIDs, sharedconv.ToPBUUID[idtype.TagID]),
	}

	if course, ok := registeredCourse.CourseAssociation.Get(); ok {
		pbRegisteredCourse.Code = lo.ToPtr(course.Code.String())
	}

	pbRegisteredCourse.Methods, err = base.MapWithErr(registeredCourse.GetMethods(), ToPBCourseMethod)
	if err != nil {
		return
	}

	pbRegisteredCourse.Schedules, err = base.MapWithErr(registeredCourse.GetSchedules(), ToPBSchedule)
	if err != nil {
		return
	}

	return
}
