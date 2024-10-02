package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	announcementrepository "github.com/twin-te/twin-te/back/module/announcement/repository"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	authrepository "github.com/twin-te/twin-te/back/module/auth/repository"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	donationrepository "github.com/twin-te/twin-te/back/module/donation/repository"
	schoolcalendarport "github.com/twin-te/twin-te/back/module/schoolcalendar/port"
	schoolcalendarrepository "github.com/twin-te/twin-te/back/module/schoolcalendar/repository"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	timetablerepository "github.com/twin-te/twin-te/back/module/timetable/repository"
)

// checkAggregatesCmd represents the check-aggregates command
var checkAggregatesCmd = &cobra.Command{
	Use:   "check-aggregates",
	Short: "Check if all aggregates are correctly reconstructed from repository",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		announcementRepository := announcementrepository.New(db)
		authRepository := authrepository.New(db)
		donationRepository := donationrepository.New(db)
		schoolcalendarRepository := schoolcalendarrepository.New()
		timetableRepository := timetablerepository.New(db)

		alreadyReads, err := announcementRepository.ListAlreadyReads(context.Background(), announcementport.ListAlreadyReadsConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d already reads are correctly reconstructed from repository", len(alreadyReads))

		announcements, err := announcementRepository.ListAnnouncements(context.Background(), announcementport.ListAnnouncementsConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d announcements are correctly reconstructed from repository", len(announcements))

		sessions, err := authRepository.ListSessions(context.Background(), authport.ListSessionsConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d sessions are correctly reconstructed from repository", len(sessions))

		users, err := authRepository.ListUsers(context.Background(), authport.ListUsersConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d users are correctly reconstructed from repository", len(users))

		paymentUsers, err := donationRepository.ListPaymentUsers(context.Background(), donationport.ListPaymentUsersConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d payment users are correctly reconstructed from repository", len(paymentUsers))

		events, err := schoolcalendarRepository.ListEvents(context.Background(), schoolcalendarport.ListEventsConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d events are correctly reconstructed from repository", len(events))

		moduleDetails, err := schoolcalendarRepository.ListModuleDetails(context.Background(), schoolcalendarport.ListModuleDetailsConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d module details are correctly reconstructed from repository", len(moduleDetails))

		courses, err := timetableRepository.ListCourses(context.Background(), timetableport.ListCoursesConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d courses are correctly reconstructed from repository", len(courses))

		registeredCourses, err := timetableRepository.ListRegisteredCourses(context.Background(), timetableport.ListRegisteredCoursesConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d registered courses are correctly reconstructed from repository", len(registeredCourses))

		tags, err := timetableRepository.ListTags(context.Background(), timetableport.ListTagsConds{}, sharedport.LockNone)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%d tags are correctly reconstructed from repository", len(tags))

		log.Println("all aggregates are correctly reconstructed from repository")
	},
}

func init() {
	rootCmd.AddCommand(checkAggregatesCmd)
}
