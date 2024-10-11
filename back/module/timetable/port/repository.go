package timetableport

import (
	"context"

	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindCourse(ctx context.Context, filter CourseFilter, lock sharedport.Lock) (mo.Option[*timetabledomain.Course], error)
	ListCourses(ctx context.Context, filter CourseFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*timetabledomain.Course, error)
	CreateCourses(ctx context.Context, courses ...*timetabledomain.Course) error
	UpdateCourse(ctx context.Context, course *timetabledomain.Course) error

	FindRegisteredCourse(ctx context.Context, filter RegisteredCourseFilter, lock sharedport.Lock) (mo.Option[*timetabledomain.RegisteredCourse], error)
	ListRegisteredCourses(ctx context.Context, filter RegisteredCourseFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*timetabledomain.RegisteredCourse, error)
	CreateRegisteredCourses(ctx context.Context, registeredCourses ...*timetabledomain.RegisteredCourse) error
	UpdateRegisteredCourse(ctx context.Context, registeredCourse *timetabledomain.RegisteredCourse) error
	DeleteRegisteredCourses(ctx context.Context, filter RegisteredCourseFilter) (rowsAffected int, err error)

	FindTag(ctx context.Context, filter TagFilter, lock sharedport.Lock) (mo.Option[*timetabledomain.Tag], error)
	ListTags(ctx context.Context, filter TagFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*timetabledomain.Tag, error)
	CreateTags(ctx context.Context, tags ...*timetabledomain.Tag) error
	UpdateTag(ctx context.Context, tag *timetabledomain.Tag) error
	DeleteTags(ctx context.Context, filter TagFilter) (rowsAffected int, err error)
}

type CourseFilter struct {
	ID    mo.Option[idtype.CourseID]
	IDs   mo.Option[[]idtype.CourseID]
	Year  mo.Option[shareddomain.AcademicYear]
	Code  mo.Option[timetabledomain.Code]
	Codes mo.Option[[]timetabledomain.Code]
}

func (f *CourseFilter) IsUniqueFilter() bool {
	return f.ID.IsPresent() || (f.Year.IsPresent() && f.Code.IsPresent())
}

type RegisteredCourseFilter struct {
	ID        mo.Option[idtype.RegisteredCourseID]
	UserID    mo.Option[idtype.UserID]
	Year      mo.Option[shareddomain.AcademicYear]
	CourseIDs mo.Option[[]idtype.CourseID]
	TagID     mo.Option[idtype.TagID]
}

func (f *RegisteredCourseFilter) IsUniqueFilter() bool {
	return f.ID.IsPresent()
}

type TagFilter struct {
	ID     mo.Option[idtype.TagID]
	IDs    mo.Option[[]idtype.TagID]
	UserID mo.Option[idtype.UserID]
}

func (f *TagFilter) IsUniqueFilter() bool {
	return f.ID.IsPresent()
}
