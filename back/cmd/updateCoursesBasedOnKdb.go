package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/twin-te/twin-te/back/appctx"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/module/auth/accesscontroller"
	authrepository "github.com/twin-te/twin-te/back/module/auth/adapter/repository"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetablefactory "github.com/twin-te/twin-te/back/module/timetable/adapter/factory"
	timetableintegrator "github.com/twin-te/twin-te/back/module/timetable/adapter/integrator"
	timetablequery "github.com/twin-te/twin-te/back/module/timetable/adapter/query"
	timetablerepository "github.com/twin-te/twin-te/back/module/timetable/adapter/repository"
	timetableusecase "github.com/twin-te/twin-te/back/module/timetable/usecase"
)

var (
	year            int
	kdbJSONFilePath string
)

// UpdateCoursesBasedOnKdBCmd represents the update-courses-based-on-kdb command
var UpdateCoursesBasedOnKdBCmd = &cobra.Command{
	Use:   "update-courses-based-on-kdb",
	Short: "Update courses based on KdB",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		authRepository := authrepository.New(db)
		accessController := accesscontroller.New(authRepository)

		timetableFactory := timetablefactory.New(db)
		timetableIntegrator := timetableintegrator.New(kdbJSONFilePath)
		timetableQuery := timetablequery.New(db)
		timetableRepository := timetablerepository.New(db)
		timetableUseCase := timetableusecase.New(accessController, timetableFactory, timetableIntegrator, timetableQuery, timetableRepository)

		year, err := shareddomain.ParseAcademicYear(year)
		if err != nil {
			log.Fatalln(err)
		}

		ctx := appctx.SetActor(context.Background(), authdomain.NewUnknown(authdomain.PermissionExecuteBatchJob))

		if err := timetableUseCase.UpdateCoursesBasedOnKdB(ctx, year); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(UpdateCoursesBasedOnKdBCmd)

	UpdateCoursesBasedOnKdBCmd.Flags().IntVar(&year, "year", 0, "academic year of courses you want to update")
	UpdateCoursesBasedOnKdBCmd.Flags().StringVar(&kdbJSONFilePath, "kdb-json-file-path", "", "kdb json file path that is used in timetable integrator")
}
