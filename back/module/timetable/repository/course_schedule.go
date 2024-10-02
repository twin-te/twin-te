package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	"gorm.io/gorm"

	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func (r *impl) updateCourseSchedules(db *gorm.DB, course *timetabledomain.Course) error {
	toCreate, toDelete := lo.Difference(course.Schedules, course.EntityBeforeUpdated.Schedules)

	if len(toCreate) != 0 {
		dbCourseSchedules := base.MapWithArg(toCreate, course.ID, toDBCourseSchedule)

		if err := db.Create(dbCourseSchedules).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbCourseSchedules := base.MapWithArg(toDelete, course.ID, toDBCourseSchedule)

		return db.Where("course_id = ?", course.ID.String()).
			Where("(module,day,period,room) IN ?", base.Map(dbCourseSchedules, func(dbCourseSchedule model.CourseSchedule) []any {
				return []any{
					dbCourseSchedule.Module,
					dbCourseSchedule.Day,
					dbCourseSchedule.Period,
					dbCourseSchedule.Room,
				}
			})).
			Delete(&model.CourseSchedule{}).
			Error
	}

	return nil
}

func fromDBCourseSchedule(dbCourseSchedule model.CourseSchedule) (timetabledomain.Schedule, error) {
	return timetabledomain.ParseSchedule(
		dbCourseSchedule.Module,
		dbCourseSchedule.Day,
		int(dbCourseSchedule.Period),
		dbCourseSchedule.Room,
	)
}

func toDBCourseSchedule(schedule timetabledomain.Schedule, courseID idtype.CourseID) model.CourseSchedule {
	return model.CourseSchedule{
		CourseID: courseID.StringPtr(),
		Module:   schedule.Module.String(),
		Day:      schedule.Day.String(),
		Period:   int16(schedule.Period.Int()),
		Room:     schedule.Rooms,
	}
}
