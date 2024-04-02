package donationv1conv

import (
	"fmt"

	sharedconv "github.com/twin-te/twinte-back/handler/api/rpc/shared/conv"
	donationv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/donation/v1"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
)

func ToPBPaymentType(paymentType donationdomain.PaymentType) donationv1.PaymentType {
	switch paymentType {
	case donationdomain.PaymentTypeOneTime:
		return donationv1.PaymentType_PAYMENT_TYPE_ONE_TIME
	case donationdomain.PaymentTypeSubscription:
		return donationv1.PaymentType_PAYMENT_TYPE_SUBSCRIPTION
	}
	panic(fmt.Sprintf("never happened %#v", paymentType))
}

func ToPBPaymentStatus(paymentStatus donationdomain.PaymentStatus) donationv1.PaymentStatus {
	switch paymentStatus {
	case donationdomain.PaymentStatusPending:
		return donationv1.PaymentStatus_PAYMENT_STATUS_PENDING
	case donationdomain.PaymentStatusCanceled:
		return donationv1.PaymentStatus_PAYMENT_STATUS_CANCELED
	case donationdomain.PaymentStatusSucceeded:
		return donationv1.PaymentStatus_PAYMENT_STATUS_SUCCEEDED
	}
	panic(fmt.Sprintf("never happened %#v", paymentStatus))
}

func ToPBPaymentHistory(paymentHistory *donationdomain.PaymentHistory) *donationv1.PaymentHistory {
	return &donationv1.PaymentHistory{
		Id:        paymentHistory.ID.String(),
		Type:      ToPBPaymentType(paymentHistory.Type),
		Status:    ToPBPaymentStatus(paymentHistory.Status),
		Amount:    int32(paymentHistory.Amount),
		CreatedAt: sharedconv.ToPBRFC3339DateTime(paymentHistory.CreatedAt),
	}
}
