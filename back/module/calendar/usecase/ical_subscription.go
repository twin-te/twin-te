package calendarusecase

import (
	"context"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/apperr"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) FindIcalSubscription(ctx context.Context) (mo.Option[*calendardomain.IcalSubscription], error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return mo.None[*calendardomain.IcalSubscription](), err
	}

	return uc.r.FindIcalSubscriptionByUserID(ctx, userID, sharedport.LockNone)
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
		if sub, ok := existing.Get(); ok {
			resultID = sub.ID
			return nil
		}

		sub, err := calendardomain.ConstructIcalSubscription(func(s *calendardomain.IcalSubscription) error {
			s.ID = idtype.NewIcalSubscriptionID()
			s.UserID = userID
			s.Mode = calendardomain.IcalSubscriptionModeSync
			return nil
		})
		if err != nil {
			return err
		}
		resultID = sub.ID
		return rtx.CreateIcalSubscription(ctx, sub)
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

func (uc *impl) UpdateIcalSubscription(ctx context.Context, mode calendardomain.IcalSubscriptionMode, targetTagIDs []idtype.TagID) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	// トランザクションを開始する前に、対象タグ ID がすべてユーザー所有であることを検証する。
	if len(targetTagIDs) != 0 {
		ownedTagIDs, err := uc.timetable.ListTagIDsByUserID(ctx, userID, targetTagIDs)
		if err != nil {
			return err
		}
		if notOwned, _ := lo.Difference(targetTagIDs, ownedTagIDs); len(notOwned) != 0 {
			return sharederr.NewInvalidArgument("invalid tag ids %+v", notOwned)
		}
	}

	return uc.r.Transaction(ctx, func(rtx calendarport.Repository) error {
		existing, err := rtx.FindIcalSubscriptionByUserID(ctx, userID, sharedport.LockExclusive)
		if err != nil {
			return err
		}
		sub, ok := existing.Get()
		if !ok {
			return apperr.New(sharederr.CodeInvalidArgument, "ical subscription is not enabled")
		}

		updated, err := calendardomain.ConstructIcalSubscription(func(s *calendardomain.IcalSubscription) error {
			s.ID = sub.ID
			s.UserID = sub.UserID
			s.Mode = mode
			s.TargetTagIDs = targetTagIDs
			return nil
		})
		if err != nil {
			return err
		}

		return rtx.UpdateIcalSubscription(ctx, updated)
	}, false)
}

func (uc *impl) resolveIcalSubscriptionByID(ctx context.Context, id idtype.IcalSubscriptionID) (*calendardomain.IcalSubscription, error) {
	optSub, err := uc.r.FindIcalSubscriptionByID(ctx, id, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	sub, ok := optSub.Get()
	if !ok {
		return nil, apperr.New(sharederr.CodeUnauthenticated, "invalid token")
	}

	return sub, nil
}
