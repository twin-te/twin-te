package timetablerepository

import (
	"context"
	_ "embed"

	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
	"gorm.io/gorm"
)

var _ timetableport.Repository = (*impl)(nil)

type impl struct {
	db *gorm.DB
}

func (r *impl) Transaction(ctx context.Context, fn func(rtx timetableport.Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&impl{db: tx})
	}, nil)
}

func New(db *gorm.DB) *impl {
	return &impl{
		db: db,
	}
}
