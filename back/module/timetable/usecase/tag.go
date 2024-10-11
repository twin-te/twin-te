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
	sharedhelper "github.com/twin-te/twin-te/back/module/shared/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
	timetableerr "github.com/twin-te/twin-te/back/module/timetable/err"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

func (uc impl) CreateTag(ctx context.Context, name shareddomain.RequiredString) (*timetabledomain.Tag, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	tag, err := uc.f.NewTag(
		userID,
		name,
	)
	if err != nil {
		return nil, err
	}

	return tag, uc.r.CreateTags(ctx, tag)
}

func (uc impl) ListTags(ctx context.Context) ([]*timetabledomain.Tag, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	return uc.r.ListTags(ctx, timetableport.TagFilter{
		UserID: mo.Some(userID),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
}

func (uc impl) UpdateTag(ctx context.Context, in timetablemodule.UpdateTagIn) (tag *timetabledomain.Tag, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		tagOption, err := rtx.FindTag(ctx, timetableport.TagFilter{
			ID:     mo.Some(in.ID),
			UserID: mo.Some(userID),
		}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		var found bool
		tag, found = tagOption.Get()
		if !found {
			return apperr.New(timetableerr.CodeTagNotFound, fmt.Sprintf("not found tag whose id is %s", in.ID))
		}

		tag.BeforeUpdateHook()
		tag.Update(timetabledomain.TagDataToUpdate{Name: in.Name})
		return rtx.UpdateTag(ctx, tag)
	}, false)

	return
}

func (uc impl) DeleteTag(ctx context.Context, id idtype.TagID) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	var rowsAffected int
	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		registeredCourses, err := rtx.ListRegisteredCourses(ctx, timetableport.RegisteredCourseFilter{
			UserID: mo.Some(userID),
			TagID:  mo.Some(id),
		}, sharedport.LimitOffset{}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		for _, registeredCourse := range registeredCourses {
			registeredCourse.BeforeUpdateHook()
			registeredCourse.DetachTag(id)
			if err := rtx.UpdateRegisteredCourse(ctx, registeredCourse); err != nil {
				return err
			}
		}

		rowsAffected, err = rtx.DeleteTags(ctx, timetableport.TagFilter{
			ID:     mo.Some(id),
			UserID: mo.Some(userID),
		})
		return err
	}, false)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperr.New(timetableerr.CodeTagNotFound, fmt.Sprintf("not found tag whose id is %s", id))
	}

	return nil
}

func (uc impl) RearrangeTags(ctx context.Context, ids []idtype.TagID) (tags []*timetabledomain.Tag, err error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return
	}

	if err = sharedhelper.ValidateDuplicates(ids); err != nil {
		return
	}

	err = uc.r.Transaction(ctx, func(rtx timetableport.Repository) error {
		tags, err = rtx.ListTags(ctx, timetableport.TagFilter{
			UserID: mo.Some(userID),
		}, sharedport.LimitOffset{}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		savedTagIDs := base.Map(tags, func(tag *timetabledomain.Tag) idtype.TagID {
			return tag.ID
		})

		if err := sharedhelper.ValidateDifference(savedTagIDs, ids); err != nil {
			return err
		}

		lo.ForEach(tags, func(tag *timetabledomain.Tag, _ int) {
			tag.BeforeUpdateHook()
		})

		timetabledomain.RearrangeTags(tags, ids)

		for _, tag := range tags {
			if err := rtx.UpdateTag(ctx, tag); err != nil {
				return err
			}
		}

		return nil
	}, false)

	return
}
