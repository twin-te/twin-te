package timetabledomain

import (
	"fmt"

	"github.com/twin-te/twin-te/back/base"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=CourseMethod -trimprefix=CourseMethod -output=course_method_string.gen.go
type CourseMethod int

const (
	CourseMethodOnlineAsynchronous CourseMethod = iota + 1
	CourseMethodOnlineSynchronous
	CourseMethodFaceToFace
	CourseMethodOthers
)

var AllCourseMethods = []CourseMethod{
	CourseMethodOnlineAsynchronous,
	CourseMethodOnlineSynchronous,
	CourseMethodFaceToFace,
	CourseMethodOthers,
}

func ParseCourseMethod(s string) (CourseMethod, error) {
	courseMethod, ok := base.FindByString(AllCourseMethods, s)
	if ok {
		return courseMethod, nil
	}
	return 0, fmt.Errorf("failed to parse CourseMethod %#v", s)
}
