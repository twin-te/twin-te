package calendarmodule

import (
	"context"

	"github.com/samber/mo"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

// UseCase represents application specific business rules.
//
// The following error codes are not stated explicitly in the each method, but may be returned.
//   - shared.InvalidArgument
//   - shared.Unauthenticated
//   - shared.Unauthorized
type UseCase interface {
	// ExportTimetableToICal returns iCalendar bytes for the authenticated user's timetable.
	//
	// [Authentication] required
	ExportTimetableToICal(ctx context.Context, year shareddomain.AcademicYear, tagIDs []idtype.TagID, isRdateSupported bool) ([]byte, error)

	// FindIcalSubscriptionID returns the iCal subscription ID if the user has enabled it.
	//
	// [Authentication] required
	FindIcalSubscriptionID(ctx context.Context) (mo.Option[idtype.IcalSubscriptionID], error)

	// EnableIcalSubscription creates or returns the iCal subscription ID for the user.
	//
	// [Authentication] required
	EnableIcalSubscription(ctx context.Context) (idtype.IcalSubscriptionID, error)

	// DisableIcalSubscription removes the public iCal URL for the user.
	//
	// [Authentication] required
	DisableIcalSubscription(ctx context.Context) error

	// ResolveUserIDByIcalSubscriptionID returns the user ID for the given public iCal subscription ID.
	//
	// [Authentication] optional
	ResolveUserIDByIcalSubscriptionID(ctx context.Context, id idtype.IcalSubscriptionID) (idtype.UserID, error)
}
