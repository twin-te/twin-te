/*
	References
		- https://stripe.com/docs/payments/checkout/how-checkout-works?payment-ui=stripe-hosted-page
		- https://stripe.com/docs/checkout/quickstart?lang=go
		- https://stripe.com/docs/billing/subscriptions/build-subscriptions?ui=checkout
*/

package donationintegrator

import (
	"context"
	"fmt"

	"github.com/samber/mo"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/twin-te/twin-te/back/appenv"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func (i *impl) CreateOneTimeCheckoutSession(ctx context.Context, paymentUserID mo.Option[idtype.PaymentUserID], amount int) (idtype.CheckoutSessionID, error) {
	params := &stripe.CheckoutSessionParams{
		Params: stripe.Params{
			Context: ctx,
		},
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		SubmitType:         stripe.String("donate"),
		Customer:           base.OptionMapByString(paymentUserID).ToPointer(),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Quantity: stripe.Int64(1),
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String("Twin:te寄付"),
						Description: stripe.String("寄付いただいたお金はTwin:teの運用や開発に使用します"),
						Images:      stripe.StringSlice([]string{"https://www.twinte.net/ogp.jpg"}),
					},
					Currency:   stripe.String("jpy"),
					UnitAmount: stripe.Int64(int64(amount)),
				},
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(fmt.Sprintf("%s?type=onetime&amount=%d", appenv.STRIPE_CHECKOUT_SUCCESS_URL, amount)),
		CancelURL:  stripe.String(appenv.STRIPE_CHECKOUT_CANCEL_URL),
	}

	s, err := session.New(params)
	if err != nil {
		return "", err
	}

	return idtype.ParseCheckoutSessionID(s.ID)
}

func (i *impl) CreateSubscriptionCheckoutSession(ctx context.Context, paymentUserID idtype.PaymentUserID, subscriptionPlanID idtype.SubscriptionPlanID) (idtype.CheckoutSessionID, error) {
	params := &stripe.CheckoutSessionParams{
		Params: stripe.Params{
			Context: ctx,
		},
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Customer:           stripe.String(paymentUserID.String()),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(subscriptionPlanID.String()),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(fmt.Sprintf("%s?type=subscription&plan_id=%s", appenv.STRIPE_CHECKOUT_SUCCESS_URL, subscriptionPlanID)),
		CancelURL:  stripe.String(appenv.STRIPE_CHECKOUT_CANCEL_URL),
	}

	s, err := session.New(params)
	if err != nil {
		return "", err
	}

	return idtype.ParseCheckoutSessionID(s.ID)
}
