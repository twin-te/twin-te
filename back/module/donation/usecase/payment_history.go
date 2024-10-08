package donationusecase

import (
	"context"

	"github.com/samber/lo"
	"github.com/samber/mo"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func (uc *impl) ListPaymentHistories(ctx context.Context) ([]*donationdomain.PaymentHistory, error) {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return nil, err
	}

	return uc.i.ListPaymentHistories(ctx, mo.Some(paymentUser.ID))
}

func (uc *impl) GetTotalAmount(ctx context.Context) (int, error) {
	uc.totalAmountCacheMutex.RLock()
	defer uc.totalAmountCacheMutex.RUnlock()

	return uc.totalAmountCache, nil
}

func (uc *impl) updateTotalAmountCache(ctx context.Context) error {
	paymentHistories, err := uc.i.ListPaymentHistories(ctx, mo.None[idtype.PaymentUserID]())
	if err != nil {
		return err
	}

	paymentHistories = lo.Filter(paymentHistories, func(paymentHistory *donationdomain.PaymentHistory, _ int) bool {
		return paymentHistory.Status == donationdomain.PaymentStatusSucceeded
	})

	totalAmount := lo.Reduce(paymentHistories, func(totalAmount int, paymentHistory *donationdomain.PaymentHistory, _ int) int {
		return totalAmount + paymentHistory.Amount
	}, 0)

	uc.totalAmountCacheMutex.Lock()
	uc.totalAmountCache = totalAmount
	uc.totalAmountCacheMutex.Unlock()

	return nil
}
