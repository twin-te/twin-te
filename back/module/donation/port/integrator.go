package donationport

import (
	"context"

	"github.com/samber/mo"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Integrator interface {
	// Stripe

	CreateOneTimeCheckoutSession(ctx context.Context, paymentUserID mo.Option[idtype.PaymentUserID], amount int) (idtype.CheckoutSessionID, error)
	CreateSubscriptionCheckoutSession(ctx context.Context, paymentUserID idtype.PaymentUserID, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error)
	ListPaymentHistories(ctx context.Context, paymentUserID mo.Option[idtype.PaymentUserID]) ([]*donationdomain.PaymentHistory, error)

	FindSubscriptionPlan(ctx context.Context, id idtype.SubscriptionPlanID) (mo.Option[*donationdomain.SubscriptionPlan], error)
	ListSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error)

	// ListSubscriptions returns the subscriptions associated with the given payment user id.
	// The each returned subscription has plan association loaded.
	ListSubscriptions(ctx context.Context, paymentUserID idtype.PaymentUserID) ([]*donationdomain.Subscription, error)

	DeleteSubscription(ctx context.Context, id idtype.SubscriptionID) (err error)
}
