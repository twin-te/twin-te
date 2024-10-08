package timetablerepository

import (
	"context"
	"fmt"
	"slices"

	"github.com/samber/lo"
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

func (r *impl) FindRegisteredCourse(ctx context.Context, conds timetableport.FindRegisteredCourseConds, lock sharedport.Lock) (*timetabledomain.RegisteredCourse, error) {
	db := r.db.WithContext(ctx).Where("id = ?", conds.ID.String())

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	db = db.Clauses(clause.Locking{
		Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
		Table:    clause.Table{Name: clause.CurrentTable},
	})

	dbRegisteredCourse := new(timetabledbmodel.RegisteredCourse)

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Tags").Take(dbRegisteredCourse).Error
		return dbhelper.ConvertErrRecordNotFound(err)
	})
	if err != nil {
		return nil, err
	}

	return timetabledbmodel.FromDBRegisteredCourse(dbRegisteredCourse)
}

func (r *impl) ListRegisteredCourses(ctx context.Context, conds timetableport.ListRegisteredCoursesConds, lock sharedport.Lock) ([]*timetabledomain.RegisteredCourse, error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	if conds.Year != nil {
		db = db.Where("year = ?", conds.Year.Int())
	}

	if conds.CourseIDs != nil {
		db = db.Where("course_id IN ?", base.MapByString(*conds.CourseIDs))
	}

	db = db.Clauses(clause.Locking{
		Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
		Table:    clause.Table{Name: clause.CurrentTable},
	})

	var dbRegisteredCourses []*timetabledbmodel.RegisteredCourse

	err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Preload("Tags").Find(&dbRegisteredCourses).Error
	})
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
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(dbRegisteredCourses).Error
	}, nil)
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

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(columns) > 0 {
			if err := tx.Select(columns).Updates(dbRegisteredCourse).Error; err != nil {
				return err
			}
		}
		return r.updateRegisteredCourseTagIDs(tx, registeredCourse)
	}, nil)
}

func (r *impl) DeleteRegisteredCourses(ctx context.Context, conds timetableport.DeleteRegisteredCoursesConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if conds.ID != nil {
		db = db.Where("id = ?", conds.ID.String())
	}

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	return int(db.Delete(&timetabledbmodel.RegisteredCourse{}).RowsAffected), db.Error
}

func (r *impl) LoadCourseAssociationToRegisteredCourse(ctx context.Context, registeredCourses []*timetabledomain.RegisteredCourse, lock sharedport.Lock) error {
	courseIDToRegisteredCourse := make(map[idtype.CourseID]*timetabledomain.RegisteredCourse, len(registeredCourses))
	for _, registeredCourse := range registeredCourses {
		if registeredCourse.HasBasedCourse() && registeredCourse.CourseAssociation.IsAbsent() {
			courseIDToRegisteredCourse[registeredCourse.CourseID.MustGet()] = registeredCourse
		}
	}

	courses, err := r.ListCourses(ctx, timetableport.ListCoursesConds{
		IDs: lo.ToPtr(lo.Keys(courseIDToRegisteredCourse)),
	}, lock)
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
