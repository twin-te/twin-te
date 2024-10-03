package timetableport

import (
	"context"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type Integrator interface {
	// GetCourseWithoutIDsFromKdB returns the latest course data retrieved from KdB.
	GetCourseWithoutIDsFromKdB(ctx context.Context, year shareddomain.AcademicYear) ([]timetabledomain.CourseWithoutID, error)
}
