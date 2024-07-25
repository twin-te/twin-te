package timetableport

import (
	"context"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetabledto "github.com/twin-te/twin-te/back/module/timetable/dto"
)

type Gateway interface {
	// GetCourseWithoutIDsFromKdB returns the latest course data retrieved from KdB.
	GetCourseWithoutIDsFromKdB(ctx context.Context, year shareddomain.AcademicYear) ([]timetabledto.CourseWithoutID, error)
}
