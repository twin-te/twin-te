package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateCourseRecommendedGrades(db *gorm.DB, course *timetabledomain.Course) error {
	before := course.BeforeUpdated.MustGet()
	toCreate, toDelete := lo.Difference(course.RecommendedGrades, before.RecommendedGrades)

	if len(toCreate) != 0 {
		dbRecommendedGrades := base.MapWithArg(toCreate, course.ID, timetabledbmodel.ToDBRecommendedGrade)
		if err := db.Create(dbRecommendedGrades).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbRecommendedGrades := base.MapWithArg(toDelete, course.ID, timetabledbmodel.ToDBRecommendedGrade)
		return db.
			Where("course_id = ?", course.ID.String()).
			Where("grade IN ?", base.Map(dbRecommendedGrades, func(dbRecommendedGrade timetabledbmodel.CourseRecommendedGrade) any {
				return dbRecommendedGrade.RecommendedGrade
			})).
			Delete(&timetabledbmodel.CourseRecommendedGrade{}).
			Error
	}

	return nil
}
