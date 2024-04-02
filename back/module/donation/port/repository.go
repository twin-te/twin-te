package donationport

import (
	"context"

	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error) error

	// PaymentUser

	FindPaymentUser(ctx context.Context, conds FindPaymentUserConds, lock sharedport.Lock) (*donationdomain.PaymentUser, error)
	ListPaymentUsers(ctx context.Context, conds ListPaymentUsersConds, lock sharedport.Lock) ([]*donationdomain.PaymentUser, error)
	CreatePaymentUsers(ctx context.Context, paymentUsers ...*donationdomain.PaymentUser) error
	UpdatePaymentUser(ctx context.Context, paymentUser *donationdomain.PaymentUser) error
}

type FindPaymentUserConds struct {
	UserID idtype.UserID
}

type ListPaymentUsersConds struct {
	RequireDisplayName bool
}
