package calendarv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	calendarv1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/calendar/v1"
	"github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/calendar/v1/calendarv1connect"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
)

var _ calendarv1connect.CalendarServiceHandler = (*impl)(nil)

type impl struct {
	uc calendarmodule.UseCase
}

func (svc *impl) GetIcalSubscriptionUrl(ctx context.Context, req *connect.Request[calendarv1.GetIcalSubscriptionUrlRequest]) (res *connect.Response[calendarv1.GetIcalSubscriptionUrlResponse], err error) {
	url, ok, err := svc.uc.GetIcalSubscriptionUrl(ctx)
	if !ok || err != nil {
		return nil, err
	}

	res = connect.NewResponse(&calendarv1.GetIcalSubscriptionUrlResponse{
		Url: &url,
	})

	return
}

func (svc *impl) EnableIcalSubscription(ctx context.Context, req *connect.Request[calendarv1.EnableIcalSubscriptionRequest]) (*connect.Response[calendarv1.EnableIcalSubscriptionResponse], error) {
	url, err := svc.uc.EnableIcalSubscription(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&calendarv1.EnableIcalSubscriptionResponse{Url: url}), nil
}

func (svc *impl) DisableIcalSubscription(ctx context.Context, req *connect.Request[calendarv1.DisableIcalSubscriptionRequest]) (*connect.Response[calendarv1.DisableIcalSubscriptionResponse], error) {
	if err := svc.uc.DisableIcalSubscription(ctx); err != nil {
		return nil, err
	}
	return connect.NewResponse(&calendarv1.DisableIcalSubscriptionResponse{}), nil
}

func New(uc calendarmodule.UseCase) *impl {
	return &impl{uc: uc}
}
