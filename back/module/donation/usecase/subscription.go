package donationusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/twin-te/twin-te/back/apperr"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationerr "github.com/twin-te/twin-te/back/module/donation/err"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) GetSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error) {
	return uc.g.ListSubscriptionPlans(ctx)
}

func (uc *impl) GetSubscription(ctx context.Context) (*donationdomain.Subscription, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return nil, err
	}

	subscription, err := uc.g.FindSubscription(ctx, paymentUser.ID)
	if errors.Is(err, sharedport.ErrNotFound) {
		return nil, apperr.New(
			donationerr.CodeSubscriptionNotFound,
			fmt.Sprintf("not found subscription associated with the user whose id is %s", userID),
		)
	}

	return subscription, err
}

func (uc *impl) Unsubscribe(ctx context.Context, id idtype.SubscriptionID) error {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	subscription, err := uc.GetSubscription(ctx)
	if err != nil {
		return err
	}

	if id != subscription.ID {
		return apperr.New(
			donationerr.CodeSubscriptionNotFound,
			fmt.Sprintf("not found subscription whose id is %s", id),
		)
	}

	return uc.g.DeleteSubscription(ctx, id)
}
