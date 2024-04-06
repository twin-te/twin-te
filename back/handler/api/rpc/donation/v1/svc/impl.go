package donationv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twin-te/back/base"
	donationv1conv "github.com/twin-te/twin-te/back/handler/api/rpc/donation/v1/conv"
	donationv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/donation/v1"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/donation/v1/donationv1connect"
	donationmodule "github.com/twin-te/twin-te/back/module/donation"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var _ donationv1connect.DonationServiceHandler = (*impl)(nil)

type impl struct {
	uc donationmodule.UseCase
}

func (svc *impl) CreateOneTimeCheckoutSession(ctx context.Context, req *connect.Request[donationv1.CreateOneTimeCheckoutSessionRequest]) (res *connect.Response[donationv1.CreateOneTimeCheckoutSessionResponse], err error) {
	checkoutSessionId, err := svc.uc.CreateOneTimeCheckoutSession(ctx, int(req.Msg.Amount))
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.CreateOneTimeCheckoutSessionResponse{
		CheckoutSessionId: checkoutSessionId.String(),
	})

	return
}

func (svc *impl) CreateSubscriptionCheckoutSession(ctx context.Context, req *connect.Request[donationv1.CreateSubscriptionCheckoutSessionRequest]) (res *connect.Response[donationv1.CreateSubscriptionCheckoutSessionResponse], err error) {
	planID, err := idtype.ParseSubscriptionPlanID(req.Msg.PlanId)
	if err != nil {
		return
	}

	checkoutSessionId, err := svc.uc.CreateSubscriptionCheckoutSession(ctx, planID)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.CreateSubscriptionCheckoutSessionResponse{
		CheckoutSessionId: checkoutSessionId.String(),
	})

	return
}

func (svc *impl) GetPaymentUser(ctx context.Context, req *connect.Request[donationv1.GetPaymentUserRequest]) (res *connect.Response[donationv1.GetPaymentUserResponse], err error) {
	paymentUser, err := svc.uc.GetOrCreatePaymentUser(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.GetPaymentUserResponse{
		PaymentUser: donationv1conv.ToPBPaymentUser(paymentUser),
	})

	return
}

func (svc *impl) UpdatePaymentUser(ctx context.Context, req *connect.Request[donationv1.UpdatePaymentUserRequest]) (res *connect.Response[donationv1.UpdatePaymentUserResponse], err error) {
	in := donationmodule.UpdateOrCreatePaymentUserIn{}

	if req.Msg.DisplayName != nil {
		in.DisplayName, err = base.ToPtrWithErr(donationdomain.ParseDisplayName(*req.Msg.DisplayName))
		if err != nil {
			return
		}
	}

	if req.Msg.Link != nil {
		in.Link, err = base.ToPtrWithErr(donationdomain.ParseLink(*req.Msg.Link))
		if err != nil {
			return
		}
	}

	paymentUser, err := svc.uc.UpdateOrCreatePaymentUser(ctx, in)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.UpdatePaymentUserResponse{
		PaymentUser: donationv1conv.ToPBPaymentUser(paymentUser),
	})

	return
}

func (svc *impl) GetPaymentHistories(ctx context.Context, req *connect.Request[donationv1.GetPaymentHistoriesRequest]) (res *connect.Response[donationv1.GetPaymentHistoriesResponse], err error) {
	paymentHistories, err := svc.uc.GetPaymentHistories(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.GetPaymentHistoriesResponse{
		PaymentHistories: base.Map(paymentHistories, donationv1conv.ToPBPaymentHistory),
	})

	return
}

func (svc *impl) GetSubscriptionPlans(ctx context.Context, req *connect.Request[donationv1.GetSubscriptionPlansRequest]) (res *connect.Response[donationv1.GetSubscriptionPlansResponse], err error) {
	subscriptionPlans, err := svc.uc.GetSubscriptionPlans(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.GetSubscriptionPlansResponse{
		SubscriptionPlans: base.Map(subscriptionPlans, donationv1conv.ToPBSubscriptionPlan),
	})

	return
}

func (svc *impl) GetActiveSubscription(ctx context.Context, req *connect.Request[donationv1.GetActiveSubscriptionRequest]) (res *connect.Response[donationv1.GetActiveSubscriptionResponse], err error) {
	subscription, err := svc.uc.GetActiveSubscription(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.GetActiveSubscriptionResponse{
		Subscription: donationv1conv.ToPBSubscription(subscription),
	})

	return
}

func (svc *impl) Unsubscribe(ctx context.Context, req *connect.Request[donationv1.UnsubscribeRequest]) (res *connect.Response[donationv1.UnsubscribeResponse], err error) {
	id, err := idtype.ParseSubscriptionID(req.Msg.Id)
	if err != nil {
		return
	}

	err = svc.uc.Unsubscribe(ctx, id)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.UnsubscribeResponse{})

	return
}

func (svc *impl) GetTotalAmount(ctx context.Context, req *connect.Request[donationv1.GetTotalAmountRequest]) (res *connect.Response[donationv1.GetTotalAmountResponse], err error) {
	totalAmount, err := svc.uc.GetTotalAmount(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.GetTotalAmountResponse{
		TotalAmount: int32(totalAmount),
	})

	return
}

func (svc *impl) GetContributors(ctx context.Context, req *connect.Request[donationv1.GetContributorsRequest]) (res *connect.Response[donationv1.GetContributorsResponse], err error) {
	contributors, err := svc.uc.GetContributors(ctx)
	if err != nil {
		return
	}

	res = connect.NewResponse(&donationv1.GetContributorsResponse{
		Contributors: base.Map(contributors, donationv1conv.ToPBContributor),
	})

	return
}

func New(uc donationmodule.UseCase) *impl {
	return &impl{uc: uc}
}
