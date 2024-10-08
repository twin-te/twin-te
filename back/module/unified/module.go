package unifiedmodule

import (
	"context"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// GetByDate returns the resources related to the given date.
	// Only registered courses which will be held on the given date will be returned.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - schoolcalendar.ModuleNotFound
	GetByDate(ctx context.Context, date civil.Date) (events []*schoolcalendardomain.Event, module schoolcalendardomain.Module, registeredCourses []*timetableappdto.RegisteredCourse, err error)
}
