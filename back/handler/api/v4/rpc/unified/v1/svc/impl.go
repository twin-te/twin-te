package unifiedv1svc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twin-te/back/base"
	schoolcalendarv1conv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/schoolcalendar/v1/conv"
	unifiedv1 "github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/unified/v1"
	"github.com/twin-te/twin-te/back/handler/api/v4/rpcgen/unified/v1/unifiedv1connect"
	sharedconv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/shared/conv"
	unifiedmodule "github.com/twin-te/twin-te/back/module/unified"
	timetablev1conv "github.com/twin-te/twin-te/back/handler/api/v4/rpc/timetable/v1/conv"
)


var _ unifiedv1connect.UnifiedServiceHandler = (*impl)(nil)

type impl struct {
	uc unifiedmodule.UseCase
}

func (svc *impl) GetByDate(ctx context.Context, req *connect.Request[unifiedv1.GetByDateRequest]) (res *connect.Response[unifiedv1.GetByDateResponse], err error) {
	date, err := sharedconv.FromPBRFC3339FullDate(req.Msg.Date)
	if err != nil {
		return
	}

	events, module, registeredCourses, err := svc.uc.GetByDate(ctx, date)
	if err != nil {
		return
	}
	
	pbEvents, err := base.MapWithErr(events, schoolcalendarv1conv.ToPBEvent)
	if err != nil {
		return
	}

	pbModule, err := schoolcalendarv1conv.ToPBModule(module)
	if err != nil {
		return
	}

	pbRegisteredCourses, err := base.MapWithErr(registeredCourses, timetablev1conv.ToPBRegisteredCourse)
	if err != nil {
		return
	}
	
	res = connect.NewResponse(&unifiedv1.GetByDateResponse{
		Events: pbEvents,
		Module: pbModule,
		RegisteredCourses: pbRegisteredCourses,
	})

	return
}

func New(uc unifiedmodule.UseCase) *impl {
	return &impl{uc: uc}
}
