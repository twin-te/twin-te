package timetablerepository

import (
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/twin-te/twin-te/back/base"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func (r *impl) updateCourseSchedules(db *gorm.DB, course *timetabledomain.Course) error {
	before := course.BeforeUpdated.MustGet()
	toCreate, toDelete := lo.Difference(course.Schedules, before.Schedules)

	if len(toCreate) != 0 {
		dbCourseSchedules := base.MapWithArg(toCreate, course.ID, timetabledbmodel.ToDBCourseSchedule)
		if err := db.Create(dbCourseSchedules).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbCourseSchedules := base.MapWithArg(toDelete, course.ID, timetabledbmodel.ToDBCourseSchedule)
		return db.
			Where("course_id = ?", course.ID.String()).
			Where("(module, day, period, locations) IN ?", base.Map(dbCourseSchedules, func(dbCourseSchedule timetabledbmodel.CourseSchedule) []any {
				return []any{
					dbCourseSchedule.Module,
					dbCourseSchedule.Day,
					dbCourseSchedule.Period,
					dbCourseSchedule.Locations,
				}
			})).
			Delete(&timetabledbmodel.CourseSchedule{}).
			Error
	}

	return nil
}
