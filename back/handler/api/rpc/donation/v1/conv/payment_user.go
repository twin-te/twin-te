package donationv1conv

import (
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	donationv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/donation/v1"
	donationappdto "github.com/twin-te/twin-te/back/module/donation/appdto"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
)

func ToPBPaymentUser(paymentUser *donationdomain.PaymentUser) *donationv1.PaymentUser {
	return &donationv1.PaymentUser{
		Id:          paymentUser.ID.String(),
		UserId:      sharedconv.ToPBUUID(paymentUser.UserID),
		DisplayName: base.OptionMapByString(paymentUser.DisplayName).ToPointer(),
		Link:        base.OptionMapByString(paymentUser.Link).ToPointer(),
	}
}

func ToPBContributor(Contributor donationappdto.Contributor) *donationv1.Contributor {
	return &donationv1.Contributor{
		DisplayName: Contributor.DisplayName.String(),
		Link:        base.OptionMapByString(Contributor.Link).ToPointer(),
	}
}
