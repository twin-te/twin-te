package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/twin-te/twin-te/back/appctx"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/module/auth/accesscontroller"
	authrepository "github.com/twin-te/twin-te/back/module/auth/adapter/repository"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	timetablefactory "github.com/twin-te/twin-te/back/module/timetable/adapter/factory"
	timetableintegrator "github.com/twin-te/twin-te/back/module/timetable/adapter/integrator"
	timetablequery "github.com/twin-te/twin-te/back/module/timetable/adapter/query"
	timetablerepository "github.com/twin-te/twin-te/back/module/timetable/adapter/repository"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableusecase "github.com/twin-te/twin-te/back/module/timetable/usecase"
)

var (
	migrateYear            int
	migrateKdbJSONFilePath string
	migrateDryRun          bool
)

type kdbCourseCode struct {
	Code string `json:"code"`
}

var MigrateMissingCoursesCmd = &cobra.Command{
	Use:   "migrate-missing-courses",
	Short: "Migrate registered courses whose based courses are missing from KdB to manual registrations",
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

		year, err := shareddomain.ParseAcademicYear(migrateYear)
		if err != nil {
			log.Fatalln(err)
		}

		data, err := os.ReadFile(migrateKdbJSONFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		var kdbCourses []kdbCourseCode
		if err := json.Unmarshal(data, &kdbCourses); err != nil {
			log.Fatalln(err)
		}

		importedCodes, err := base.MapWithErr(kdbCourses, func(c kdbCourseCode) (timetabledomain.Code, error) {
			return timetabledomain.ParseCode(c.Code)
		})
		if err != nil {
			log.Fatalln(err)
		}

		ctx := appctx.SetActor(context.Background(), authdomain.NewUnknown(authdomain.PermissionExecuteBatchJob))

		missingCourses, err := timetableUseCase.ListMissingCourses(ctx, year, importedCodes)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%d missing courses found for year %d:\n", len(missingCourses), year.Int())
		for _, course := range missingCourses {
			fmt.Printf("  %s  %s\n", course.Code, course.Name)
		}

		if migrateDryRun {
			return
		}

		if err := timetableUseCase.MigrateMissingCourses(ctx, year, importedCodes); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Migration completed.")
	},
}

func init() {
	rootCmd.AddCommand(MigrateMissingCoursesCmd)

	MigrateMissingCoursesCmd.Flags().IntVar(&migrateYear, "year", 0, "academic year of courses to migrate")
	MigrateMissingCoursesCmd.Flags().StringVar(&migrateKdbJSONFilePath, "kdb-json-file-path", "", "kdb json file path containing imported courses")
	MigrateMissingCoursesCmd.Flags().BoolVar(&migrateDryRun, "dry-run", false, "list missing courses without making any changes")
}
