package timetableport

import (
	"context"

	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error) error

	// Course

	FindCourse(ctx context.Context, conds FindCourseConds, lock sharedport.Lock) (*timetabledomain.Course, error)
	ListCourses(ctx context.Context, conds ListCoursesConds, lock sharedport.Lock) ([]*timetabledomain.Course, error)
	CreateCourses(ctx context.Context, courses ...*timetabledomain.Course) error
	UpdateCourse(ctx context.Context, course *timetabledomain.Course) error

	// RegisteredCourse

	FindRegisteredCourse(ctx context.Context, conds FindRegisteredCourseConds, lock sharedport.Lock) (*timetabledomain.RegisteredCourse, error)
	ListRegisteredCourses(ctx context.Context, conds ListRegisteredCoursesConds, lock sharedport.Lock) ([]*timetabledomain.RegisteredCourse, error)
	CreateRegisteredCourses(ctx context.Context, registeredCourses ...*timetabledomain.RegisteredCourse) error
	UpdateRegisteredCourse(ctx context.Context, registeredCourse *timetabledomain.RegisteredCourse) error
	DeleteRegisteredCourses(ctx context.Context, conds DeleteRegisteredCoursesConds) (rowsAffected int, err error)

	LoadCourseAssociationToRegisteredCourse(ctx context.Context, registeredCourses []*timetabledomain.RegisteredCourse, lock sharedport.Lock) error

	// Tag

	FindTag(ctx context.Context, conds FindTagConds, lock sharedport.Lock) (*timetabledomain.Tag, error)
	ListTags(ctx context.Context, conds ListTagsConds, lock sharedport.Lock) ([]*timetabledomain.Tag, error)
	CreateTags(ctx context.Context, tags ...*timetabledomain.Tag) error
	UpdateTag(ctx context.Context, tag *timetabledomain.Tag) error
	DeleteTags(ctx context.Context, conds DeleteTagsConds) (rowsAffected int, err error)
}

// Course

type FindCourseConds struct {
	Year shareddomain.AcademicYear
	Code timetabledomain.Code
}

type ListCoursesConds struct {
	IDs   *[]idtype.CourseID
	Year  *shareddomain.AcademicYear
	Codes *[]timetabledomain.Code
}

// RegisteredCourse

type FindRegisteredCourseConds struct {
	ID     idtype.RegisteredCourseID
	UserID *idtype.UserID
}

type ListRegisteredCoursesConds struct {
	UserID    *idtype.UserID
	Year      *shareddomain.AcademicYear
	CourseIDs *[]idtype.CourseID
}

type DeleteRegisteredCoursesConds struct {
	ID     *idtype.RegisteredCourseID
	UserID *idtype.UserID
}

// Tag

type FindTagConds struct {
	ID     idtype.TagID
	UserID *idtype.UserID
}

type ListTagsConds struct {
	IDs    *[]idtype.TagID
	UserID *idtype.UserID
}

type DeleteTagsConds struct {
	ID     *idtype.TagID
	UserID *idtype.UserID
}
