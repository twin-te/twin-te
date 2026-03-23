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
	copySourceYear     int
	copyMaxFutureYears int
)

var CopyCoursesToFutureYearsCmd = &cobra.Command{
	Use:   "copy-courses-to-future-years",
	Short: "Copy courses from source year to future years",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbhelper.NewDB()
		if err != nil {
			log.Fatalln(err)
		}

		authRepository := authrepository.New(db)
		accessController := accesscontroller.New(authRepository)

		timetableFactory := timetablefactory.New(db)
		timetableIntegrator := timetableintegrator.New("")
		timetableQuery := timetablequery.New(db)
		timetableRepository := timetablerepository.New(db)
		timetableUseCase := timetableusecase.New(accessController, timetableFactory, timetableIntegrator, timetableQuery, timetableRepository)

		sourceYear, err := shareddomain.ParseAcademicYear(copySourceYear)
		if err != nil {
			log.Fatalln(err)
		}

		ctx := appctx.SetActor(context.Background(), authdomain.NewUnknown(authdomain.PermissionExecuteBatchJob))

		if err := timetableUseCase.CopyCoursesToFutureYears(ctx, sourceYear, copyMaxFutureYears); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(CopyCoursesToFutureYearsCmd)

	CopyCoursesToFutureYearsCmd.Flags().IntVar(&copySourceYear, "source-year", 0, "academic year of courses to copy from")
	CopyCoursesToFutureYearsCmd.Flags().IntVar(&copyMaxFutureYears, "max-future-years", 1, "number of future years to copy to")
}
