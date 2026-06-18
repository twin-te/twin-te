package calendarport

import (
	"context"

	"github.com/samber/mo"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindIcalSubscriptionByID(ctx context.Context, id idtype.IcalSubscriptionID, lock sharedport.Lock) (mo.Option[*calendardomain.IcalSubscription], error)
	FindIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID, lock sharedport.Lock) (mo.Option[*calendardomain.IcalSubscription], error)
	CreateIcalSubscription(ctx context.Context, s *calendardomain.IcalSubscription) error
	UpdateIcalSubscription(ctx context.Context, s *calendardomain.IcalSubscription) error
	DeleteIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID) (rowsAffected int, err error)
}
