package donationdbmodel

import (
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type PaymentUser struct {
	ID          string
	UserID      string
	DisplayName mo.Option[string]
	Link        mo.Option[string]
}

func FromDBPaymentUser(dbPaymentUser *PaymentUser) (*donationdomain.PaymentUser, error) {
	return donationdomain.ConstructPaymentUser(func(pu *donationdomain.PaymentUser) (err error) {
		pu.ID, err = idtype.ParsePaymentUserID(dbPaymentUser.ID)
		if err != nil {
			return
		}

		pu.UserID, err = idtype.ParseUserID(dbPaymentUser.UserID)
		if err != nil {
			return
		}

		if displayName, ok := dbPaymentUser.DisplayName.Get(); ok {
			pu.DisplayName, err = base.ToPtrWithErr(donationdomain.ParseDisplayName(displayName))
			if err != nil {
				return
			}
		}

		if link, ok := dbPaymentUser.Link.Get(); ok {
			pu.Link, err = base.ToPtrWithErr(donationdomain.ParseLink(link))
			if err != nil {
				return
			}
		}

		return
	})
}

func ToDBPaymentUser(paymentUser *donationdomain.PaymentUser) *PaymentUser {
	return &PaymentUser{
		ID:          paymentUser.ID.String(),
		UserID:      paymentUser.UserID.String(),
		DisplayName: mo.PointerToOption(paymentUser.DisplayName.StringPtr()),
		Link:        mo.PointerToOption(paymentUser.Link.StringPtr()),
	}
}
