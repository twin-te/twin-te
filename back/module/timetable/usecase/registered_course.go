package timetableusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/apperr"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
	sharedhelper "github.com/twin-te/twin-te/back/module/shared/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableerr "github.com/twin-te/twin-te/back/module/timetable/err"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func (uc *impl) CreateRegisteredCoursesByCodes(ctx context.Context, year shareddomain.AcademicYear, codes []timetabledomain.Code) ([]*timetabledomain.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	courses, err := uc.r.ListCourses(ctx, timetableport.ListCoursesConds{
		Year:  &year,
		Codes: &codes,
	}, sharedport.LockNone)
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

	savedRegisteredCourses, err := uc.r.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID:    &userID,
		Year:      &year,
		CourseIDs: &courseIDs,
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	if len(savedRegisteredCourses) != 0 {
		alreadyRegisteredCodes := base.Map(savedRegisteredCourses, func(rc *timetabledomain.RegisteredCourse) timetabledomain.Code {
			return courseIDToCode[*rc.CourseID]
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
	})
	return registeredCourses, err
}

func (uc *impl) CreateRegisteredCourseManually(ctx context.Context, in timetablemodule.CreateRegisteredCourseManuallyIn) (*timetabledomain.RegisteredCourse, error) {
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

	return registeredCourse, uc.r.CreateRegisteredCourses(ctx, registeredCourse)
}

func (uc *impl) GetRegisteredCourses(ctx context.Context, year *shareddomain.AcademicYear) ([]*timetabledomain.RegisteredCourse, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	registeredCourses, err := uc.r.ListRegisteredCourses(ctx, timetableport.ListRegisteredCoursesConds{
		UserID: &userID,
		Year:   year,
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	return registeredCourses, uc.r.LoadCourseAssociationToRegisteredCourse(ctx, registeredCourses, sharedport.LockNone)
}

func (uc *impl) UpdateRegisteredCourse(ctx context.Context, in timetablemodule.UpdateRegisteredCourseIn) (registeredCourse *timetabledomain.RegisteredCourse, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	if in.Methods != nil {
		if err := sharedhelper.ValidateDuplicates(*in.Methods); err != nil {
			return nil, err
		}
	}

	if in.Schedules != nil {
		if err := sharedhelper.ValidateDuplicates(*in.Schedules); err != nil {
			return nil, err
		}
	}

	if in.TagIDs != nil {
		if err := sharedhelper.ValidateDuplicates(*in.TagIDs); err != nil {
			return nil, err
		}
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) (err error) {
		registeredCourse, err = rtx.FindRegisteredCourse(ctx, timetableport.FindRegisteredCourseConds{
			ID:     in.ID,
			UserID: &userID,
		}, sharedport.LockExclusive)
		if err != nil {
			if errors.Is(err, sharedport.ErrNotFound) {
				return apperr.New(timetableerr.CodeRegisteredCourseNotFound, fmt.Sprintf("not found registered course whose id is %s", in.ID))
			}
			return err
		}

		if err := uc.r.LoadCourseAssociationToRegisteredCourse(ctx, []*timetabledomain.RegisteredCourse{registeredCourse}, sharedport.LockNone); err != nil {
			return err
		}

		if in.TagIDs != nil {
			savedTags, err := rtx.ListTags(ctx, timetableport.ListTagsConds{
				IDs:    in.TagIDs,
				UserID: &userID,
			}, sharedport.LockShared)
			if err != nil {
				return err
			}

			savedTagIDs := base.Map(savedTags, func(tag *timetabledomain.Tag) idtype.TagID { return tag.ID })

			notFoundTagIDs, _ := lo.Difference(*in.TagIDs, savedTagIDs)
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
		})

		return rtx.UpdateRegisteredCourse(ctx, registeredCourse)
	})

	return
}

func (uc *impl) DeleteRegisteredCourse(ctx context.Context, id idtype.RegisteredCourseID) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := uc.r.DeleteRegisteredCourses(ctx, timetableport.DeleteRegisteredCoursesConds{
		ID:     &id,
		UserID: &userID,
	})
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperr.New(timetableerr.CodeRegisteredCourseNotFound, fmt.Sprintf("not found registered course whose id is %s", id))
	}

	return nil
}
