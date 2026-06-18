package calendarrepository

import (
	"context"

	"github.com/samber/mo"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	calendardbmodel "github.com/twin-te/twin-te/back/module/calendar/dbmodel"
	calendardomain "github.com/twin-te/twin-te/back/module/calendar/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm"
)

func (r *impl) FindIcalSubscriptionByID(ctx context.Context, id idtype.IcalSubscriptionID, lock sharedport.Lock) (mo.Option[*calendardomain.IcalSubscription], error) {
	db := r.db.WithContext(ctx).Preload("TargetTags").Where("id = ?", id.String())
	db = dbhelper.ApplyLock(db, lock)
	return takeIcalSubscription(db)
}

func (r *impl) FindIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID, lock sharedport.Lock) (mo.Option[*calendardomain.IcalSubscription], error) {
	db := r.db.WithContext(ctx).Preload("TargetTags").Where("user_id = ?", userID.String())
	db = dbhelper.ApplyLock(db, lock)
	return takeIcalSubscription(db)
}

func takeIcalSubscription(db *gorm.DB) (mo.Option[*calendardomain.IcalSubscription], error) {
	dbSub := new(calendardbmodel.IcalSubscription)
	if err := db.Take(dbSub).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*calendardomain.IcalSubscription](err)
	}
	sub, err := calendardbmodel.FromDBIcalSubscription(dbSub)
	if err != nil {
		return mo.None[*calendardomain.IcalSubscription](), err
	}
	return mo.Some(sub), nil
}

func (r *impl) CreateIcalSubscription(ctx context.Context, s *calendardomain.IcalSubscription) error {
	dbSub := calendardbmodel.ToDBIcalSubscription(s)
	return r.transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(dbSub).Error
	}, false)
}

func (r *impl) UpdateIcalSubscription(ctx context.Context, s *calendardomain.IcalSubscription) error {
	dbSub := calendardbmodel.ToDBIcalSubscription(s)
	return r.transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Model(&calendardbmodel.IcalSubscription{}).
			Where("id = ?", dbSub.ID).
			Update("mode", dbSub.Mode).Error; err != nil {
			return err
		}

		if err := tx.Where("ical_subscription_id = ?", dbSub.ID).
			Delete(&calendardbmodel.IcalSubscriptionTargetTag{}).Error; err != nil {
			return err
		}

		if len(dbSub.TargetTags) != 0 {
			if err := tx.Create(&dbSub.TargetTags).Error; err != nil {
				return err
			}
		}

		return nil
	}, false)
}

func (r *impl) DeleteIcalSubscriptionByUserID(ctx context.Context, userID idtype.UserID) (rowsAffected int, err error) {
	err = r.transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Where("user_id = ?", userID.String()).Delete(&calendardbmodel.IcalSubscription{})
		rowsAffected = int(result.RowsAffected)
		return result.Error
	}, false)
	return
}
