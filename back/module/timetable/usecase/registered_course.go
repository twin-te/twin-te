package timetableusecase

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/apperr"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
	sharedhelper "github.com/twin-te/twin-te/back/module/shared/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetableappdto "github.com/twin-te/twin-te/back/module/timetable/appdto"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableerr "github.com/twin-te/twin-te/back/module/timetable/err"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func (uc *impl) CreateRegisteredCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetableappdto.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	courses, err := uc.r.ListCourses(ctx, timetableport.CourseFilter{
		Year:  mo.Some(year),
		Codes: mo.Some(codes),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	courseIDs := base.Map(courses, func(course *timetabledomain.Course) idtype.CourseID {
		return course.ID
	})
	codeToCourse := lo.SliceToMap(courses, func(course *timetabledomain.Course) (timetabledomain.Code, *timetabledomain.Course) {
		return course.Code, course
	})
	courseIDToCode := lo.SliceToMap(courses, func(course *timetabledomain.Course) (idtype.CourseID, timetabledomain.Code) {
		return course.ID, course.Code
	})

	notFoundCodes := lo.Filter(codes, func(code timetabledomain.Code, index int) bool {
		_, ok := codeToCourse[code]
		return !ok
	})
	if len(notFoundCodes) != 0 {
		return nil, apperr.New(
			timetableerr.CodeCourseNotFound,
			fmt.Sprintf("not found courses with these codes %+v", notFoundCodes),
		)
	}

	savedRegisteredCourses, err := uc.r.ListRegisteredCourses(ctx, timetableport.RegisteredCourseFilter{
		UserID:    mo.Some(userID),
		CourseIDs: mo.Some(courseIDs),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	if len(savedRegisteredCourses) != 0 {
		alreadyRegisteredCodes := base.Map(savedRegisteredCourses, func(rc *timetabledomain.RegisteredCourse) timetabledomain.Code {
			return courseIDToCode[rc.CourseID.MustGet()]
		})

		return nil, apperr.New(
			timetableerr.CodeRegisteredCourseAlreadyExists,
			fmt.Sprintf("the courses with these codes are already registered, %+v", alreadyRegisteredCodes),
		)
	}

	registeredCourses, err := base.MapWithErr(codes, func(code timetabledomain.Code) (*timetabledomain.RegisteredCourse, error) {
		return uc.f.NewRegisteredCourseFromCourse(userID, codeToCourse[code])
	})
	if err != nil {
		return nil, err
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		return rtx.CreateRegisteredCourses(ctx, registeredCourses...)
	}, false)
	if err != nil {
		return nil, err
	}

	registeredCoursesIDs := base.Map(registeredCourses, func(registeredCourse *timetabledomain.RegisteredCourse) idtype.RegisteredCourseID {
		return registeredCourse.ID
	})

	return uc.q.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		IDs: mo.Some(registeredCoursesIDs),
	})
}

func (uc *impl) CreateRegisteredCourseManually(ctx context.Context, in timetablemodule.CreateRegisteredCourseManuallyIn) (*timetableappdto.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	registeredCourse, err := uc.f.NewRegisteredCourseMannualy(
		userID,
		in.Year,
		in.Name,
		in.Instructors,
		in.Credit,
		in.Methods,
		in.Schedules,
	)
	if err != nil {
		return nil, err
	}

	err = uc.r.CreateRegisteredCourses(ctx, registeredCourse)
	if err != nil {
		return nil, err
	}

	return base.MustGetWithErr(uc.q.FindRegisteredCourses(ctx, registeredCourse.ID))
}

func (uc *impl) ListRegisteredCourses(ctx context.Context, year mo.Option[shareddomain.AcademicYear]) ([]*timetableappdto.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	return uc.q.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID: mo.Some(userID),
		Year:   year,
	})
}

func (uc *impl) UpdateRegisteredCourse(ctx context.Context, in timetablemodule.UpdateRegisteredCourseIn) (*timetableappdto.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	if methods, ok := in.Methods.Get(); ok {
		if err := sharedhelper.ValidateDuplicates(methods); err != nil {
			return nil, err
		}
	}

	if schedules, ok := in.Schedules.Get(); ok {
		if err := sharedhelper.ValidateDuplicates(schedules); err != nil {
			return nil, err
		}
	}

	if tagIDs, ok := in.TagIDs.Get(); ok {
		if err := sharedhelper.ValidateDuplicates(tagIDs); err != nil {
			return nil, err
		}
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) (err error) {
		registeredCourseOption, err := rtx.FindRegisteredCourse(ctx, timetableport.RegisteredCourseFilter{
			ID:     mo.Some(in.ID),
			UserID: mo.Some(userID),
		}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		registeredCourse, found := registeredCourseOption.Get()
		if !found {
			return apperr.New(timetableerr.CodeRegisteredCourseNotFound, fmt.Sprintf("not found registered course whose id is %s", in.ID))
		}

		var courseOption mo.Option[*timetabledomain.Course]
		if courseID, ok := registeredCourse.CourseID.Get(); ok {
			courseOption, err = rtx.FindCourse(ctx, timetableport.CourseFilter{ID: mo.Some(courseID)}, sharedport.LockShared)
			if err != nil {
				return nil
			}
		}

		if tagIDs, ok := in.TagIDs.Get(); ok {
			savedTags, err := rtx.ListTags(ctx, timetableport.TagFilter{
				IDs:    in.TagIDs,
				UserID: mo.Some(userID),
			}, sharedport.LimitOffset{}, sharedport.LockShared)
			if err != nil {
				return err
			}

			savedTagIDs := base.Map(savedTags, func(tag *timetabledomain.Tag) idtype.TagID { return tag.ID })

			notFoundTagIDs, _ := lo.Difference(tagIDs, savedTagIDs)
			if len(notFoundTagIDs) != 0 {
				return apperr.New(sharederr.CodeInvalidArgument, fmt.Sprintf("invalid tag ids %+v", notFoundTagIDs))
			}
		}

		registeredCourse.BeforeUpdateHook()

		registeredCourse.Update(timetabledomain.RegisteredCourseDataToUpdate{
			Name:        in.Name,
			Instructors: in.Instructors,
			Credit:      in.Credit,
			Methods:     in.Methods,
			Schedules:   in.Schedules,
			Memo:        in.Memo,
			Attendance:  in.Attendance,
			Absence:     in.Absence,
			Late:        in.Late,
			TagIDs:      in.TagIDs,
		}, courseOption)

		return rtx.UpdateRegisteredCourse(ctx, registeredCourse)
	}, false)
	if err != nil {
		return nil, err
	}

	return base.MustGetWithErr(uc.q.FindRegisteredCourses(ctx, in.ID))
}

func (uc *impl) DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := uc.r.DeleteRegisteredCourses(ctx, timetableport.RegisteredCourseFilter{
		ID:     mo.Some(id),
		UserID: mo.Some(userID),
	})
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperr.New(timetableerr.CodeRegisteredCourseNotFound, fmt.Sprintf("not found registered course whose id is %s", id))
	}

	return nil
}
