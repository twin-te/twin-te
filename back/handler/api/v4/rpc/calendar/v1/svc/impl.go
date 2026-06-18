package calendarv1svc

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twin-te/back/appenv"
	"github.com/twin-te/twin-te/back/base"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/shared/conv"
	calendarv1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/calendar/v1"
	"github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/calendar/v1/calendarv1connect"
	calendarv1handler "github.com/twin-te/twin-te/back/handler/calendar/v1"
	calendarmodule "github.com/twin-te/twin-te/back/module/calendar"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
)

var _ calendarv1connect.CalendarServiceHandler = (*impl)(nil)

type impl struct {
	uc calendarmodule.UseCase
}

func (svc *impl) GetIcalSubscriptionUrl(ctx context.Context, req *connect.Request[calendarv1.GetIcalSubscriptionUrlRequest]) (res *connect.Response[calendarv1.GetIcalSubscriptionUrlResponse], err error) {
	optSub, err := svc.uc.FindIcalSubscription(ctx)
	if err != nil {
		return nil, err
	}

	msg := &calendarv1.GetIcalSubscriptionUrlResponse{}
	if sub, ok := optSub.Get(); ok {
		urlVal := buildIcalSubscriptionUrl(sub.ID)
		msg.Url = &urlVal
		msg.Mode = toPBIcalSubscriptionMode(sub.Mode)
		msg.TargetTagIds = base.Map(sub.TargetTagIDs, sharedconv.ToPBUUID[idtype.TagID])
	}

	res = connect.NewResponse(msg)

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

func (svc *impl) UpdateIcalSubscription(ctx context.Context, req *connect.Request[calendarv1.UpdateIcalSubscriptionRequest]) (*connect.Response[calendarv1.UpdateIcalSubscriptionResponse], error) {
	mode, err := fromPBIcalSubscriptionMode(req.Msg.Mode)
	if err != nil {
		return nil, err
	}

	targetTagIDs, err := base.MapWithArgAndErr(req.Msg.TargetTagIds, idtype.ParseTagID, sharedconv.FromPBUUID[idtype.TagID])
	if err != nil {
		return nil, err
	}

	if err := svc.uc.UpdateIcalSubscription(ctx, mode, targetTagIDs); err != nil {
		return nil, err
	}

	return connect.NewResponse(&calendarv1.UpdateIcalSubscriptionResponse{}), nil
}

func fromPBIcalSubscriptionMode(mode calendarv1.IcalSubscriptionMode) (calendardomain.IcalSubscriptionMode, error) {
	switch mode {
	case calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_SYNC:
		return calendardomain.IcalSubscriptionModeSync, nil
	case calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_EXCLUDE:
		return calendardomain.IcalSubscriptionModeExclude, nil
	case calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_TRANSPARENT:
		return calendardomain.IcalSubscriptionModeTransparent, nil
	default:
		return "", sharederr.NewInvalidArgument("invalid ical subscription mode")
	}
}

func toPBIcalSubscriptionMode(mode calendardomain.IcalSubscriptionMode) calendarv1.IcalSubscriptionMode {
	switch mode {
	case calendardomain.IcalSubscriptionModeSync:
		return calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_SYNC
	case calendardomain.IcalSubscriptionModeExclude:
		return calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_EXCLUDE
	case calendardomain.IcalSubscriptionModeTransparent:
		return calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_TRANSPARENT
	default:
		return calendarv1.IcalSubscriptionMode_ICAL_SUBSCRIPTION_MODE_UNSPECIFIED
	}
}

func buildIcalSubscriptionUrl(id idtype.IcalSubscriptionID) string {
	return fmt.Sprintf("%s%s/timetable.ics?token=%s", appenv.APP_URL, calendarv1handler.PathPrefix, id.String())
}

func New(uc calendarmodule.UseCase) *impl {
	return &impl{uc: uc}
}
