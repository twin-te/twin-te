package announcementrepository

import (
	"context"

	announcementdomain "github.com/twin-te/twinte-back/module/announcement/domain"
	announcementport "github.com/twin-te/twinte-back/module/announcement/port"
	"gorm.io/gorm"
)

var _ announcementport.Repository = (*impl)(nil)

type impl struct {
	db *gorm.DB

	announcements []*announcementdomain.Announcement
}

func (r *impl) Transaction(ctx context.Context, fn func(rtx announcementport.Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&impl{db: tx})
	}, nil)
}

func New(db *gorm.DB) *impl {
	return &impl{
		db: db,
	}
}
