package donationport

import (
	"context"

	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

type Gateway interface {
	// Stripe

	CreateOneTimeCheckoutSession(ctx context.Context, paymentUserID *idtype.PaymentUserID, amount int) (idtype.CheckoutSessionID, error)
	CreateSubscriptionCheckoutSession(ctx context.Context, paymentUserID idtype.PaymentUserID, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error)
	ListPaymentHistories(ctx context.Context, paymentUserID *idtype.PaymentUserID) ([]*donationdomain.PaymentHistory, error)
	ListSubscriptions(ctx context.Context, paymentUserID idtype.PaymentUserID) ([]*donationdomain.Subscription, error)
	DeleteSubscription(ctx context.Context, id idtype.SubscriptionID) (err error)
}
