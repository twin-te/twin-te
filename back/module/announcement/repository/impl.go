package announcementrepository

import (
	"context"
	"database/sql"

	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	"gorm.io/gorm"
)

var _ announcementport.Repository = (*impl)(nil)

type impl struct {
	db            *gorm.DB
	inTransaction bool
	readOnly      bool

	announcements []*announcementdomain.Announcement
}

func (r *impl) Transaction(ctx context.Context, fn func(rtx announcementport.Repository) error, readOnly bool) error {
	return r.gormTransaction(ctx, func(tx *gorm.DB) error {
		return fn(&impl{db: tx, inTransaction: true, readOnly: readOnly})
	}, readOnly)
}

func (r *impl) transaction(ctx context.Context, fn func(tx *gorm.DB) error, readOnly bool) error {
	if r.inTransaction && r.readOnly && !readOnly {
		panic("invalid implementation")
	}
	if r.inTransaction {
		return fn(r.db)
	}
	return r.gormTransaction(ctx, fn, readOnly)
}

func (r *impl) gormTransaction(ctx context.Context, fn func(tx *gorm.DB) error, readOnly bool) error {
	return r.db.WithContext(ctx).Transaction(fn, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  readOnly,
	})
}

func New(db *gorm.DB) *impl {
	return &impl{
		db:            db,
		announcements: make([]*announcementdomain.Announcement, 0),
	}
}
