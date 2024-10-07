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
	before := paymentUser.BeforeUpdated.MustGet()
	columns := make([]string, 0)

	if !base.EqualPtr(paymentUser.DisplayName, before.DisplayName) {
		columns = append(columns, "display_name")
	}

	if !base.EqualPtr(paymentUser.Link, before.Link) {
		columns = append(columns, "link")
	}

	if len(columns) == 0 {
		return nil
	}

	dbPaymentUser := toDBPaymentUser(paymentUser)
	return r.db.WithContext(ctx).
		Select(columns).
		Updates(dbPaymentUser).
		Error
}

func fromDBPaymentUser(dbPaymentUser *model.PaymentUser) (*donationdomain.PaymentUser, error) {
	return donationdomain.ConstructPaymentUser(func(pu *donationdomain.PaymentUser) (err error) {
		pu.ID, err = idtype.ParsePaymentUserID(dbPaymentUser.ID)
		if err != nil {
			return
		}

		pu.UserID, err = idtype.ParseUserID(dbPaymentUser.UserID)
		if err != nil {
			return
		}

		if dbPaymentUser.DisplayName != nil {
			pu.DisplayName, err = base.ToPtrWithErr(donationdomain.ParseDisplayName(*dbPaymentUser.DisplayName))
			if err != nil {
				return
			}
		}

		if dbPaymentUser.Link != nil {
			pu.Link, err = base.ToPtrWithErr(donationdomain.ParseLink(*dbPaymentUser.Link))
			if err != nil {
				return
			}
		}

		return
	})
}

func toDBPaymentUser(paymentUser *donationdomain.PaymentUser) *model.PaymentUser {
	return &model.PaymentUser{
		ID:          paymentUser.ID.String(),
		UserID:      paymentUser.UserID.String(),
		DisplayName: paymentUser.DisplayName.StringPtr(),
		Link:        paymentUser.Link.StringPtr(),
	}
}
