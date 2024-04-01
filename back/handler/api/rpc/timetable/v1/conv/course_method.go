package timetablev1conv

import (
	"fmt"

	timetablev1 "github.com/twin-te/twinte-back/handler/api/rpcgen/timetable/v1"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func FromPBCourseMethod(pbCourseMethod timetablev1.CourseMethod) (timetabledomain.CourseMethod, error) {
	switch pbCourseMethod {
	case timetablev1.CourseMethod_COURSE_METHOD_ONLINE_ASYNCHRONOUS:
		return timetabledomain.CourseMethodOnlineAsynchronous, nil
	case timetablev1.CourseMethod_COURSE_METHOD_ONLINE_SYNCHRONOUS:
		return timetabledomain.CourseMethodOnlineSynchronous, nil
	case timetablev1.CourseMethod_COURSE_METHOD_FACE_TO_FACE:
		return timetabledomain.CourseMethodFaceToFace, nil
	case timetablev1.CourseMethod_COURSE_METHOD_OTHERS:
		return timetabledomain.CourseMethodOthers, nil
	}
	return 0, fmt.Errorf("invalid %#v", pbCourseMethod)
}

func ToPBCourseMethod(courseMethod timetabledomain.CourseMethod) (timetablev1.CourseMethod, error) {
	switch courseMethod {
	case timetabledomain.CourseMethodOnlineAsynchronous:
		return timetablev1.CourseMethod_COURSE_METHOD_ONLINE_ASYNCHRONOUS, nil
	case timetabledomain.CourseMethodOnlineSynchronous:
		return timetablev1.CourseMethod_COURSE_METHOD_ONLINE_SYNCHRONOUS, nil
	case timetabledomain.CourseMethodFaceToFace:
		return timetablev1.CourseMethod_COURSE_METHOD_FACE_TO_FACE, nil
	case timetabledomain.CourseMethodOthers:
		return timetablev1.CourseMethod_COURSE_METHOD_OTHERS, nil
	}
	return 0, fmt.Errorf("invalid %#v", courseMethod)
}
