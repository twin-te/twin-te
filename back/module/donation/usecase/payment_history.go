package donationusecase

import (
	"context"

	"github.com/samber/lo"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
)

func (uc *impl) GetPaymentHistories(ctx context.Context) ([]*donationdomain.PaymentHistory, error) {
	_, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	paymentUser, err := uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return nil, err
	}

	return uc.g.ListPaymentHistories(ctx, &paymentUser.ID)
}

func (uc *impl) GetTotalAmount(ctx context.Context) (int, error) {
	uc.totalAmountCacheMutex.RLock()
	defer uc.totalAmountCacheMutex.RUnlock()

	return uc.totalAmountCache, nil
}

func (uc *impl) updateTotalAmountCache(ctx context.Context) error {
	paymentHistories, err := uc.g.ListPaymentHistories(ctx, nil)
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
