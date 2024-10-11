package donationintegrator

import (
	"log"

	"github.com/stripe/stripe-go/v76"
	"github.com/twin-te/twin-te/back/appenv"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
)

func init() {
	stripe.Key = appenv.STRIPE_KEY
}

var _ donationport.Integrator = (*impl)(nil)

type impl struct {
	subscriptionPlans []*donationdomain.SubscriptionPlan
}

func New() *impl {
	g := &impl{}

	if err := g.loadSubscriptionPlans(); err != nil {
		log.Printf("failed to load subscription plans, %+v", err)
	}

	return g
}
