package donationrepository

import (
	"context"

	donationport "github.com/twin-te/twinte-back/module/donation/port"
	"gorm.io/gorm"
)

var _ donationport.Repository = (*impl)(nil)

type impl struct {
	db *gorm.DB
}

func (r *impl) Transaction(ctx context.Context, fn func(rtx donationport.Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&impl{db: tx})
	}, nil)
}

func New(db *gorm.DB) *impl {
	return &impl{
		db: db,
	}
}
