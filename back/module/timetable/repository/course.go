package timetablerepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) FindCourse(ctx context.Context, conds timetableport.FindCourseConds, lock sharedport.Lock) (*timetabledomain.Course, error) {
	dbCourse := new(timetabledbmodel.Course)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			WithContext(ctx).
			Where("year = ?", conds.Year.Int()).
			Where("code = ?", conds.Code.String()).
			Clauses(clause.Locking{
				Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
				Table:    clause.Table{Name: clause.CurrentTable},
			}).
			Preload("RecommendedGrades").
			Preload("Methods").
			Preload("Schedules").
			Take(dbCourse).
			Error
		return dbhelper.ConvertErrRecordNotFound(err)
	}, nil)
	if err != nil {
		return nil, err
	}

	return timetabledbmodel.FromDBCourse(dbCourse)
}

func (r *impl) ListCourses(ctx context.Context, conds timetableport.ListCoursesConds, lock sharedport.Lock) ([]*timetabledomain.Course, error) {
	db := r.db.WithContext(ctx)

	if conds.IDs != nil {
		db = db.Where("id IN ?", base.MapByString(*conds.IDs))
	}

	if conds.Year != nil {
		db = db.Where("year = ?", conds.Year.Int())
	}

	if conds.Codes != nil {
		db = db.Where("code IN ?", base.MapByString(*conds.Codes))
	}

	var dbCourses []*timetabledbmodel.Course

	err := db.Transaction(func(tx *gorm.DB) error {
		return tx.
			Clauses(clause.Locking{
				Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
				Table:    clause.Table{Name: clause.CurrentTable},
			}).
			Preload("RecommendedGrades").
			Preload("Methods").
			Preload("Schedules").
			Find(&dbCourses).
			Error
	}, nil)
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbCourses, timetabledbmodel.FromDBCourse)
}

func (r *impl) CreateCourses(ctx context.Context, courses ...*timetabledomain.Course) error {
	dbCourses := base.MapWithArg(courses, true, timetabledbmodel.ToDBCourse)
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(dbCourses).Error
	}, nil)
}

func (r *impl) UpdateCourse(ctx context.Context, course *timetabledomain.Course) error {
	before := course.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if course.Year != before.Year {
		columns = append(columns, "year")
	}

	if course.Code != before.Code {
		columns = append(columns, "code")
	}

	if course.Name != before.Name {
		columns = append(columns, "name")
	}

	if course.Instructors != before.Instructors {
		columns = append(columns, "instructors")
	}

	if course.Credit != before.Credit {
		columns = append(columns, "credit")
	}

	if course.Overview != before.Overview {
		columns = append(columns, "overview")
	}

	if course.Remarks != before.Remarks {
		columns = append(columns, "remarks")
	}

	if !course.LastUpdatedAt.Equal(before.LastUpdatedAt) {
		columns = append(columns, "last_updated_at")
	}

	if course.HasParseError != before.HasParseError {
		columns = append(columns, "has_parse_error")
	}

	if course.IsAnnual != before.IsAnnual {
		columns = append(columns, "is_annual")
	}

	dbCourse := timetabledbmodel.ToDBCourse(course, false)

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(columns) > 0 {
			if err := tx.Select(columns).Updates(dbCourse).Error; err != nil {
				return err
			}
		}

		if err := r.updateCourseRecommendedGrades(tx, course); err != nil {
			return err
		}

		if err := r.updateCourseMethods(tx, course); err != nil {
			return err
		}

		if err := r.updateCourseSchedules(tx, course); err != nil {
			return err
		}

		return nil
	}, nil)
}
