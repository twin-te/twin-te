package donationfactory

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	donationport "github.com/twin-te/twinte-back/module/donation/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

var _ donationport.Factory = (*impl)(nil)

type impl struct{}

func (f *impl) NewPaymentUser(userID idtype.UserID, displayName *string, link *string) (*donationdomain.PaymentUser, error) {
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
