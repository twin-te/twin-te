package donationfactory

import (
	"github.com/samber/mo"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var _ donationport.Factory = (*impl)(nil)

type impl struct{}

func (f *impl) NewPaymentUser(userID idtype.UserID, displayName mo.Option[shareddomain.RequiredString], link mo.Option[donationdomain.Link]) (*donationdomain.PaymentUser, error) {
	return donationdomain.ConstructPaymentUser(func(pu *donationdomain.PaymentUser) (err error) {
		customer, err := customer.New(&stripe.CustomerParams{})
		if err != nil {
			return
		}

		pu.ID, err = idtype.ParsePaymentUserID(customer.ID)
		if err != nil {
			return
		}

		pu.UserID = userID
		pu.DisplayName = displayName
		pu.Link = link

		return
	})
}

func New() *impl {
	return &impl{}
}
