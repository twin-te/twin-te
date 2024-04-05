package donationport

import (
	"context"

	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Gateway interface {
	// Stripe

	CreateOneTimeCheckoutSession(ctx context.Context, paymentUserID *idtype.PaymentUserID, amount int) (idtype.CheckoutSessionID, error)
	CreateSubscriptionCheckoutSession(ctx context.Context, paymentUserID idtype.PaymentUserID, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error)
	ListPaymentHistories(ctx context.Context, paymentUserID *idtype.PaymentUserID) ([]*donationdomain.PaymentHistory, error)
	ListSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error)

	// FindSubscription returns the subscription associated with the given payment user id.
	// If it does not exist, ErrNotFound is returned.
	// The returned subscription has plan association loaded.
	FindSubscription(ctx context.Context, paymentUserID idtype.PaymentUserID) (*donationdomain.Subscription, error)

	DeleteSubscription(ctx context.Context, id idtype.SubscriptionID) (err error)
}
