package schoolcalendarv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twinte-back/base"
	schoolcalendarv1conv "github.com/twin-te/twinte-back/handler/api/rpc/schoolcalendar/v1/conv"
	sharedconv "github.com/twin-te/twinte-back/handler/api/rpc/shared/conv"
	schoolcalendarv1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	"github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1/schoolcalendarv1connect"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
)

var _ schoolcalendarv1connect.SchoolCalendarServiceHandler = (*impl)(nil)

type impl struct {
	uc schoolcalendarmodule.UseCase
}

func (svc *impl) GetEventsByDate(ctx context.Context, req *connect.Request[schoolcalendarv1.GetEventsByDateRequest]) (res *connect.Response[schoolcalendarv1.GetEventsByDateResponse], err error) {
	date, err := sharedconv.FromPBRFC3339FullDate(req.Msg.Date)
	if err != nil {
		return
	}

	events, err := svc.uc.GetEventsByDate(ctx, date)
	if err != nil {
		return
	}

	pbEvents, err := base.MapWithErr(events, schoolcalendarv1conv.ToPBEvent)
	if err != nil {
		return
	}

	res = connect.NewResponse(&schoolcalendarv1.GetEventsByDateResponse{
		Events: pbEvents,
	})

	return
}

func (svc *impl) GetModuleByDate(ctx context.Context, req *connect.Request[schoolcalendarv1.GetModuleByDateRequest]) (res *connect.Response[schoolcalendarv1.GetModuleByDateResponse], err error) {
	date, err := sharedconv.FromPBRFC3339FullDate(req.Msg.Date)
	if err != nil {
		return
	}

	module, err := svc.uc.GetModuleByDate(ctx, date)
	if err != nil {
		return
	}

	pbModule, err := schoolcalendarv1conv.ToPBModule(module)
	if err != nil {
		return
	}

	res = connect.NewResponse(&schoolcalendarv1.GetModuleByDateResponse{
		Module: pbModule,
	})

	return
}

func New(uc schoolcalendarmodule.UseCase) *impl {
	return &impl{uc: uc}
}
