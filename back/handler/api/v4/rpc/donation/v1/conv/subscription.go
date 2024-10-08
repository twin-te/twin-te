package donationv1conv

import (
	sharedconv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/shared/conv"
	donationv1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/donation/v1"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
)

func ToPBSubscription(subscription *donationdomain.Subscription) *donationv1.Subscription {
	return &donationv1.Subscription{
		Id:        subscription.ID.String(),
		Plan:      ToPBSubscriptionPlan(subscription.PlanAssociation.MustGet()),
		IsActive:  subscription.IsActive,
		CreatedAt: sharedconv.ToPBRFC3339DateTime(subscription.CreatedAt),
	}
}

func ToPBSubscriptionPlan(subscriptionPlan *donationdomain.SubscriptionPlan) *donationv1.SubscriptionPlan {
	return &donationv1.SubscriptionPlan{
		Id:     subscriptionPlan.ID.String(),
		Name:   subscriptionPlan.Name,
		Amount: int32(subscriptionPlan.Amount),
	}
}
