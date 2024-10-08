package timetablemodule

import (
	"context"

	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
//
// If the registered course is returned, course association is loaded.
type UseCase interface {
	// ListCoursesByCodes returns the courses specified by the given year and codes.
	// Even if the target courses are not found, no error will be returned.
	//
	// [Authentication] not required
	ListCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.Course, error)

	// SearchCourses returns the courses satisfied with the conditions.
	//
	// [Authentication] not required
	SearchCourses(ctx context.Context, in SearchCoursesIn) ([]*timetabledomain.Course, error)

	// UpdateCoursesBasedOnKdB retrieves data about courses from kdb and updates courses.
	//
	// [Authentication] not required
	//
	// [Permission]
	//   - PermissionExecuteBatchJob
	UpdateCoursesBasedOnKdB(ctx context.Context, year shareddomain.AcademicYear) error

	// CreateRegisteredCoursesByCodes creates new registered courses by the given year and codes.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.CourseNotFound
	//   - timetable.RegisteredCourseAlreadyExists
	CreateRegisteredCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetableappdto.RegisteredCourse, error)

	// CreateRegisteredCourseManually creates a new registered course mannually.
	//
	// [Authentication] required
	CreateRegisteredCourseManually(ctx context.Context, in CreateRegisteredCourseManuallyIn) (*timetableappdto.RegisteredCourse, error)

	// ListRegisteredCourses returns the registered courses.
	//
	// [Authentication] required
	ListRegisteredCourses(ctx context.Context, year mo.Option[shareddomain.AcademicYear]) ([]*timetableappdto.RegisteredCourse, error)

	// UpdateRegisteredCourse updates registered course specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.RegisteredCourseNotFound
	UpdateRegisteredCourse(ctx context.Context, in UpdateRegisteredCourseIn) (*timetableappdto.RegisteredCourse, error)

	// DeleteRegisteredCourse deletes registered course specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.RegisteredCourseNotFound
	DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error

	// CreateTag creates a new tag.
	//
	// [Authentication] required
	CreateTag(ctx context.Context, name shareddomain.RequiredString) (tag *timetabledomain.Tag, err error)

	// ListTags returns the tags.
	//
	// [Authentication] required
	ListTags(ctx context.Context) ([]*timetabledomain.Tag, error)

	// UpdateTag updates the tag specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.TagNotFound
	UpdateTag(ctx context.Context, in UpdateTagIn) (*timetabledomain.Tag, error)

	// DeleteTag deletes the tag specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - timetable.TagNotFound
	DeleteTag(ctx context.Context, id idtype.TagID) error

	// RearrangeTags rearranges the tags.
	// Please specify all tag ids associated with the user.
	//
	// [Authentication] required
	RearrangeTags(ctx context.Context, tagIDs []idtype.TagID) ([]*timetabledomain.Tag, error)
}

type SearchCoursesIn struct {
	Year         shareddomain.AcademicYear
	Keywords     []string // return the courses whose name contains all specified keywords
	CodePrefixes struct {
		Included []string // return the courses whose code has all specified prefixes.
		Excluded []string // return the courses whose code does not have all specified prefixes.
	}
	Schedules struct {
		FullyIncluded       []timetabledomain.Schedule // return the courses whose schedules are fully included in the specified schedules
		PartiallyOverlapped []timetabledomain.Schedule // return the courses whose schedules are partially overlapped with the specified schedules
	}
	Limit  int
	Offset int
}

type CreateRegisteredCourseManuallyIn struct {
	Year        shareddomain.AcademicYear
	Name        shareddomain.RequiredString
	Instructors string
	Credit      timetabledomain.Credit
	Methods     []timetabledomain.CourseMethod
	Schedules   []timetabledomain.Schedule
}

type UpdateRegisteredCourseIn struct {
	ID          idtype.RegisteredCourseID
	Name        mo.Option[shareddomain.RequiredString]
	Instructors mo.Option[string]
	Credit      mo.Option[timetabledomain.Credit]
	Methods     mo.Option[[]timetabledomain.CourseMethod]
	Schedules   mo.Option[[]timetabledomain.Schedule]
	Memo        mo.Option[string]
	Attendance  mo.Option[shareddomain.NonNegativeInt]
	Absence     mo.Option[shareddomain.NonNegativeInt]
	Late        mo.Option[shareddomain.NonNegativeInt]
	TagIDs      mo.Option[[]idtype.TagID]
}

type UpdateTagIn struct {
	ID   idtype.TagID
	Name mo.Option[shareddomain.RequiredString]
}
