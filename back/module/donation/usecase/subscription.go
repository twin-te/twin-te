package donationusecase

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/apperr"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationerr "github.com/twin-te/twin-te/back/module/donation/err"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func (uc *impl) ListSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error) {
	return uc.i.ListSubscriptionPlans(ctx)
}

func (uc *impl) GetActiveSubscription(ctx context.Context) (*donationdomain.Subscription, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return nil, err
	}

	subscriptions, err := uc.i.ListSubscriptions(ctx, paymentUser.ID)
	if err != nil {
		return nil, err
	}

	activeSubscription, ok := lo.Find(subscriptions, func(subscription *donationdomain.Subscription) bool {
		return subscription.IsActive
	})
	if ok {
		return activeSubscription, nil
	}

	return nil, apperr.New(
		donationerr.CodeSubscriptionNotFound,
		fmt.Sprintf("not found active subscription associated with the user whose id is %s", userID),
	)
}

func (uc *impl) Unsubscribe(ctx context.Context, id idtype.SubscriptionID) error {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return err
	}

	subscriptions, err := uc.i.ListSubscriptions(ctx, paymentUser.ID)
	if err != nil {
		return err
	}

	if !lo.ContainsBy(subscriptions, func(subscription *donationdomain.Subscription) bool {
		return subscription.ID == id
	}) {
		return apperr.New(
			donationerr.CodeSubscriptionNotFound,
			fmt.Sprintf("not found subscription whose id is %s", id),
		)
	}

	return uc.i.DeleteSubscription(ctx, id)
}
