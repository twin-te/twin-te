package calendarusecase

import (
	"context"

	"github.com/samber/mo"
	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) GetIcalSubscriptionID(ctx context.Context) (mo.Option[idtype.IcalSubscriptionID], error) {
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

func (uc *impl) ResolveUserIDByIcalSubscriptionID(ctx context.Context, id idtype.IcalSubscriptionID) (idtype.UserID, bool, error) {
	userID, err := uc.r.FindIcalSubscriptionByID(ctx, id, sharedport.LockNone)
	if err != nil {
		return idtype.UserID{}, false, err
	}
	if userID.IsAbsent() {
		return idtype.UserID{}, false, nil
	}
	uid, _ := userID.Get()
	return uid, true, nil
}

