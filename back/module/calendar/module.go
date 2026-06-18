package calendarmodule

import (
	"context"

	"github.com/samber/mo"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
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

	// ExportTimetableToICalByIcalSubscriptionID returns iCalendar bytes for the user bound to the given iCal subscription ID.
	//
	// [Authentication] not required
	ExportTimetableToICalByIcalSubscriptionID(ctx context.Context, id idtype.IcalSubscriptionID, year shareddomain.AcademicYear, tagIDs []idtype.TagID, isRdateSupported bool) ([]byte, error)

	// FindIcalSubscription はユーザーが連携を有効化していれば iCal 連携（id, mode, 対象タグ ID）を返す。
	//
	// [Authentication] required
	FindIcalSubscription(ctx context.Context) (mo.Option[*calendardomain.IcalSubscription], error)

	// EnableIcalSubscription はユーザーの iCal 連携 ID を作成または既存のものを返す。
	// 新規作成時のモードは SYNC、対象タグは空がデフォルト。
	//
	// [Authentication] required
	EnableIcalSubscription(ctx context.Context) (idtype.IcalSubscriptionID, error)

	// DisableIcalSubscription はユーザーの公開 iCal URL を削除する。
	//
	// [Authentication] required
	DisableIcalSubscription(ctx context.Context) error

	// UpdateIcalSubscription はユーザーの iCal 連携のモードと対象タグ ID を上書きする。
	//
	// [Authentication] required
	UpdateIcalSubscription(ctx context.Context, mode calendardomain.IcalSubscriptionMode, targetTagIDs []idtype.TagID) error
}
