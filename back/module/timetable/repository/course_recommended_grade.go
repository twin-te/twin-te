package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	"github.com/twin-te/twinte-back/db/gen/model"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateCourseRecommendedGrades(db *gorm.DB, course *timetabledomain.Course) error {
	toCreate, toDelete := lo.Difference(course.RecommendedGrades, course.EntityBeforeUpdated.RecommendedGrades)

	if len(toCreate) != 0 {
		dbRecommendedGrades := base.MapWithArg(toCreate, course.ID, toDBRecommendedGrade)

		if err := db.Create(dbRecommendedGrades).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbRecommendedGrades := base.MapWithArg(toCreate, course.ID, toDBRecommendedGrade)

		return db.Where("course_id = ?", course.ID.String()).
			Where("grade IN ?", base.Map(dbRecommendedGrades, func(dbRecommendedGrade model.CourseRecommendedGrade) any {
				return dbRecommendedGrade.Grade
			})).
			Delete(&model.CourseRecommendedGrade{}).
			Error
	}

	return nil
}

func fromDBRecommendedGrade(dbRecommendedGrade model.CourseRecommendedGrade) (timetabledomain.RecommendedGrade, error) {
	return timetabledomain.ParseRecommendedGrade(int(dbRecommendedGrade.Grade))
}

func toDBRecommendedGrade(recommendedGrade timetabledomain.RecommendedGrade, courseID idtype.CourseID) model.CourseRecommendedGrade {
	return model.CourseRecommendedGrade{
		CourseID: courseID.StringPtr(),
		Grade:    int16(recommendedGrade),
	}
}
