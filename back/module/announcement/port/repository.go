package announcementport

import (
	"context"
	"time"

	"github.com/samber/mo"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindAlreadyRead(ctx context.Context, filter AlreadyReadFilter, lock sharedport.Lock) (mo.Option[*announcementdomain.AlreadyRead], error)
	ListAlreadyReads(ctx context.Context, filter AlreadyReadFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error)
	CreateAlreadyReads(ctx context.Context, alreadyReads ...*announcementdomain.AlreadyRead) error
	DeleteAlreadyReads(ctx context.Context, filter AlreadyReadFilter) (rowsAffected int, err error)

	FindAnnouncement(ctx context.Context, filter AnnouncementFilter, lock sharedport.Lock) (mo.Option[*announcementdomain.Announcement], error)
	ListAnnouncements(ctx context.Context, filter AnnouncementFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*announcementdomain.Announcement, error)
	CreateAnnouncements(ctx context.Context, announcements ...*announcementdomain.Announcement) error
}

type AlreadyReadFilter struct {
	UserID          mo.Option[idtype.UserID]
	AnnouncementID  mo.Option[idtype.AnnouncementID]
	AnnouncementIDs mo.Option[[]idtype.AnnouncementID]
}

func (f *AlreadyReadFilter) IsUniqueFilter() bool {
	return f.UserID.IsPresent() && f.AnnouncementID.IsPresent()
}

type AnnouncementFilter struct {
	ID                mo.Option[idtype.AnnouncementID]
	IDs               mo.Option[[]idtype.AnnouncementID]
	PublishedAtBefore mo.Option[time.Time]
}

func (f *AnnouncementFilter) IsUniqueFilter() bool {
	return f.ID.IsPresent()
}
