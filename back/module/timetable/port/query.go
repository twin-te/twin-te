package timetableport

import (
	"context"

	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

type Query interface {
	FindRegisteredCourses(ctx context.Context, id idtype.RegisteredCourseID) (mo.Option[*timetableappdto.RegisteredCourse], error)
	ListRegisteredCourses(ctx context.Context, conds ListRegisteredCoursesConds) ([]*timetableappdto.RegisteredCourse, error)
}

type ListRegisteredCoursesConds struct {
	IDs    mo.Option[[]idtype.RegisteredCourseID]
	UserID mo.Option[idtype.UserID]
	Year   mo.Option[shareddomain.AcademicYear]
}
