package donationusecase

import (
	"context"
	"fmt"

	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

func (uc *impl) CreateOneTimeCheckoutSession(ctx context.Context, amount int) (idtype.CheckoutSessionID, error) {
	if amount <= 0 {
		return "", fmt.Errorf("amount must be greater than 0, but got %d", amount)
	}

	var paymentUserID *idtype.PaymentUserID

	_, err := uc.a.Authenticate(ctx)
	if err == nil {
		paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
		if err != nil {
			return "", err
		}
		paymentUserID = &paymentUser.ID
	}

	return uc.g.CreateOneTimeCheckoutSession(ctx, paymentUserID, amount)
}

func (uc *impl) CreateSubscriptionCheckoutSession(ctx context.Context, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error) {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return "", err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return "", nil
	}

	return uc.g.CreateSubscriptionCheckoutSession(ctx, paymentUser.ID, subscriptionPlanID)
}
