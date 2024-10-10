package timetablerepository

import (
	"context"
	"fmt"
	"slices"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) FindRegisteredCourse(ctx context.Context, filter timetableport.RegisteredCourseFilter, lock sharedport.Lock) (mo.Option[*timetabledomain.RegisteredCourse], error) {
	if !filter.IsUniqueFilter() {
		return mo.None[*timetabledomain.RegisteredCourse](), fmt.Errorf("%v is not unique", filter)
	}

	dbRegisteredCourse := new(timetabledbmodel.RegisteredCourse)

	err := r.transaction(ctx, func(tx *gorm.DB) error {
		tx = applyRegisteredCourseFilter(tx, filter)
		tx = dbhelper.ApplyLock(tx, lock)
		return tx.Preload("Tags").Take(dbRegisteredCourse).Error
	}, true)
	if err != nil {
		return dbhelper.ConvertErrRecordNotFound[*timetabledomain.RegisteredCourse](err)
	}

	return base.SomeWithErr(timetabledbmodel.FromDBRegisteredCourse(dbRegisteredCourse))
}

func (r *impl) ListRegisteredCourses(ctx context.Context, filter timetableport.RegisteredCourseFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*timetabledomain.RegisteredCourse, error) {
	var dbRegisteredCourses []*timetabledbmodel.RegisteredCourse

	err := r.transaction(ctx, func(tx *gorm.DB) error {
		tx = applyRegisteredCourseFilter(tx, filter)
		tx = dbhelper.ApplyLimitOffset(tx, limitOffset)
		tx = dbhelper.ApplyLock(tx, lock)
		return tx.
			Preload("Tags").
			Find(&dbRegisteredCourses).
			Error
	}, true)
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbRegisteredCourses, timetabledbmodel.FromDBRegisteredCourse)
}

func (r *impl) CreateRegisteredCourses(ctx context.Context, registeredCourses ...*timetabledomain.RegisteredCourse) error {
	dbRegisteredCourses, err := base.MapWithArgAndErr(registeredCourses, true, timetabledbmodel.ToDBRegisteredCourse)
	if err != nil {
		return err
	}
	dbRegisteredCourseTags := lo.Flatten(base.Map(dbRegisteredCourses, func(dbRegisteredCourse *timetabledbmodel.RegisteredCourse) []timetabledbmodel.RegisteredCourseTag {
		return dbRegisteredCourse.Tags
	}))
	return r.transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(dbRegisteredCourses).Error; err != nil {
			return err
		}
		if len(dbRegisteredCourseTags) > 0 {
			if err := tx.Create(dbRegisteredCourseTags).Error; err != nil {
				return err
			}
		}
		return nil
	}, false)
}

func (r *impl) UpdateRegisteredCourse(ctx context.Context, registeredCourse *timetabledomain.RegisteredCourse) error {
	before := registeredCourse.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if registeredCourse.UserID != before.UserID {
		columns = append(columns, "user_id")
	}

	if registeredCourse.Year != before.Year {
		columns = append(columns, "year")
	}

	if registeredCourse.CourseID != before.CourseID {
		columns = append(columns, "course_id")
	}

	if registeredCourse.Name != before.Name {
		columns = append(columns, "name")
	}

	if registeredCourse.Instructors != before.Instructors {
		columns = append(columns, "instructors")
	}

	if registeredCourse.Credit != before.Credit {
		columns = append(columns, "credit")
	}

	if !base.OptionEqualBy(registeredCourse.Methods, before.Methods, slices.Equal) {
		columns = append(columns, "methods")
	}

	if !base.OptionEqualBy(registeredCourse.Schedules, before.Schedules, slices.Equal) {
		columns = append(columns, "schedules")
	}

	if registeredCourse.Memo != before.Memo {
		columns = append(columns, "memo")
	}

	if registeredCourse.Attendance != before.Attendance {
		columns = append(columns, "attendance")
	}

	if registeredCourse.Absence != before.Absence {
		columns = append(columns, "absence")
	}

	if registeredCourse.Late != before.Late {
		columns = append(columns, "late")
	}

	dbRegisteredCourse, err := timetabledbmodel.ToDBRegisteredCourse(registeredCourse, false)
	if err != nil {
		return err
	}

	return r.transaction(ctx, func(tx *gorm.DB) error {
		if len(columns) > 0 {
			if err := tx.Select(columns).Updates(dbRegisteredCourse).Error; err != nil {
				return err
			}
		}
		return r.updateRegisteredCourseTagIDs(tx, registeredCourse)
	}, false)
}

func (r *impl) DeleteRegisteredCourses(ctx context.Context, filter timetableport.RegisteredCourseFilter) (rowsAffected int, err error) {
	err = r.transaction(ctx, func(tx *gorm.DB) error {
		tx = applyRegisteredCourseFilter(tx, filter)
		rowsAffected = int(tx.Delete(&timetabledbmodel.RegisteredCourse{}).RowsAffected)
		return tx.Error

	}, false)
	return
}

func (r *impl) LoadCourseAssociationToRegisteredCourse(ctx context.Context, registeredCourses []*timetabledomain.RegisteredCourse, lock sharedport.Lock) error {
	courseIDToRegisteredCourse := make(map[idtype.CourseID]*timetabledomain.RegisteredCourse, len(registeredCourses))
	for _, registeredCourse := range registeredCourses {
		if registeredCourse.HasBasedCourse() && registeredCourse.CourseAssociation.IsAbsent() {
			courseIDToRegisteredCourse[registeredCourse.CourseID.MustGet()] = registeredCourse
		}
	}

	courses, err := r.ListCourses(ctx, timetableport.CourseFilter{
		IDs: mo.Some(lo.Keys(courseIDToRegisteredCourse)),
	}, sharedport.LimitOffset{}, lock)
	if err != nil {
		return err
	}

	for _, course := range courses {
		courseIDToRegisteredCourse[course.ID].CourseAssociation.Set(course)
	}

	for courseID, registeredCourse := range courseIDToRegisteredCourse {
		if registeredCourse.CourseAssociation.IsAbsent() {
			return fmt.Errorf("can't load course (%s) to registered course (%s)", courseID, registeredCourse.ID)
		}
	}

	return nil
}

func applyRegisteredCourseFilter(db *gorm.DB, filter timetableport.RegisteredCourseFilter) *gorm.DB {
	if id, ok := filter.ID.Get(); ok {
		db = db.Where("id = ?", id.String())
	}

	if userID, ok := filter.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if year, ok := filter.Year.Get(); ok {
		db = db.Where("year = ?", year.Int())
	}

	if courseIDs, ok := filter.CourseIDs.Get(); ok {
		db = db.Where("course_id IN ?", base.MapByString(courseIDs))
	}

	return db
}
