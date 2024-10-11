package timetablequery

import (
	"context"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
	timetabledbmodel "github.com/twin-te/twin-te/back/module/timetable/dbmodel"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
	"gorm.io/gorm"
)

func (q *impl) FindRegisteredCourses(ctx context.Context, id idtype.RegisteredCourseID) (mo.Option[*timetableappdto.RegisteredCourse], error) {
	registeredCourses, err := q.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{IDs: mo.Some([]idtype.RegisteredCourseID{id})})
	if err != nil {
		return mo.None[*timetableappdto.RegisteredCourse](), err
	}

	if len(registeredCourses) == 0 {
		return mo.None[*timetableappdto.RegisteredCourse](), nil
	}

	return mo.Some(registeredCourses[0]), nil
}

func (q *impl) ListRegisteredCourses(ctx context.Context, conds timetableport.ListRegisteredCoursesConds) ([]*timetableappdto.RegisteredCourse, error) {
	var (
		dbRegisteredCourses []*timetabledbmodel.RegisteredCourse
		dbCourses           []*timetabledbmodel.Course
	)

	err := q.gormTransaction(ctx, func(tx *gorm.DB) error {
		if ids, ok := conds.IDs.Get(); ok {
			tx = tx.Where("id IN ?", base.MapByString(ids))
		}

		if userID, ok := conds.UserID.Get(); ok {
			tx = tx.Where("user_id = ?", userID.String())
		}

		if year, ok := conds.Year.Get(); ok {
			tx = tx.Where("year = ?", year.Int())
		}

		if err := tx.
			Preload("Tags").
			Find(&dbRegisteredCourses).
			Error; err != nil {
			return err
		}

		return tx.
			Preload("RecommendedGrades").
			Preload("Methods").
			Preload("Schedules").
			Find(&dbCourses).
			Error
	})
	if err != nil {
		return nil, err
	}

	dbCourseIDToDBCourse := lo.SliceToMap(dbCourses, func(dbCourse *timetabledbmodel.Course) (string, *timetabledbmodel.Course) {
		return dbCourse.ID, dbCourse
	})

	return base.MapWithErr(dbRegisteredCourses, func(dbRegisteredCourse *timetabledbmodel.RegisteredCourse) (*timetableappdto.RegisteredCourse, error) {
		var (
			registeredCourse = new(timetableappdto.RegisteredCourse)
			err              error
		)

		registeredCourse.ID, err = idtype.ParseRegisteredCourseID(dbRegisteredCourse.ID)
		if err != nil {
			return nil, err
		}

		registeredCourse.UserID, err = idtype.ParseUserID(dbRegisteredCourse.UserID)
		if err != nil {
			return nil, err
		}

		registeredCourse.Year, err = shareddomain.ParseAcademicYear(dbRegisteredCourse.Year)
		if err != nil {
			return nil, err
		}

		if dbCourseID, ok := dbRegisteredCourse.CourseID.Get(); ok {
			dbCourse := dbCourseIDToDBCourse[dbCourseID]

			registeredCourse.CourseID, err = base.SomeWithErr(idtype.ParseCourseID(dbCourse.ID))
			if err != nil {
				return nil, err
			}

			registeredCourse.Code, err = base.SomeWithErr(timetabledomain.ParseCode(dbCourse.Code))
			if err != nil {
				return nil, err
			}

			registeredCourse.Name, err = timetabledomain.ParseName(dbRegisteredCourse.Name.OrElse(dbCourse.Name))
			if err != nil {
				return nil, err
			}

			registeredCourse.Credit, err = timetabledomain.ParseCredit(dbRegisteredCourse.Credit.OrElse(dbCourse.Credit))
			if err != nil {
				return nil, err
			}

			registeredCourse.Instructors = dbRegisteredCourse.Instructors.OrElse(dbCourse.Instructors)

			if methods, ok := dbRegisteredCourse.Methods.Get(); ok {
				registeredCourse.Methods, err = timetabledbmodel.FromDBRegisteredCourseMethods(methods)
				if err != nil {
					return nil, err
				}
			} else {
				registeredCourse.Methods, err = base.MapWithErr(dbCourse.Methods, timetabledbmodel.FromDBCourseMethod)
				if err != nil {
					return nil, err
				}
			}

			if schedules, ok := dbhelper.NullToOption(dbRegisteredCourse.Schedules).Get(); ok {
				registeredCourse.Schedules, err = timetabledbmodel.FromDBRegisteredCourseSchedules(schedules)
				if err != nil {
					return nil, err
				}
			} else {
				registeredCourse.Schedules, err = base.MapWithErr(dbCourse.Schedules, timetabledbmodel.FromDBCourseSchedule)
				if err != nil {
					return nil, err
				}
			}
		}

		registeredCourse.Memo = dbRegisteredCourse.Memo

		registeredCourse.Attendance, err = timetabledomain.ParseAttendance(dbRegisteredCourse.Attendance)
		if err != nil {
			return nil, err
		}

		registeredCourse.Absence, err = timetabledomain.ParseAbsence(dbRegisteredCourse.Absence)
		if err != nil {
			return nil, err
		}

		registeredCourse.Late, err = timetabledomain.ParseLate(dbRegisteredCourse.Late)
		if err != nil {
			return nil, err
		}

		registeredCourse.TagIDs, err = base.MapWithErr(dbRegisteredCourse.Tags, timetabledbmodel.FromDBRegisteredCourseTag)
		if err != nil {
			return nil, err
		}

		return registeredCourse, nil
	})
}
