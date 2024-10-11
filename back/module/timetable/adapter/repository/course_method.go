package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateCourseMethods(db *gorm.DB, course *timetabledomain.Course) error {
	before := course.BeforeUpdated.MustGet()
	toCreate, toDelete := lo.Difference(course.Methods, before.Methods)

	if len(toCreate) != 0 {
		dbCourseMethods := base.MapWithArg(toCreate, course.ID, timetabledbmodel.ToDBCourseMethod)
		if err := db.Create(dbCourseMethods).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbCourseMethods := base.MapWithArg(toDelete, course.ID, timetabledbmodel.ToDBCourseMethod)
		return db.
			Where("course_id = ?", course.ID.String()).
			Where("method IN ?", base.Map(dbCourseMethods, func(dbCourseMethod timetabledbmodel.CourseMethod) any {
				return dbCourseMethod.Method
			})).
			Delete(&timetabledbmodel.CourseMethod{}).
			Error
	}

	return nil
}
