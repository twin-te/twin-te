package donationgateway

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/plan"
	"github.com/stripe/stripe-go/v76/subscription"
	"github.com/twin-te/twin-te/back/base"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (g *impl) ListSubscriptions(ctx context.Context, paymentUserID idtype.PaymentUserID) ([]*donationdomain.Subscription, error) {
	var startingAfter *string

	stripeSubscriptions := make([]*stripe.Subscription, 0)

	for {
		iter := subscription.List(&stripe.SubscriptionListParams{
			ListParams: stripe.ListParams{
				Context:       ctx,
				Limit:         stripe.Int64(100),
				StartingAfter: startingAfter,
			},
			Customer: stripe.String(paymentUserID.String()),
			Status:   stripe.String("all"),
		})

		if err := iter.Err(); err != nil {
			return nil, err
		}

		data := iter.SubscriptionList().Data

		stripeSubscriptions = append(stripeSubscriptions, data...)

		if !iter.Meta().HasMore {
			break
		}

		startingAfter = &data[len(data)-1].ID

		time.Sleep(50 * time.Microsecond)
	}

	return base.MapWithErr(stripeSubscriptions, fromStripeSubscription)
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

func (g *impl) FindSubscriptionPlan(ctx context.Context, id idtype.SubscriptionPlanID) (*donationdomain.SubscriptionPlan, error) {
	subscriptionPlan, ok := lo.Find(g.subscriptionPlans, func(plan *donationdomain.SubscriptionPlan) bool {
		return plan.ID == id
	})

	if ok {
		return subscriptionPlan, nil
	}

	return nil, sharedport.ErrNotFound
}

func (g *impl) ListSubscriptionPlans(ctx context.Context) ([]*donationdomain.SubscriptionPlan, error) {
	return base.MapByClone(g.subscriptionPlans), nil
}

func (g *impl) loadSubscriptionPlans() (err error) {
	var startingAfter *string

	stripePlans := make([]*stripe.Plan, 0)

	for {
		iter := plan.List(&stripe.PlanListParams{
			ListParams: stripe.ListParams{
				Limit:         stripe.Int64(100),
				StartingAfter: startingAfter,
			},
			Active: stripe.Bool(true),
		})

		if err := iter.Err(); err != nil {
			return err
		}

		data := iter.PlanList().Data

		stripePlans = append(stripePlans, data...)

		if !iter.Meta().HasMore {
			break
		}

		startingAfter = &data[len(data)-1].ID

		time.Sleep(50 * time.Microsecond)
	}

	g.subscriptionPlans, err = base.MapWithErr(stripePlans, fromStripePlan)
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

		if len(stripeSubscription.Items.Data) != 1 {
			return fmt.Errorf("subscription (%s) must have only one plan, but got %+v", stripeSubscription.ID, stripeSubscription.Items.Data)
		}

		plan, err := fromStripePlan(stripeSubscription.Items.Data[0].Plan)
		if err != nil {
			return
		}
		s.PlanID = plan.ID
		s.PlanAssociation.Set(plan)

		s.IsActive = stripeSubscription.Status == stripe.SubscriptionStatusActive
		s.CreatedAt = time.Unix(stripeSubscription.Created, 0)

		return nil
	})
}

func fromStripePlan(stripePlan *stripe.Plan) (*donationdomain.SubscriptionPlan, error) {
	return donationdomain.ConstructSubscriptionPlan(func(sp *donationdomain.SubscriptionPlan) (err error) {
		sp.ID, err = idtype.ParseSubscriptionPlanID(stripePlan.ID)
		if err != nil {
			return
		}

		sp.Name = stripePlan.Nickname
		sp.Amount = int(stripePlan.Amount)

		return
	})
}
