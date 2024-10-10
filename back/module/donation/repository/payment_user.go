package donationrepository

import (
	"context"
	"fmt"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	donationdbmodel "github.com/twin-te/twin-te/back/module/donation/dbmodel"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm"
)

func (r *impl) FindPaymentUser(ctx context.Context, filter donationport.PaymentUserFilter, lock sharedport.Lock) (mo.Option[*donationdomain.PaymentUser], error) {
	if !filter.IsUniqueFilter() {
		return mo.None[*donationdomain.PaymentUser](), fmt.Errorf("%v is not unique", filter)
	}

	db := r.db.WithContext(ctx)
	db = applyPaymentUserFilter(db, filter)
	db = dbhelper.ApplyLock(db, lock)

	dbPaymentUser := new(donationdbmodel.PaymentUser)
	if err := db.Take(&dbPaymentUser).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*donationdomain.PaymentUser](err)
	}

	return base.SomeWithErr(donationdbmodel.FromDBPaymentUser(dbPaymentUser))
}

func (r *impl) ListPaymentUsers(ctx context.Context, filter donationport.PaymentUserFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*donationdomain.PaymentUser, error) {
	db := r.db.WithContext(ctx)
	db = applyPaymentUserFilter(db, filter)
	db = dbhelper.ApplyLimitOffset(db, limitOffset)
	db = dbhelper.ApplyLock(db, lock)

	var dbPaymentUsers []*donationdbmodel.PaymentUser
	if err := db.Find(&dbPaymentUsers).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbPaymentUsers, donationdbmodel.FromDBPaymentUser)
}

func (r *impl) CreatePaymentUsers(ctx context.Context, paymentUsers ...*donationdomain.PaymentUser) error {
	dbPaymentUsers := base.Map(paymentUsers, donationdbmodel.ToDBPaymentUser)
	return r.db.WithContext(ctx).Create(dbPaymentUsers).Error
}

func (r *impl) UpdatePaymentUser(ctx context.Context, paymentUser *donationdomain.PaymentUser) error {
	before := paymentUser.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if paymentUser.DisplayName != before.DisplayName {
		columns = append(columns, "display_name")
	}

	if paymentUser.Link != before.Link {
		columns = append(columns, "link")
	}

	dbPaymentUser := donationdbmodel.ToDBPaymentUser(paymentUser)
	return r.db.
		WithContext(ctx).
		Select(columns).
		Updates(dbPaymentUser).
		Error
}

func applyPaymentUserFilter(db *gorm.DB, filter donationport.PaymentUserFilter) *gorm.DB {
	if userID, ok := filter.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if filter.RequireDisplayName {
		db = db.Where("display_name IS NOT NULL")
	}

	return db
}
