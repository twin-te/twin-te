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

	FindAnnouncement(ctx context.Context, conds FindAnnouncementConds, lock sharedport.Lock) (mo.Option[*announcementdomain.Announcement], error)
	ListAnnouncements(ctx context.Context, conds ListAnnouncementsConds, lock sharedport.Lock) ([]*announcementdomain.Announcement, error)
	CreateAnnouncements(ctx context.Context, announcements ...*announcementdomain.Announcement) error

	FindAlreadyRead(ctx context.Context, conds FindAlreadyReadConds, lock sharedport.Lock) (mo.Option[*announcementdomain.AlreadyRead], error)
	ListAlreadyReads(ctx context.Context, conds ListAlreadyReadsConds, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error)
	CreateAlreadyReads(ctx context.Context, alreadyReads ...*announcementdomain.AlreadyRead) error
	DeleteAlreadyReads(ctx context.Context, conds DeleteAlreadyReadsConds) (rowsAffected int, err error)
}

// Announcement

type FindAnnouncementConds struct {
	ID                idtype.AnnouncementID
	PublishedAtBefore mo.Option[time.Time]
}

type ListAnnouncementsConds struct {
	IDs               mo.Option[[]idtype.AnnouncementID]
	PublishedAtBefore mo.Option[time.Time]
}

// AlreadyRead

type FindAlreadyReadConds struct {
	UserID         idtype.UserID
	AnnouncementID idtype.AnnouncementID
}

type ListAlreadyReadsConds struct {
	UserID          mo.Option[idtype.UserID]
	AnnouncementIDs mo.Option[[]idtype.AnnouncementID]
}

type DeleteAlreadyReadsConds struct {
	UserID         mo.Option[idtype.UserID]
	AnnouncementID mo.Option[idtype.AnnouncementID]
}
