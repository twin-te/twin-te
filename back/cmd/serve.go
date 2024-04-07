package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/twin-te/twin-te/back/appenv"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/handler"
	announcementdata "github.com/twin-te/twin-te/back/module/announcement/data"
	announcementfactory "github.com/twin-te/twin-te/back/module/announcement/factory"
	announcementrepository "github.com/twin-te/twin-te/back/module/announcement/repository"
	announcementusecase "github.com/twin-te/twin-te/back/module/announcement/usecase"
	"github.com/twin-te/twin-te/back/module/auth/accesscontroller"
	authfactory "github.com/twin-te/twin-te/back/module/auth/factory"
	authrepository "github.com/twin-te/twin-te/back/module/auth/repository"
	authusecase "github.com/twin-te/twin-te/back/module/auth/usecase"
	donationfactory "github.com/twin-te/twin-te/back/module/donation/factory"
	donationgateway "github.com/twin-te/twin-te/back/module/donation/gateway"
	donationrepository "github.com/twin-te/twin-te/back/module/donation/repository"
	donationusecase "github.com/twin-te/twin-te/back/module/donation/usecase"
	schoolcalendardata "github.com/twin-te/twin-te/back/module/schoolcalendar/data"
	schoolcalendarrepository "github.com/twin-te/twin-te/back/module/schoolcalendar/repository"
	schoolcalendarusecase "github.com/twin-te/twin-te/back/module/schoolcalendar/usecase"
	timetablefactory "github.com/twin-te/twin-te/back/module/timetable/factory"
	timetablegateway "github.com/twin-te/twin-te/back/module/timetable/gateway"
	timetablerepository "github.com/twin-te/twin-te/back/module/timetable/repository"
	timetableusecase "github.com/twin-te/twin-te/back/module/timetable/usecase"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
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
		donationGateway := donationgateway.New()
		donationRepository := donationrepository.New(db)
		donationUseCase := donationusecase.New(accessController, donationFactory, donationGateway, donationRepository)

		schoolcalendarRepository := schoolcalendarrepository.New()
		schoolcalendarUseCase := schoolcalendarusecase.New(accessController, schoolcalendarRepository)

		timetableFactory := timetablefactory.New(db)
		timetableGateWay := timetablegateway.New("")
		timetableRepository := timetablerepository.New(db)
		timetableUseCase := timetableusecase.New(accessController, timetableFactory, timetableGateWay, timetableRepository)

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
		)

		mux := http.NewServeMux()
		mux.Handle("/", h)

		log.Printf("listen and serve on %s\n", appenv.ADDR)

		if err := http.ListenAndServe(appenv.ADDR, mux); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
