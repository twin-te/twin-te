package announcementrepository

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (r *impl) FindAnnouncement(ctx context.Context, filter announcementport.AnnouncementFilter, lock sharedport.Lock) (mo.Option[*announcementdomain.Announcement], error) {
	if !filter.IsUniqueFilter() {
		return mo.None[*announcementdomain.Announcement](), fmt.Errorf("%v is not unique", filter)
	}

	announcements := applyAnnouncementFilter(r.announcements, filter)
	if len(announcements) == 0 {
		return mo.None[*announcementdomain.Announcement](), nil
	}

	return mo.Some(announcements[0].Clone()), nil
}

func (r *impl) ListAnnouncements(ctx context.Context, filter announcementport.AnnouncementFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*announcementdomain.Announcement, error) {
	announcements := applyAnnouncementFilter(r.announcements, filter)
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

func applyAnnouncementFilter(announcements []*announcementdomain.Announcement, filter announcementport.AnnouncementFilter) []*announcementdomain.Announcement {
	if id, ok := filter.ID.Get(); ok {
		announcements = lo.Filter(announcements, func(announcement *announcementdomain.Announcement, _ int) bool {
			return announcement.ID == id
		})
	}

	if ids, ok := filter.IDs.Get(); ok {
		announcements = lo.Filter(announcements, func(announcement *announcementdomain.Announcement, _ int) bool {
			return lo.Contains(ids, announcement.ID)
		})
	}

	if publishedAtBefore, ok := filter.PublishedAtBefore.Get(); ok {
		announcements = lo.Filter(announcements, func(announcement *announcementdomain.Announcement, _ int) bool {
			return announcement.PublishedAt.Before(publishedAtBefore)
		})
	}

	return announcements
}
