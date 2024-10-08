package donationusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/apperr"
	donationerr "github.com/twin-te/twin-te/back/module/donation/err"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) CreateOneTimeCheckoutSession(ctx context.Context, amount int) (idtype.CheckoutSessionID, error) {
	if amount <= 0 {
		return "", sharederr.NewInvalidArgument("amount must be greater than 0, but got %d", amount)
	}

	var paymentUserID mo.Option[idtype.PaymentUserID]

	_, err := uc.a.Authenticate(ctx)
	if err == nil {
		paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
		if err != nil {
			return "", err
		}
		paymentUserID = mo.Some(paymentUser.ID)
	}

	return uc.i.CreateOneTimeCheckoutSession(ctx, paymentUserID, amount)
}

func (uc *impl) CreateSubscriptionCheckoutSession(ctx context.Context, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return "", err
	}

	_, err = uc.i.FindSubscriptionPlan(ctx, subscriptionPlanID)
	if err != nil {
		if errors.Is(err, sharedport.ErrNotFound) {
			return "", apperr.New(
				donationerr.CodeSubscriptionPlanNotFound,
				fmt.Sprintf("not found subscription plan (%s)", subscriptionPlanID),
			)
		}
		return "", err
	}

	activeSubscription, err := uc.GetActiveSubscription(ctx)
	switch {
	case err == nil:
		return "", apperr.New(
			donationerr.CodeActiveSubscriptionAlreadyExists,
			fmt.Sprintf("user (%s) has already active subscription (%s)", userID, activeSubscription.ID),
		)
	case !apperr.Is(err, donationerr.CodeSubscriptionNotFound):
		return "", err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return "", nil
	}

	return uc.i.CreateSubscriptionCheckoutSession(ctx, paymentUser.ID, subscriptionPlanID)
}
