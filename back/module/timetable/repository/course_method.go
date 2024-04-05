package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateCourseMethods(db *gorm.DB, course *timetabledomain.Course) error {
	toCreate, toDelete := lo.Difference(course.Methods, course.EntityBeforeUpdated.Methods)

	if len(toCreate) != 0 {
		dbCourseMethods := base.MapWithArg(toDelete, course.ID, toDBCourseMethod)

		if err := db.Create(dbCourseMethods).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbCourseMethods := base.MapWithArg(toDelete, course.ID, toDBCourseMethod)

		return db.Where("course_id = ?", course.ID.String()).
			Where("method IN ?", base.Map(dbCourseMethods, func(dbCourseMethod model.CourseMethod) any {
				return dbCourseMethod.Method
			})).
			Delete(&model.CourseMethod{}).
			Error
	}

	return nil
}

func fromDBCourseMethod(method model.CourseMethod) (timetabledomain.CourseMethod, error) {
	return timetabledomain.ParseCourseMethod(method.Method)
}

func toDBCourseMethod(method timetabledomain.CourseMethod, courseID idtype.CourseID) model.CourseMethod {
	return model.CourseMethod{
		CourseID: courseID.StringPtr(),
		Method:   method.String(),
	}
}
