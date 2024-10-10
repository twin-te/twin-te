package donationusecase

import (
	"context"
	"sync"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	donationmodule "github.com/twin-te/twin-te/back/module/donation"
	donationappdto "github.com/twin-te/twin-te/back/module/donation/appdto"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"golang.org/x/sync/errgroup"
)

func (uc *impl) GetOrCreatePaymentUser(ctx context.Context) (*donationdomain.PaymentUser, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUserOption, err := uc.r.FindPaymentUser(ctx, donationport.PaymentUserFilter{
		UserID: mo.Some(userID),
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	return base.OptionOrElseByWithErr(paymentUserOption, func() (*donationdomain.PaymentUser, error) {
		paymentUser, err := uc.f.NewPaymentUser(userID, mo.None[shareddomain.RequiredString](), mo.None[donationdomain.Link]())
		if err != nil {
			return nil, err
		}
		return paymentUser, uc.r.CreatePaymentUsers(ctx, paymentUser)
	})
}

func (uc *impl) UpdateOrCreatePaymentUser(ctx context.Context, in donationmodule.UpdateOrCreatePaymentUserIn) (paymentUser *donationdomain.PaymentUser, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	err = uc.r.Transaction(ctx, func(rtx donationport.Repository) (err error) {
		paymentUserOption, err := rtx.FindPaymentUser(ctx, donationport.PaymentUserFilter{UserID: mo.Some(userID)}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		var found bool
		paymentUser, found = paymentUserOption.Get()
		if found {
			paymentUser.BeforeUpdateHook()
			paymentUser.Update(donationdomain.PaymentUserDataToUpdate{
				DisplayName: in.DisplayName,
				Link:        in.Link,
			})
			return rtx.UpdatePaymentUser(ctx, paymentUser)
		} else {
			paymentUser, err = uc.f.NewPaymentUser(
				userID,
				in.DisplayName.OrElse(mo.None[shareddomain.RequiredString]()),
				in.Link.OrElse(mo.None[donationdomain.Link]()),
			)
			if err != nil {
				return
			}
			return uc.r.CreatePaymentUsers(ctx, paymentUser)
		}
	}, false)
	return paymentUser, err
}

func (uc *impl) ListContributors(ctx context.Context) ([]donationappdto.Contributor, error) {
	uc.contributorsCacheMutex.RLock()
	defer uc.contributorsCacheMutex.RUnlock()

	return base.CopySlice(uc.contributorsCache), nil
}

func (uc *impl) updateContributorsCache(ctx context.Context) error {
	paymentUsers, err := uc.r.ListPaymentUsers(ctx, donationport.PaymentUserFilter{
		RequireDisplayName: true,
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return err
	}

	paymentUserIDToIsContributor := make(map[idtype.PaymentUserID]bool, len(paymentUsers))
	var paymentUserIDToIsContributorMutex sync.Mutex

	eg, ctx := errgroup.WithContext(ctx)
	for _, paymentUser := range paymentUsers {
		paymentUser := paymentUser
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				paymentHistories, err := uc.i.ListPaymentHistories(ctx, mo.Some(paymentUser.ID))
				if err != nil {
					return err
				}

				paymentUserIDToIsContributorMutex.Lock()
				paymentUserIDToIsContributor[paymentUser.ID] = lo.SomeBy(paymentHistories, func(paymentHistory *donationdomain.PaymentHistory) bool {
					return paymentHistory.Status == donationdomain.PaymentStatusSucceeded && paymentHistory.Amount > 0
				})
				paymentUserIDToIsContributorMutex.Unlock()
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	paymentUsers = lo.Filter(paymentUsers, func(paymentUser *donationdomain.PaymentUser, index int) bool {
		return paymentUserIDToIsContributor[paymentUser.ID]
	})

	contributors := base.Map(paymentUsers, func(paymentUser *donationdomain.PaymentUser) donationappdto.Contributor {
		return donationappdto.Contributor{
			DisplayName: paymentUser.DisplayName.MustGet(),
			Link:        paymentUser.Link,
		}
	})

	uc.contributorsCacheMutex.Lock()
	uc.contributorsCache = contributors
	uc.contributorsCacheMutex.Unlock()

	return nil
}
