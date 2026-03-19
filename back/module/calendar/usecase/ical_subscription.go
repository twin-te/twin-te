package calendarusecase

import (
	"context"
	"fmt"

	"github.com/samber/mo"
	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) FindIcalSubscriptionID(ctx context.Context) (mo.Option[idtype.IcalSubscriptionID], error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return mo.None[idtype.IcalSubscriptionID](), err
	}

	subID, err := uc.r.FindIcalSubscriptionByUserID(ctx, userID, sharedport.LockNone)
	if err != nil {
		return mo.None[idtype.IcalSubscriptionID](), err
	}

	return subID, nil
}

func (uc *impl) EnableIcalSubscription(ctx context.Context) (idtype.IcalSubscriptionID, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return idtype.IcalSubscriptionID{}, err
	}

	var resultID idtype.IcalSubscriptionID
	err = uc.r.Transaction(ctx, func(rtx calendarport.Repository) error {
		existing, err := rtx.FindIcalSubscriptionByUserID(ctx, userID, sharedport.LockExclusive)
		if err != nil {
			return err
		}
		if existing.IsPresent() {
			resultID, _ = existing.Get()
			return nil
		}
		resultID = idtype.NewIcalSubscriptionID()
		return rtx.CreateIcalSubscription(ctx, resultID, userID)
	}, false)
	if err != nil {
		return idtype.IcalSubscriptionID{}, err
	}
	return resultID, nil
}

func (uc *impl) DisableIcalSubscription(ctx context.Context) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}
	_, err = uc.r.DeleteIcalSubscriptionByUserID(ctx, userID)
	return err
}

func (uc *impl) ResolveUserIDByIcalSubscriptionID(ctx context.Context, id idtype.IcalSubscriptionID) (idtype.UserID, error) {
	optUserID, err := uc.r.FindUserByIcalSubscriptionID(ctx, id, sharedport.LockNone)
	if err != nil {
		return idtype.UserID{}, err
	}

	userID, ok := optUserID.Get()
	if !ok {
		return idtype.UserID{}, fmt.Errorf("invalid token")
	}

	return userID, nil
}
