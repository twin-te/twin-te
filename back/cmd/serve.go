package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/spf13/cobra"
	"github.com/twin-te/twin-te/back/appenv"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/handler"
	announcementfactory "github.com/twin-te/twin-te/back/module/announcement/adapter/factory"
	announcementrepository "github.com/twin-te/twin-te/back/module/announcement/adapter/repository"
	announcementdata "github.com/twin-te/twin-te/back/module/announcement/data"
	announcementusecase "github.com/twin-te/twin-te/back/module/announcement/usecase"
	"github.com/twin-te/twin-te/back/module/auth/accesscontroller"
	authfactory "github.com/twin-te/twin-te/back/module/auth/adapter/factory"
	authrepository "github.com/twin-te/twin-te/back/module/auth/adapter/repository"
	authusecase "github.com/twin-te/twin-te/back/module/auth/usecase"
	donationfactory "github.com/twin-te/twin-te/back/module/donation/adapter/factory"
	donationintegrator "github.com/twin-te/twin-te/back/module/donation/adapter/integrator"
	donationrepository "github.com/twin-te/twin-te/back/module/donation/adapter/repository"
	donationusecase "github.com/twin-te/twin-te/back/module/donation/usecase"
	schoolcalendarrepository "github.com/twin-te/twin-te/back/module/schoolcalendar/adapter/repository"
	schoolcalendardata "github.com/twin-te/twin-te/back/module/schoolcalendar/data"
	schoolcalendarusecase "github.com/twin-te/twin-te/back/module/schoolcalendar/usecase"
	timetablefactory "github.com/twin-te/twin-te/back/module/timetable/adapter/factory"
	timetableintegrator "github.com/twin-te/twin-te/back/module/timetable/adapter/integrator"
	timetablequery "github.com/twin-te/twin-te/back/module/timetable/adapter/query"
	timetablerepository "github.com/twin-te/twin-te/back/module/timetable/adapter/repository"
	timetableusecase "github.com/twin-te/twin-te/back/module/timetable/usecase"
	unifiedusecase "github.com/twin-te/twin-te/back/module/unified/usecase"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              appenv.SENTRY_DSN,
			SendDefaultPII:   true,
			EnableTracing:    true,
			TracesSampleRate: 0.3,
		})

		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}

		defer sentry.Flush(2 * time.Second)
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		nowFunc := func() time.Time { return time.Now().Truncate(time.Microsecond) }

		authFactory := authfactory.New(nowFunc)
		authRepository := authrepository.New(db)
		accessController := accesscontroller.New(authRepository)
		authUseCase := authusecase.New(accessController, authFactory, authRepository)

		announcementFactory := announcementfactory.New(nowFunc)
		announcementRepository := announcementrepository.New(db)
		announcementUsecase := announcementusecase.New(accessController, announcementFactory, announcementRepository)

		donationFactory := donationfactory.New()
		donationIntegrator := donationintegrator.New()
		donationRepository := donationrepository.New(db)
		donationUseCase := donationusecase.New(accessController, donationFactory, donationIntegrator, donationRepository)

		schoolcalendarRepository := schoolcalendarrepository.New()
		schoolcalendarUseCase := schoolcalendarusecase.New(accessController, schoolcalendarRepository)

		timetableFactory := timetablefactory.New(db)
		timetableIntegrator := timetableintegrator.New("")
		timetableQuery := timetablequery.New(db)
		timetableRepository := timetablerepository.New(db)
		timetableUseCase := timetableusecase.New(accessController, timetableFactory, timetableIntegrator, timetableQuery, timetableRepository)

		unifiedUseCase := unifiedusecase.New(accessController, schoolcalendarUseCase, timetableUseCase)

		announcements, err := announcementdata.LoadAnnouncements()
		if err != nil {
			log.Fatalln(err)
		}
		err = announcementRepository.CreateAnnouncements(context.Background(), announcements...)
		if err != nil {
			log.Fatalln(err)
		}

		events, err := schoolcalendardata.LoadEvents()
		if err != nil {
			log.Fatalln(err)
		}
		err = schoolcalendarRepository.CreateEvents(context.Background(), events...)
		if err != nil {
			log.Fatalln(err)
		}

		moduleDetails, err := schoolcalendardata.LoadModuleDetails()
		if err != nil {
			log.Fatalln(err)
		}
		err = schoolcalendarRepository.CreateModuleDetails(context.Background(), moduleDetails...)
		if err != nil {
			log.Fatalln(err)
		}

		h := handler.New(
			accessController,
			announcementUsecase,
			authUseCase,
			donationUseCase,
			schoolcalendarUseCase,
			timetableUseCase,
			unifiedUseCase,
		)

		mux := http.NewServeMux()
		mux.Handle("/", h)
		sentryhandler := sentryhttp.New(sentryhttp.Options{}).Handle(mux)

		log.Printf("listen and serve on %s\n", appenv.ADDR)

		if err := http.ListenAndServe(appenv.ADDR, sentryhandler); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
