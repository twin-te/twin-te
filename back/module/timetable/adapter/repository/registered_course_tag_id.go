package timetablerepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	"gorm.io/gorm"
)

func (r *impl) updateRegisteredCourseTagIDs(db *gorm.DB, registeredCourse *timetabledomain.RegisteredCourse) error {
	before := registeredCourse.BeforeUpdated.MustGet()
	toCreate, toDelete := lo.Difference(registeredCourse.TagIDs, before.TagIDs)

	if len(toCreate) != 0 {
		dbTags := base.MapWithArg(toCreate, registeredCourse.ID, timetabledbmodel.ToDBRegisteredCourseTag)
		if err := db.Create(dbTags).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		return db.
			Where("registered_course_id = ?", registeredCourse.ID.String()).
			Where("tag_id IN ?", base.MapByString(toDelete)).
			Delete(&timetabledbmodel.RegisteredCourseTag{}).
			Error
	}

	return nil
}
