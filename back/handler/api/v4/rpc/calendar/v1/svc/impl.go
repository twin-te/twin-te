package calendarv1svc

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twin-te/back/appenv"
	calendarv1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/calendar/v1"
	"github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/calendar/v1/calendarv1connect"
	calendarv1handler "github.com/twin-te/twin-te/back/handler/calendar/v1"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var _ calendarv1connect.CalendarServiceHandler = (*impl)(nil)

type impl struct {
	uc calendarmodule.UseCase
}

func (svc *impl) GetIcalSubscriptionUrl(ctx context.Context, req *connect.Request[calendarv1.GetIcalSubscriptionUrlRequest]) (res *connect.Response[calendarv1.GetIcalSubscriptionUrlResponse], err error) {
	optID, err := svc.uc.GetIcalSubscriptionID(ctx)
	if err != nil {
		return nil, err
	}

	var url *string
	if id, ok := optID.Get(); ok {
		urlVal := buildIcalSubscriptionUrl(id)
		url = &urlVal
	}

	res = connect.NewResponse(&calendarv1.GetIcalSubscriptionUrlResponse{
		Url: url,
	})

	return
}

func (svc *impl) EnableIcalSubscription(ctx context.Context, req *connect.Request[calendarv1.EnableIcalSubscriptionRequest]) (*connect.Response[calendarv1.EnableIcalSubscriptionResponse], error) {
	id, err := svc.uc.EnableIcalSubscription(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&calendarv1.EnableIcalSubscriptionResponse{Url: buildIcalSubscriptionUrl(id)}), nil
}

func (svc *impl) DisableIcalSubscription(ctx context.Context, req *connect.Request[calendarv1.DisableIcalSubscriptionRequest]) (*connect.Response[calendarv1.DisableIcalSubscriptionResponse], error) {
	if err := svc.uc.DisableIcalSubscription(ctx); err != nil {
		return nil, err
	}
	return connect.NewResponse(&calendarv1.DisableIcalSubscriptionResponse{}), nil
}

func buildIcalSubscriptionUrl(id idtype.IcalSubscriptionID) string {
	return fmt.Sprintf("%s%s/timetable.ics?token=%s", appenv.APP_URL, calendarv1handler.PathPrefix, id.String())
}

func New(uc calendarmodule.UseCase) *impl {
	return &impl{uc: uc}
}
