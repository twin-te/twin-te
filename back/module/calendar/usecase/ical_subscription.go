package calendarusecase

import (
	"context"
	"fmt"

	"github.com/twin-te/twin-te/back/appenv"
	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

const icalSubscriptionPath = "/calendar/v1beta/timetable.ics"

func (uc *impl) GetIcalSubscriptionUrl(ctx context.Context) (url string, ok bool, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return "", false, err
	}

	subID, err := uc.r.FindIcalSubscriptionByUserID(ctx, userID, sharedport.LockNone)
	if err != nil {
		return "", false, err
	}
	if subID.IsAbsent() {
		return "", false, nil
	}
	id, _ := subID.Get()
	return buildIcalSubscriptionUrl(id), true, nil
}

func (uc *impl) EnableIcalSubscription(ctx context.Context) (url string, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return "", err
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
		return "", err
	}
	return buildIcalSubscriptionUrl(resultID), nil
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

func buildIcalSubscriptionUrl(id idtype.IcalSubscriptionID) string {
	return fmt.Sprintf("%s%s?token=%s", appenv.APP_URL, icalSubscriptionPath, id.String())
}
