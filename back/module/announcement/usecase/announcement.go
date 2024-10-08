package announcementusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/apperr"
	"github.com/twin-te/twin-te/back/base"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementerr "github.com/twin-te/twin-te/back/module/announcement/err"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedhelper "github.com/twin-te/twin-te/back/module/shared/helper"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) ListAnnouncements(ctx context.Context) (announcements []*announcementdomain.Announcement, idToReadFlag map[idtype.AnnouncementID]bool, err error) {
	announcements, err = uc.r.ListAnnouncements(ctx, announcementport.AnnouncementFilter{
		PublishedAtBefore: mo.Some(time.Now()),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return
	}

	userID, unauthenticatedErr := uc.a.Authenticate(ctx)
	if unauthenticatedErr != nil {
		return
	}

	ids := base.Map(announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	})

	alreadyReads, err := uc.r.ListAlreadyReads(ctx, announcementport.AlreadyReadFilter{
		UserID:          mo.Some(userID),
		AnnouncementIDs: mo.Some(ids),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return
	}

	idToReadFlag = lo.SliceToMap(ids, func(id idtype.AnnouncementID) (idtype.AnnouncementID, bool) {
		return id, false
	})

	for _, alreadyRead := range alreadyReads {
		idToReadFlag[alreadyRead.AnnouncementID] = true
	}

	return
}

func (uc *impl) ReadAnnouncements(ctx context.Context, ids []idtype.AnnouncementID) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	if err := sharedhelper.ValidateDuplicates((ids)); err != nil {
		return err
	}

	announcements, err := uc.r.ListAnnouncements(ctx, announcementport.AnnouncementFilter{
		IDs:               mo.Some(ids),
		PublishedAtBefore: mo.Some(time.Now()),
	}, sharedport.LimitOffset{}, sharedport.LockShared)
	if err != nil {
		return err
	}

	_, notFoundIDs := lo.Difference(base.Map(announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	}), ids)
	if len(notFoundIDs) != 0 {
		return apperr.New(announcementerr.CodeAnnouncementNotFound, fmt.Sprintf("not found announcements whose id are, %+v", notFoundIDs))
	}

	savedAlreadyReads, err := uc.r.ListAlreadyReads(ctx, announcementport.AlreadyReadFilter{
		UserID:          mo.Some(userID),
		AnnouncementIDs: mo.Some(ids),
	}, sharedport.LimitOffset{}, sharedport.LockNone)
	if err != nil {
		return err
	}

	targetAnnouncementIDs, _ := lo.Difference(
		ids,
		base.Map(savedAlreadyReads, func(savedAlreadyRead *announcementdomain.AlreadyRead) idtype.AnnouncementID {
			return savedAlreadyRead.AnnouncementID
		}),
	)

	alreadyReads, err := base.MapWithErr(targetAnnouncementIDs, func(id idtype.AnnouncementID) (*announcementdomain.AlreadyRead, error) {
		return uc.f.NewAlreadyRead(userID, id)
	})
	if err != nil {
		return err
	}

	return uc.r.CreateAlreadyReads(ctx, alreadyReads...)
}
