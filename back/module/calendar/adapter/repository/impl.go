package calendarrepository

import (
	"context"
	"database/sql"

	calendarport "github.com/twin-te/twin-te/back/module/calendar/port"
	"gorm.io/gorm"
)

var _ calendarport.Repository = (*impl)(nil)

type impl struct {
	db            *gorm.DB
	inTransaction bool
	readOnly      bool
}

func (r *impl) Transaction(ctx context.Context, fn func(rtx calendarport.Repository) error, readOnly bool) error {
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
		db: db,
	}
}
