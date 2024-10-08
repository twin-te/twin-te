package donationrepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	donationdbmodel "github.com/twin-te/twin-te/back/module/donation/dbmodel"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindPaymentUser(ctx context.Context, conds donationport.FindPaymentUserConds, lock sharedport.Lock) (*donationdomain.PaymentUser, error) {
	db := r.db.WithContext(ctx).
		Where("user_id = ?", conds.UserID.String())

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbPaymentUser := new(donationdbmodel.PaymentUser)
	if err := db.Take(&dbPaymentUser).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return donationdbmodel.FromDBPaymentUser(dbPaymentUser)
}

func (r *impl) ListPaymentUsers(ctx context.Context, conds donationport.ListPaymentUsersConds, lock sharedport.Lock) ([]*donationdomain.PaymentUser, error) {
	db := r.db.WithContext(ctx)

	if conds.RequireDisplayName {
		db = db.Where("display_name IS NOT NULL")
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

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
