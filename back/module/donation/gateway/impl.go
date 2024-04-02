package donationgateway

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/twin-te/twinte-back/appenv"
	donationport "github.com/twin-te/twinte-back/module/donation/port"
)

func init() {
	stripe.Key = appenv.STRIPE_KEY
}

var _ donationport.Gateway = (*impl)(nil)

type impl struct{}

func New() *impl {
	return &impl{}
}
