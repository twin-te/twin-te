package donationgateway

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/twin-te/twinte-back/base"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

func (g *impl) ListPaymentHistories(ctx context.Context, paymentUserID *idtype.PaymentUserID) ([]*donationdomain.PaymentHistory, error) {
	var startingAfter *string

	var customer *string
	if paymentUserID != nil {
		customer = stripe.String(paymentUserID.String())
	}

	paymentIntents := make([]*stripe.PaymentIntent, 0)

	for {
		iter := paymentintent.List(&stripe.PaymentIntentListParams{
			ListParams: stripe.ListParams{
				Context:       ctx,
				Limit:         stripe.Int64(100),
				StartingAfter: startingAfter,
			},
			Customer: customer,
			Expand:   stripe.StringSlice([]string{"data.invoice"}),
		})

		if err := iter.Err(); err != nil {
			return nil, err
		}

		data := iter.PaymentIntentList().Data

		paymentIntents = append(paymentIntents, data...)

		if !iter.Meta().HasMore {
			break
		}

		startingAfter = &data[len(data)-1].ID

		time.Sleep(40 * time.Microsecond)
	}

	return base.MapWithErr(paymentIntents, fromStripePaymentIntent)
}

func fromStripePaymentIntent(paymentIntent *stripe.PaymentIntent) (*donationdomain.PaymentHistory, error) {
	return donationdomain.ConstructPaymentHistory(func(ph *donationdomain.PaymentHistory) (err error) {
		ph.ID, err = idtype.ParsePaymentHistoryID(paymentIntent.ID)
		if err != nil {
			return
		}

		if paymentIntent.Customer != nil {
			ph.PaymentUserID, err = base.ToPtrWithErr(idtype.ParsePaymentUserID(paymentIntent.Customer.ID))
			if err != nil {
				return
			}
		}

		// According to the documentation (https://stripe.com/docs/invoicing/overview), it is stated as below.
		// > サブスクリプションは請求サイクルごとにインボイスを自動生成します。
		//
		// That's why, in the case of subscription, invoice seems to be generated automatically.
		//
		// In the case of one-time payment, invoice doesn't seem to be generated
		// unless stripe.CheckoutSessionParams.InvoiceCreation.Enabled is set to true explicitly in creating CheckoutSession.
		ph.Type = lo.Ternary(
			paymentIntent.Invoice == nil,
			donationdomain.PaymentTypeOneTime,
			donationdomain.PaymentTypeSubscription,
		)

		switch paymentIntent.Status {
		case stripe.PaymentIntentStatusCanceled:
			ph.Status = donationdomain.PaymentStatusCanceled
		case stripe.PaymentIntentStatusSucceeded:
			ph.Status = donationdomain.PaymentStatusSucceeded
		default:
			ph.Status = donationdomain.PaymentStatusPending
		}

		ph.Amount = int(paymentIntent.Amount)
		ph.CreatedAt = time.Unix(paymentIntent.Created, 0)

		return
	})
}
