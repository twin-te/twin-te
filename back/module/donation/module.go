package donationmodule

import (
	"context"

	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
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
	// CreateOneTimeCheckoutSession creates one-time checkout session.
	//
	// [Authentication] optional
	CreateOneTimeCheckoutSession(ctx context.Context, amount int) (idtype.CheckoutSessionID, error)

	// CreateSubscriptionCheckoutSession creates subscription checkout session.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - donation.ActiveSubscriptionAlreadyExists
	//   - donation.SubscriptionPlanNotFound
	CreateSubscriptionCheckoutSession(ctx context.Context, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error)

	// GetOrCreatePaymentUser returns the payment user.
	// If one does not exists, a new payment user will be created.
	//
	// [Authentication] required
	GetOrCreatePaymentUser(ctx context.Context) (*donationdomain.PaymentUser, error)

	// UpdateOrCreatePaymentUser updates payment user, if one exists.
	// If one does not exist, a new payment user will be created based on the given information.
	//
	// [Authentication] required
	UpdateOrCreatePaymentUser(ctx context.Context, in UpdateOrCreatePaymentUserIn) (*donationdomain.PaymentUser, error)

	// GetPaymentHistories returns the payment histories.
	//
	// [Authentication] required
	GetPaymentHistories(ctx context.Context) ([]*donationdomain.PaymentHistory, error)

	// GetSubscriptionPlans returns the subscription plans.
	//
	// [Authentication] not required
	GetSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error)

	// GetActiveSubscription returns the subscription which is active and has plan association loaded.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - donation.SubscriptionNotFound
	GetActiveSubscription(ctx context.Context) (*donationdomain.Subscription, error)

	// Unsubscribe unsubscribes the subscription specified by the given id.
	//
	// [Authentication] required
	//
	// [Error Code]
	//   - donation.SubscriptionNotFound
	Unsubscribe(ctx context.Context, id idtype.SubscriptionID) error

	// GetTotalAmount returns the total amount donated.
	//
	// [Authentication] not required
	GetTotalAmount(ctx context.Context) (int, error)

	// GetContributors returns contributors.
	// Contributor is payment user who has made at least one donation and has registered name for display.
	//
	// [Authentication] not required
	GetContributors(ctx context.Context) ([]Contributor, error)
}

type UpdateOrCreatePaymentUserIn struct {
	DisplayName *shareddomain.RequiredString
	Link        *donationdomain.Link
}

type Contributor struct {
	DisplayName shareddomain.RequiredString
	Link        *donationdomain.Link
}
