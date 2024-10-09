package donationrepository

import (
	"context"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	donationdbmodel "github.com/twin-te/twin-te/back/module/donation/dbmodel"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (r *impl) FindPaymentUser(ctx context.Context, conds donationport.FindPaymentUserConds, lock sharedport.Lock) (mo.Option[*donationdomain.PaymentUser], error) {
	db := r.db.WithContext(ctx).
		Where("user_id = ?", conds.UserID.String())

	db = dbhelper.ApplyLock(db, lock)

	dbPaymentUser := new(donationdbmodel.PaymentUser)
	if err := db.Take(&dbPaymentUser).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*donationdomain.PaymentUser](err)
	}

	return base.SomeWithErr(donationdbmodel.FromDBPaymentUser(dbPaymentUser))
}

func (r *impl) ListPaymentUsers(ctx context.Context, conds donationport.ListPaymentUsersConds, lock sharedport.Lock) ([]*donationdomain.PaymentUser, error) {
	db := r.db.WithContext(ctx)

	if conds.RequireDisplayName {
		db = db.Where("display_name IS NOT NULL")
	}

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

	if len(columns) == 0 {
		return nil
	}

	dbPaymentUser := donationdbmodel.ToDBPaymentUser(paymentUser)
	return r.db.
		WithContext(ctx).
		Select(columns).
		Updates(dbPaymentUser).
		Error
}
