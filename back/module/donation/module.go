package donationmodule

import (
	"context"

	"github.com/samber/mo"
	donationappdto "github.com/twin-te/twin-te/back/module/donation/appdto"
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

	// ListPaymentHistories returns the payment histories.
	//
	// [Authentication] required
	ListPaymentHistories(ctx context.Context) ([]*donationdomain.PaymentHistory, error)

	// ListSubscriptionPlans returns the subscription plans.
	//
	// [Authentication] not required
	ListSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error)

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

	// ListContributors returns contributors.
	// Contributor is payment user who has made at least one donation and has registered name for display.
	//
	// [Authentication] not required
	ListContributors(ctx context.Context) ([]donationappdto.Contributor, error)
}

type UpdateOrCreatePaymentUserIn struct {
	DisplayName mo.Option[mo.Option[shareddomain.RequiredString]]
	Link        mo.Option[mo.Option[donationdomain.Link]]
}
