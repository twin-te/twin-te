package donationgateway

import (
	"context"
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
	"github.com/twin-te/twinte-back/base"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

func (g *impl) ListSubscriptions(ctx context.Context, paymentUserID idtype.PaymentUserID) ([]*donationdomain.Subscription, error) {
	var startingAfter *string

	subscriptions := make([]*stripe.Subscription, 0)

	for {
		iter := subscription.List(&stripe.SubscriptionListParams{
			ListParams: stripe.ListParams{
				Context:       ctx,
				Limit:         stripe.Int64(100),
				StartingAfter: startingAfter,
			},
			Customer: stripe.String(paymentUserID.String()),
		})

		if err := iter.Err(); err != nil {
			return nil, err
		}

		data := iter.SubscriptionList().Data

		subscriptions = append(subscriptions, data...)

		if !iter.Meta().HasMore {
			break
		}

		startingAfter = &data[len(data)-1].ID

		time.Sleep(50 * time.Microsecond)
	}

	return base.MapWithErr(subscriptions, fromStripeSubscription)
}

func (g *impl) DeleteSubscription(ctx context.Context, id idtype.SubscriptionID) (err error) {
	params := &stripe.SubscriptionCancelParams{
		Params: stripe.Params{
			Context: ctx,
		},
	}

	stripeSubscription, err := subscription.Cancel(id.String(), params)
	if err != nil {
		return
	}

	if stripeSubscription.Status != stripe.SubscriptionStatusCanceled {
		return fmt.Errorf("failed to cancel subscription whose id is %s", id)
	}

	return
}

func fromStripeSubscription(stripeSubscription *stripe.Subscription) (*donationdomain.Subscription, error) {
	return donationdomain.ConstructSubscription(func(s *donationdomain.Subscription) (err error) {
		s.ID, err = idtype.ParseSubscriptionID(stripeSubscription.ID)
		if err != nil {
			return
		}

		s.PaymentUserID, err = idtype.ParsePaymentUserID(stripeSubscription.Customer.ID)
		if err != nil {
			return
		}

		s.Plans, err = base.MapWithErr(stripeSubscription.Items.Data, func(item *stripe.SubscriptionItem) (*donationdomain.SubscriptionPlan, error) {
			return donationdomain.ConstructSubscriptionPlan(func(sp *donationdomain.SubscriptionPlan) (err error) {
				sp.ID, err = idtype.ParseSubscriptionPlanID(item.Plan.ID)
				if err != nil {
					return
				}

				sp.Name = item.Plan.Nickname
				sp.Amount = int(item.Plan.Amount)

				return
			})
		})
		if err != nil {
			return
		}

		s.IsActive = stripeSubscription.Status == stripe.SubscriptionStatusActive
		s.CreatedAt = time.Unix(stripeSubscription.Created, 0)

		return nil
	})
}
