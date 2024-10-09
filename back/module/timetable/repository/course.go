package timetablerepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) FindCourse(ctx context.Context, conds timetableport.FindCourseConds, lock sharedport.Lock) (mo.Option[*timetabledomain.Course], error) {
	dbCourse := new(timetabledbmodel.Course)

	err := r.transaction(ctx, func(tx *gorm.DB) error {
		tx = dbhelper.ApplyLock(tx, lock)
		return tx.
			Where("year = ?", conds.Year.Int()).
			Where("code = ?", conds.Code.String()).
			Preload("RecommendedGrades").
			Preload("Methods").
			Preload("Schedules").
			Take(dbCourse).
			Error
	}, true)
	if err != nil {
		return dbhelper.ConvertErrRecordNotFound[*timetabledomain.Course](err)
	}

	return base.SomeWithErr(timetabledbmodel.FromDBCourse(dbCourse))
}

func (r *impl) ListCourses(ctx context.Context, conds timetableport.ListCoursesConds, lock sharedport.Lock) ([]*timetabledomain.Course, error) {
	var dbCourses []*timetabledbmodel.Course

	err := r.transaction(ctx, func(tx *gorm.DB) error {
		if ids, ok := conds.IDs.Get(); ok {
			tx = tx.Where("id IN ?", base.MapByString(ids))
		}

		if year, ok := conds.Year.Get(); ok {
			tx = tx.Where("year = ?", year.Int())
		}

		if codes, ok := conds.Codes.Get(); ok {
			tx = tx.Where("code IN ?", base.MapByString(codes))
		}

		tx = dbhelper.ApplyLock(tx, lock)

		return tx.
			Preload("RecommendedGrades").
			Preload("Methods").
			Preload("Schedules").
			Find(&dbCourses).
			Error
	}, true)
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbCourses, timetabledbmodel.FromDBCourse)
}

func (r *impl) CreateCourses(ctx context.Context, courses ...*timetabledomain.Course) error {
	dbCourses := base.MapWithArg(courses, true, timetabledbmodel.ToDBCourse)
	dbRecommendedGrades := lo.Flatten(base.Map(dbCourses, func(dbCourse *timetabledbmodel.Course) []timetabledbmodel.CourseRecommendedGrade {
		return dbCourse.RecommendedGrades
	}))
	dbMethods := lo.Flatten(base.Map(dbCourses, func(dbCourse *timetabledbmodel.Course) []timetabledbmodel.CourseMethod {
		return dbCourse.Methods
	}))
	dbSchedules := lo.Flatten(base.Map(dbCourses, func(dbCourse *timetabledbmodel.Course) []timetabledbmodel.CourseSchedule {
		return dbCourse.Schedules
	}))
	return r.transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(dbCourses).Error; err != nil {
			return err
		}
		if len(dbRecommendedGrades) > 0 {
			if err := tx.Create(dbRecommendedGrades).Error; err != nil {
				return err
			}
		}
		if len(dbMethods) > 0 {
			if err := tx.Create(dbMethods).Error; err != nil {
				return err
			}
		}
		if len(dbSchedules) > 0 {
			if err := tx.Create(dbSchedules).Error; err != nil {
				return err
			}
		}
		return nil
	}, false)
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

	return r.transaction(ctx, func(tx *gorm.DB) error {
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
	}, false)
}
