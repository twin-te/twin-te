package timetableport

import (
	"context"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

type Integrator interface {
	// ListCourseWithoutIDsFromKdB returns the latest course data retrieved from KdB.
	ListCourseWithoutIDsFromKdB(ctx context.Context, year shareddomain.AcademicYear) ([]timetableappdto.CourseWithoutID, error)
}
