package timetablequery

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func (q *impl) gormTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return q.db.WithContext(ctx).Transaction(fn, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  true,
	})
}

func New(db *gorm.DB) *impl {
	return &impl{db: db}
}
