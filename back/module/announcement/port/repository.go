package announcementport

import (
	"context"
	"time"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error) error

	FindAnnouncement(ctx context.Context, conds FindAnnouncementConds, lock sharedport.Lock) (*announcementdomain.Announcement, error)
	ListAnnouncements(ctx context.Context, conds ListAnnouncementsConds, lock sharedport.Lock) ([]*announcementdomain.Announcement, error)
	CreateAnnouncements(ctx context.Context, announcements ...*announcementdomain.Announcement) error

	FindAlreadyRead(ctx context.Context, conds FindAlreadyReadConds, lock sharedport.Lock) (*announcementdomain.AlreadyRead, error)
	ListAlreadyReads(ctx context.Context, conds ListAlreadyReadsConds, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error)
	CreateAlreadyReads(ctx context.Context, alreadyReads ...*announcementdomain.AlreadyRead) error
	DeleteAlreadyReads(ctx context.Context, conds DeleteAlreadyReadsConds) (rowsAffected int, err error)
}

// Announcement

type FindAnnouncementConds struct {
	ID                idtype.AnnouncementID
	PublishedAtBefore *time.Time
}

type ListAnnouncementsConds struct {
	IDs               *[]idtype.AnnouncementID
	PublishedAtBefore *time.Time
}

// AlreadyRead

type FindAlreadyReadConds struct {
	UserID         idtype.UserID
	AnnouncementID idtype.AnnouncementID
}

type ListAlreadyReadsConds struct {
	UserID          *idtype.UserID
	AnnouncementIDs *[]idtype.AnnouncementID
}

type DeleteAlreadyReadsConds struct {
	UserID         *idtype.UserID
	AnnouncementID *idtype.AnnouncementID
}
