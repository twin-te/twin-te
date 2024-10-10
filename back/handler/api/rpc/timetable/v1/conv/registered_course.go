package timetablev1conv

import (
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	timetablev1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/timetable/v1"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

func ToPBRegisteredCourse(registeredCourse *timetableappdto.RegisteredCourse) (pbRegisteredCourse *timetablev1.RegisteredCourse, err error) {
	pbRegisteredCourse = &timetablev1.RegisteredCourse{
		Id:          sharedconv.ToPBUUID(registeredCourse.ID),
		UserId:      sharedconv.ToPBUUID(registeredCourse.UserID),
		Year:        sharedconv.ToPBAcademicYear(registeredCourse.Year),
		Code:        base.OptionMapByString(registeredCourse.Code).ToPointer(),
		Name:        registeredCourse.Name.String(),
		Instructors: registeredCourse.Instructors,
		Credit:      registeredCourse.Credit.String(),
		Memo:        registeredCourse.Memo,
		Attendance:  int32(registeredCourse.Attendance),
		Absence:     int32(registeredCourse.Absence),
		Late:        int32(registeredCourse.Late),
		TagIds:      base.Map(registeredCourse.TagIDs, sharedconv.ToPBUUID[idtype.TagID]),
	}

	pbRegisteredCourse.Methods, err = base.MapWithErr(registeredCourse.Methods, ToPBCourseMethod)
	if err != nil {
		return
	}

	pbRegisteredCourse.Schedules, err = base.MapWithErr(registeredCourse.Schedules, ToPBSchedule)
	if err != nil {
		return
	}

	return
}
