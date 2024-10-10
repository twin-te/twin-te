package donationport

import (
	"context"

	"github.com/samber/mo"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindPaymentUser(ctx context.Context, conds PaymentUserFilter, lock sharedport.Lock) (mo.Option[*donationdomain.PaymentUser], error)
	ListPaymentUsers(ctx context.Context, filter PaymentUserFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*donationdomain.PaymentUser, error)
	CreatePaymentUsers(ctx context.Context, paymentUsers ...*donationdomain.PaymentUser) error
	UpdatePaymentUser(ctx context.Context, paymentUser *donationdomain.PaymentUser) error
}

type PaymentUserFilter struct {
	ID                 mo.Option[idtype.PaymentUserID]
	UserID             mo.Option[idtype.UserID]
	RequireDisplayName bool
}

func (f *PaymentUserFilter) IsUniqueFilter() bool {
	return f.UserID.IsPresent()
}
