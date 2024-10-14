package announcementrepository

import (
	"context"
	"fmt"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	announcementdbmodel "github.com/twin-te/twin-te/back/module/announcement/dbmodel"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm"
)

func (r *impl) FindAlreadyRead(ctx context.Context, filter announcementport.AlreadyReadFilter, lock sharedport.Lock) (mo.Option[*announcementdomain.AlreadyRead], error) {
	if !filter.IsUniqueFilter() {
		return mo.None[*announcementdomain.AlreadyRead](), fmt.Errorf("%v is not unique", filter)
	}

	db := r.db.WithContext(ctx)
	db = applyAlreadyReadFilter(db, filter)
	db = dbhelper.ApplyLock(db, lock)

	dbAlreadyRead := new(announcementdbmodel.AlreadyRead)
	if err := db.Take(dbAlreadyRead).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*announcementdomain.AlreadyRead](err)
	}

	return base.SomeWithErr(announcementdbmodel.FromDBAlreadyRead(dbAlreadyRead))
}

func (r *impl) ListAlreadyReads(ctx context.Context, filter announcementport.AlreadyReadFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error) {
	db := r.db.WithContext(ctx)
	db = applyAlreadyReadFilter(db, filter)
	db = dbhelper.ApplyLimitOffset(db, limitOffset)
	db = dbhelper.ApplyLock(db, lock)

	var dbAlreadyReads []*announcementdbmodel.AlreadyRead
	if err := db.Find(&dbAlreadyReads).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbAlreadyReads, announcementdbmodel.FromDBAlreadyRead)
}

func (r *impl) CreateAlreadyReads(ctx context.Context, alreadyReads ...*announcementdomain.AlreadyRead) error {
	dbAlreadyReads := base.Map(alreadyReads, announcementdbmodel.ToDBAlreadyRead)
	return r.transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(dbAlreadyReads).Error
	}, false)
}

func (r *impl) DeleteAlreadyReads(ctx context.Context, filter announcementport.AlreadyReadFilter) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)
	db = applyAlreadyReadFilter(db, filter)
	return int(db.Delete(&announcementdbmodel.AlreadyRead{}).RowsAffected), db.Error
}

func applyAlreadyReadFilter(db *gorm.DB, filter announcementport.AlreadyReadFilter) *gorm.DB {
	if userID, ok := filter.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if announcementID, ok := filter.AnnouncementID.Get(); ok {
		db = db.Where("announcement_id = ?", announcementID.String())
	}

	if announcementIDs, ok := filter.AnnouncementIDs.Get(); ok {
		db = db.Where("announcement_id IN ?", base.MapByString(announcementIDs))
	}

	return db
}
