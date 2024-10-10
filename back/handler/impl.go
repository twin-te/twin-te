package handler

import (
	"net/http"
	"strings"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"github.com/twin-te/twin-te/back/appenv"
	apiv4rpc "github.com/twin-te/twin-te/back/handler/api/v4/rpc"
	authv4 "github.com/twin-te/twin-te/back/handler/auth/v4"
	calendarv1beta "github.com/twin-te/twin-te/back/handler/calendar/v1beta"
	announcementmodule "github.com/twin-te/twin-te/back/module/announcement"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	donationmodule "github.com/twin-te/twin-te/back/module/donation"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
)

var _ http.Handler = (*impl)(nil)

type impl struct {
	authv4Handler         http.Handler
	calendarv1betaHandler http.Handler
	apiv4RPCHandler       http.Handler
}

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api/v4"):
		http.StripPrefix("/api/v4", h.apiv4RPCHandler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/auth/v4"):
		http.StripPrefix("/auth/v4", h.authv4Handler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/calendar/v1beta"):
		http.StripPrefix("/calendar/v1beta", h.calendarv1betaHandler).ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

func New(
	accessController authmodule.AccessController,
	announcementUsecase announcementmodule.UseCase,
	authUseCase authmodule.UseCase,
	donationUseCase donationmodule.UseCase,
	schoolcalendarUseCase schoolcalendarmodule.UseCase,
	timetableUseCase timetablemodule.UseCase,
) *impl {
	authv4Handler := authv4.New(
		accessController,
		authUseCase,
	)

	calendarv1betaHandler := calendarv1beta.New()

	var apiv4RPCHandler http.Handler = apiv4rpc.New(
		accessController,
		announcementUsecase,
		authUseCase,
		donationUseCase,
		schoolcalendarUseCase,
		timetableUseCase,
	)

	apiv4RPCHandler = cors.New(cors.Options{
		AllowedOrigins:   appenv.CORS_ALLOWED_ORIGINS,
		AllowedMethods:   connectcors.AllowedMethods(),
		AllowCredentials: true,
		AllowedHeaders:   connectcors.AllowedHeaders(),
		ExposedHeaders:   connectcors.ExposedHeaders(),
	}).Handler(apiv4RPCHandler)

	h := &impl{
		authv4Handler:         authv4Handler,
		calendarv1betaHandler: calendarv1betaHandler,
		apiv4RPCHandler:       apiv4RPCHandler,
	}

	return h
}
