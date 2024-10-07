package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateCourseRecommendedGrades(db *gorm.DB, course *timetabledomain.Course) error {
	before := course.BeforeUpdated.MustGet()
	toCreate, toDelete := lo.Difference(course.RecommendedGrades, before.RecommendedGrades)

	if len(toCreate) != 0 {
		dbRecommendedGrades := base.MapWithArg(toCreate, course.ID, toDBRecommendedGrade)

		if err := db.Create(dbRecommendedGrades).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbRecommendedGrades := base.MapWithArg(toDelete, course.ID, toDBRecommendedGrade)

		return db.Where("course_id = ?", course.ID.String()).
			Where("grade IN ?", base.Map(dbRecommendedGrades, func(dbRecommendedGrade model.CourseRecommendedGrade) any {
				return dbRecommendedGrade.RecommendedGrade
			})).
			Delete(&model.CourseRecommendedGrade{}).
			Error
	}

	return nil
}

func fromDBRecommendedGrade(dbRecommendedGrade model.CourseRecommendedGrade) (timetabledomain.RecommendedGrade, error) {
	return timetabledomain.ParseRecommendedGrade(int(dbRecommendedGrade.RecommendedGrade))
}

func toDBRecommendedGrade(recommendedGrade timetabledomain.RecommendedGrade, courseID idtype.CourseID) model.CourseRecommendedGrade {
	return model.CourseRecommendedGrade{
		CourseID:         courseID.String(),
		RecommendedGrade: int16(recommendedGrade),
	}
}
