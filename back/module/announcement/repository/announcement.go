package announcementrepository

import (
	"context"
	"fmt"
	"slices"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (r *impl) FindAnnouncement(ctx context.Context, conds announcementport.FindAnnouncementConds, lock sharedport.Lock) (*announcementdomain.Announcement, error) {
	announcement, ok := lo.Find(r.announcements, func(announcement *announcementdomain.Announcement) bool {
		return conds.ID == announcement.ID
	})
	if !ok {
		return nil, sharedport.ErrNotFound
	}

	if conds.PublishedAtBefore != nil && !announcement.PublishedAt.Before(*conds.PublishedAtBefore) {
		return nil, sharedport.ErrNotFound
	}

	return announcement.Clone(), nil
}

func (r *impl) ListAnnouncements(ctx context.Context, conds announcementport.ListAnnouncementsConds, lock sharedport.Lock) ([]*announcementdomain.Announcement, error) {
	announcements := r.announcements

	if conds.IDs != nil {
		announcements = lo.Filter(announcements, func(announcement *announcementdomain.Announcement, _ int) bool {
			return slices.Contains(*conds.IDs, announcement.ID)
		})
	}

	if conds.PublishedAtBefore != nil {
		announcements = lo.Filter(announcements, func(announcement *announcementdomain.Announcement, _ int) bool {
			return announcement.PublishedAt.Before(*conds.PublishedAtBefore)
		})
	}

	return base.MapByClone(announcements), nil
}

func (r *impl) CreateAnnouncements(ctx context.Context, announcements ...*announcementdomain.Announcement) error {
	ids := base.Map(announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	})

	savedIDs := base.Map(r.announcements, func(announcement *announcementdomain.Announcement) idtype.AnnouncementID {
		return announcement.ID
	})

	intersect := lo.Intersect(ids, savedIDs)
	if len(intersect) != 0 {
		return fmt.Errorf("duplicate ids: %+v", intersect)
	}

	r.announcements = append(r.announcements, announcements...)

	return nil
}
