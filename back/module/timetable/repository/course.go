package timetablerepository

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) FindCourse(ctx context.Context, conds timetableport.FindCourseConds, lock sharedport.Lock) (*timetabledomain.Course, error) {
	dbCourse := new(model.Course)

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

	return fromDBCourse(dbCourse)
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

	var dbCourses []*model.Course

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

	return base.MapWithErr(dbCourses, fromDBCourse)
}

func (r *impl) CreateCourses(ctx context.Context, courses ...*timetabledomain.Course) error {
	dbCourses := base.MapWithArg(courses, true, toDBCourse)
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
		columns = append(columns, "instructor")
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
		columns = append(columns, "last_update")
	}

	if course.HasParseError != before.HasParseError {
		columns = append(columns, "has_parse_error")
	}

	if course.IsAnnual != before.IsAnnual {
		columns = append(columns, "is_annual")
	}

	dbCourse := toDBCourse(course, false)

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

func fromDBCourse(dbCourse *model.Course) (*timetabledomain.Course, error) {
	return timetabledomain.ConstructCourse(func(c *timetabledomain.Course) (err error) {
		c.ID, err = idtype.ParseCourseID(dbCourse.ID)
		if err != nil {
			return err
		}

		c.Year, err = shareddomain.ParseAcademicYear(int(dbCourse.Year))
		if err != nil {
			return err
		}

		c.Code, err = timetabledomain.ParseCode(dbCourse.Code)
		if err != nil {
			return err
		}

		c.Name, err = timetabledomain.ParseName(dbCourse.Name)
		if err != nil {
			return err
		}

		c.Instructors = dbCourse.Instructors

		c.Credit, err = timetabledomain.ParseCredit(fmt.Sprintf("%.1f", dbCourse.Credit))
		if err != nil {
			return err
		}

		c.Overview = dbCourse.Overview
		c.Remarks = dbCourse.Remarks
		c.LastUpdatedAt = dbCourse.LastUpdatedAt
		c.HasParseError = dbCourse.HasParseError
		c.IsAnnual = dbCourse.IsAnnual

		if c.RecommendedGrades, err = base.MapWithErr(dbCourse.RecommendedGrades, fromDBRecommendedGrade); err != nil {
			return err
		}

		if c.Methods, err = base.MapWithErr(dbCourse.Methods, fromDBCourseMethod); err != nil {
			return err
		}

		if c.Schedules, err = base.MapWithErr(dbCourse.Schedules, fromDBCourseSchedule); err != nil {
			return err
		}

		return nil
	})
}

func toDBCourse(course *timetabledomain.Course, withAssociations bool) *model.Course {
	dbCourse := &model.Course{
		ID:            course.ID.String(),
		Year:          int16(course.Year),
		Code:          course.Code.String(),
		Name:          course.Name.String(),
		Instructors:   course.Instructors,
		Credit:        course.Credit.Float(),
		Overview:      course.Overview,
		Remarks:       course.Remarks,
		LastUpdatedAt: course.LastUpdatedAt,
		HasParseError: course.HasParseError,
		IsAnnual:      course.IsAnnual,
	}

	if withAssociations {
		dbCourse.RecommendedGrades = base.MapWithArg(course.RecommendedGrades, course.ID, toDBRecommendedGrade)
		dbCourse.Methods = base.MapWithArg(course.Methods, course.ID, toDBCourseMethod)
		dbCourse.Schedules = base.MapWithArg(course.Schedules, course.ID, toDBCourseSchedule)
	}

	return dbCourse
}
