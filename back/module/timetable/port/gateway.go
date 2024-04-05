package timetableport

import (
	"context"

	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

type Gateway interface {
	// GetCourseWithoutIDsFromKdB returns the latest course data retrieved from KdB.
	GetCourseWithoutIDsFromKdB(ctx context.Context, year shareddomain.AcademicYear) ([]CourseWithoutID, error)
}
