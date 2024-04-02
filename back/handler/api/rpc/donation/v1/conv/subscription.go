package donationv1conv

import (
	"github.com/twin-te/twinte-back/base"
	sharedconv "github.com/twin-te/twinte-back/handler/api/rpc/shared/conv"
	donationv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/donation/v1"
	donationdomain "github.com/twin-te/twinte-back/module/donation/domain"
)

func ToPBSubscription(subscription *donationdomain.Subscription) *donationv1.Subscription {
	return &donationv1.Subscription{
		Id:        subscription.ID.String(),
		Plans:     base.Map(subscription.Plans, ToPBSubscriptionPlan),
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
