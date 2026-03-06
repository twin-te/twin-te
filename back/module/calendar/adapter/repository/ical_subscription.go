package calendarrepository

import (
	"context"

	"github.com/samber/mo"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	"github.com/twin-te/twin-te/back/module/calendar/dbmodel"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm"
)

func (r *impl) FindIcalSubscriptionByID(ctx context.Context, id idtype.IcalSubscriptionID, lock sharedport.Lock) (mo.Option[idtype.UserID], error) {
	db := r.db.WithContext(ctx).Where("id = ?", id.String())
	db = dbhelper.ApplyLock(db, lock)

	dbSub := new(calendardbmodel.IcalSubscription)
	if err := db.Take(dbSub).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[idtype.UserID](err)
	}
	_, userID, err := calendardbmodel.FromDBIcalSubscription(dbSub.ID, dbSub.UserID)
	if err != nil {
		return mo.None[idtype.UserID](), err
	}
	return mo.Some(userID), nil
}

func (r *impl) FindIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID, lock sharedport.Lock) (mo.Option[idtype.IcalSubscriptionID], error) {
	db := r.db.WithContext(ctx).Where("user_id = ?", userID.String())
	db = dbhelper.ApplyLock(db, lock)

	dbSub := new(calendardbmodel.IcalSubscription)
	if err := db.Take(dbSub).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[idtype.IcalSubscriptionID](err)
	}
	subID, _, err := calendardbmodel.FromDBIcalSubscription(dbSub.ID, dbSub.UserID)
	if err != nil {
		return mo.None[idtype.IcalSubscriptionID](), err
	}
	return mo.Some(subID), nil
}

func (r *impl) CreateIcalSubscription(ctx context.Context, id idtype.IcalSubscriptionID, userID idtype.UserID) error {
	dbSub := calendardbmodel.ToDBIcalSubscription(id, userID)
	return r.transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(dbSub).Error
	}, false)
}

func (r *impl) DeleteIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID) (rowsAffected int, err error) {
	var rows int
	err = r.transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Where("user_id = ?", userID.String()).Delete(&calendardbmodel.IcalSubscription{})
		rows = int(result.RowsAffected)
		return result.Error
	}, false)
	return rows, err
}
