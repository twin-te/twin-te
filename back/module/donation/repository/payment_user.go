package donationrepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	donationdomain "github.com/twin-te/twin-te/back/module/donation/domain"
	donationport "github.com/twin-te/twin-te/back/module/donation/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindPaymentUser(ctx context.Context, conds donationport.FindPaymentUserConds, lock sharedport.Lock) (*donationdomain.PaymentUser, error) {
	db := r.db.WithContext(ctx).
		Where("twinte_user_id = ?", conds.UserID.String())

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbPaymentUser := new(model.PaymentUser)
	if err := db.Take(&dbPaymentUser).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return fromDBPaymentUser(dbPaymentUser)
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

	var dbPaymentUsers []*model.PaymentUser
	if err := db.Find(&dbPaymentUsers).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbPaymentUsers, fromDBPaymentUser)
}

func (r *impl) CreatePaymentUsers(ctx context.Context, paymentUsers ...*donationdomain.PaymentUser) error {
	dbPaymentUsers := base.Map(paymentUsers, toDBPaymentUser)
	return r.db.WithContext(ctx).Create(dbPaymentUsers).Error
}

func (r *impl) UpdatePaymentUser(ctx context.Context, paymentUser *donationdomain.PaymentUser) error {
	cols := make([]string, 0)

	if !base.EqualPtr(paymentUser.DisplayName, paymentUser.EntityBeforeUpdated.DisplayName) {
		cols = append(cols, "display_name")
	}

	if !base.EqualPtr(paymentUser.Link, paymentUser.EntityBeforeUpdated.Link) {
		cols = append(cols, "link")
	}

	if len(cols) == 0 {
		return nil
	}

	dbPaymentUser := toDBPaymentUser(paymentUser)
	return r.db.WithContext(ctx).
		Select(cols).
		Updates(dbPaymentUser).
		Error
}

func fromDBPaymentUser(dbPaymentUser *model.PaymentUser) (*donationdomain.PaymentUser, error) {
	return donationdomain.ConstructPaymentUser(func(pu *donationdomain.PaymentUser) (err error) {
		pu.ID, err = idtype.ParsePaymentUserID(dbPaymentUser.ID)
		if err != nil {
			return
		}

		pu.UserID, err = idtype.ParseUserID(dbPaymentUser.TwinteUserID)
		if err != nil {
			return
		}

		pu.DisplayName = dbPaymentUser.DisplayName
		pu.Link = dbPaymentUser.Link

		return
	})
}

func toDBPaymentUser(paymentUser *donationdomain.PaymentUser) *model.PaymentUser {
	return &model.PaymentUser{
		ID:           paymentUser.ID.String(),
		TwinteUserID: paymentUser.UserID.String(),
		DisplayName:  paymentUser.DisplayName,
		Link:         paymentUser.Link,
	}
}
