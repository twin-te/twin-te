package donationusecase

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/apperr"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	donationerr "github.com/twin-te/twinte-back/module/donation/err"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

func (uc *impl) GetSubscriptions(ctx context.Context) ([]*donationdomain.Subscription, error) {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return nil, err
	}

	return uc.g.ListSubscriptions(ctx, paymentUser.ID)
}

func (uc *impl) Unsubscribe(ctx context.Context, id idtype.SubscriptionID) error {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	subscriptions, err := uc.GetSubscriptions(ctx)
	if err != nil {
		return err
	}

	if !lo.ContainsBy(subscriptions, func(item *donationdomain.Subscription) bool {
		return item.ID == id
	}) {
		return apperr.New(
			donationerr.CodeSubscriptionNotFound,
			fmt.Sprintf("not found subscription whose id is %s", id),
		)
	}

	return uc.g.DeleteSubscription(ctx, id)
}
