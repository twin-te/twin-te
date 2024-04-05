package apirpc

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	announcementv1svc "github.com/twin-te/twin-te/back/handler/api/rpc/announcement/v1/svc"
	authv1svc "github.com/twin-te/twin-te/back/handler/api/rpc/auth/v1/svc"
	donationv1svc "github.com/twin-te/twin-te/back/handler/api/rpc/donation/v1/svc"
	schoolcalendarv1svc "github.com/twin-te/twin-te/back/handler/api/rpc/schoolcalendar/v1/svc"
	timetablev1svc "github.com/twin-te/twin-te/back/handler/api/rpc/timetable/v1/srv"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/announcement/v1/announcementv1connect"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/auth/v1/authv1connect"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/donation/v1/donationv1connect"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/schoolcalendar/v1/schoolcalendarv1connect"
	"github.com/twin-te/twin-te/back/handler/api/rpcgen/timetable/v1/timetablev1connect"
	"github.com/twin-te/twin-te/back/handler/common/interceptor"
	announcementmodule "github.com/twin-te/twin-te/back/module/announcement"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	donationmodule "github.com/twin-te/twin-te/back/module/donation"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
)

var _ http.Handler = (*impl)(nil)

// impl handles requests with paths beginning with the following prefixes
//   - "/announcement.v1"
//   - "/auth.v1"
//   - "/donation.v1"
//   - "/schoolcalendar.v1"
//   - "/timetable.v1"
type impl struct {
	pattenToHandler map[string]http.Handler
}

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for patten, handler := range h.pattenToHandler {
		if strings.HasPrefix(r.URL.Path, patten) {
			handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func (h *impl) register(patten string, handler http.Handler) {
	h.pattenToHandler[patten] = handler
}

func New(
	accessController authmodule.AccessController,
	announcementUsecase announcementmodule.UseCase,
	authUseCase authmodule.UseCase,
	donationUseCase donationmodule.UseCase,
	schoolcalendarUseCase schoolcalendarmodule.UseCase,
	timetableUseCase timetablemodule.UseCase,
) *impl {
	h := &impl{
		pattenToHandler: map[string]http.Handler{},
	}

	handlerOptions := []connect.HandlerOption{
		connect.WithRecover(func(ctx context.Context, s connect.Spec, h http.Header, a any) error {
			return connect.NewError(connect.CodeInternal, fmt.Errorf("panicked during %s: %+v", s.Procedure, a))
		}),
		connect.WithInterceptors(interceptor.NewErrorInterceptor(), interceptor.NewAuthInterceptor(accessController)),
	}

	// "/announcement.v1"
	announcementv1Svc := announcementv1svc.New(announcementUsecase)
	h.register(announcementv1connect.NewAnnouncementServiceHandler(announcementv1Svc, handlerOptions...))

	// "/auth.v1"
	authv1Svc := authv1svc.New(authUseCase)
	h.register(authv1connect.NewAuthServiceHandler(authv1Svc, handlerOptions...))

	// "/donation.v1"
	donationv1Svc := donationv1svc.New(donationUseCase)
	h.register(donationv1connect.NewDonationServiceHandler(donationv1Svc, handlerOptions...))

	// "/schoolcalendar.v1"
	schoolcalendarv1Svc := schoolcalendarv1svc.New(schoolcalendarUseCase)
	h.register(schoolcalendarv1connect.NewSchoolCalendarServiceHandler(schoolcalendarv1Svc, handlerOptions...))

	// "/timetable.v1"
	timetablev1Svc := timetablev1svc.New(timetableUseCase)
	h.register(timetablev1connect.NewTimetableServiceHandler(timetablev1Svc, handlerOptions...))

	return h
}
