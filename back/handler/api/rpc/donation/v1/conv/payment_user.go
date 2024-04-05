package donationv1conv

import (
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	donationv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/donation/v1"
	donationmodule "github.com/twin-te/twin-te/back/module/donation"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
)

func ToPBPaymentUser(paymentUser *donationdomain.PaymentUser) *donationv1.PaymentUser {
	return &donationv1.PaymentUser{
		Id:          paymentUser.ID.String(),
		UserId:      sharedconv.ToPBUUID(paymentUser.UserID),
		DisplayName: paymentUser.DisplayName.StringPtr(),
		Link:        paymentUser.Link.StringPtr(),
	}
}

func ToPBContributor(Contributor donationmodule.Contributor) *donationv1.Contributor {
	return &donationv1.Contributor{
		DisplayName: Contributor.DisplayName.String(),
		Link:        Contributor.Link.StringPtr(),
	}
}
