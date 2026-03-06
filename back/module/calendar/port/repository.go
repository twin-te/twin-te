package calendarport

import (
	"context"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindIcalSubscriptionByID(ctx context.Context, id idtype.IcalSubscriptionID, lock sharedport.Lock) (mo.Option[idtype.UserID], error)
	FindIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID, lock sharedport.Lock) (mo.Option[idtype.IcalSubscriptionID], error)
	CreateIcalSubscription(ctx context.Context, id idtype.IcalSubscriptionID, userID idtype.UserID) error
	DeleteIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID) (rowsAffected int, err error)
}
