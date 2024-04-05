package schoolcalendarmodule

import (
	"context"

	"cloud.google.com/go/civil"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// GetEvents returns the events specified by the given year.
	//
	// [Authentication] not required
	GetEvents(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.Event, error)

	// GetEventsByDate returns the events specified by the given date.
	//
	// [Authentication] not required
	GetEventsByDate(ctx context.Context, date civil.Date) ([]*schoolcalendardomain.Event, error)

	// GetModuleDetails returns the module details specified by the given year.
	//
	// [Authentication] not required
	GetModuleDetails(ctx context.Context, year shareddomain.AcademicYear) ([]*schoolcalendardomain.ModuleDetail, error)

	// GetModuleByDate returns the module corresponding to the given date.
	//
	// [Authentication] not required
	//
	// [Error Code]
	//   - schoolcalendar.ModuleNotFound
	GetModuleByDate(ctx context.Context, date civil.Date) (schoolcalendardomain.Module, error)
}
