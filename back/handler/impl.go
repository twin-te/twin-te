package handler

import (
	"net/http"
	"strings"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"github.com/twin-te/twin-te/back/appenv"
	apirpc "github.com/twin-te/twin-te/back/handler/api/rpc"
	authv3 "github.com/twin-te/twin-te/back/handler/auth/v3"
	calendarv1beta "github.com/twin-te/twin-te/back/handler/calendar/v1beta"
	announcementmodule "github.com/twin-te/twin-te/back/module/announcement"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	donationmodule "github.com/twin-te/twin-te/back/module/donation"
	schoolcalendarmodule "github.com/twin-te/twin-te/back/module/schoolcalendar"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
)

var _ http.Handler = (*impl)(nil)

type impl struct {
	authv3Handler         http.Handler
	calendarv1betaHandler http.Handler
	apiRPCHandler         http.Handler
}

func (h *impl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api"):
		http.StripPrefix("/api", h.apiRPCHandler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/auth/v3"):
		http.StripPrefix("/auth/v3", h.authv3Handler).ServeHTTP(w, r)
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
	authv3Handler := authv3.New(
		accessController,
		authUseCase,
	)

	calendarv1betaHandler := calendarv1beta.New()

	var apiRPCHandler http.Handler = apirpc.New(
		accessController,
		announcementUsecase,
		authUseCase,
		donationUseCase,
		schoolcalendarUseCase,
		timetableUseCase,
	)

	apiRPCHandler = cors.New(cors.Options{
		AllowedOrigins:   appenv.CORS_ALLOWED_ORIGINS,
		AllowedMethods:   connectcors.AllowedMethods(),
		AllowCredentials: true,
		AllowedHeaders:   connectcors.AllowedHeaders(),
		ExposedHeaders:   connectcors.ExposedHeaders(),
	}).Handler(apiRPCHandler)

	h := &impl{
		authv3Handler:         authv3Handler,
		calendarv1betaHandler: calendarv1betaHandler,
		apiRPCHandler:         apiRPCHandler,
	}

	return h
}
