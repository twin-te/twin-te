package announcementrepository

import (
	"context"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	announcementdbmodel "github.com/twin-te/twin-te/back/module/announcement/dbmodel"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (r *impl) FindAlreadyRead(ctx context.Context, conds announcementport.FindAlreadyReadConds, lock sharedport.Lock) (mo.Option[*announcementdomain.AlreadyRead], error) {
	db := r.db.
		WithContext(ctx).
		Where("user_id = ?", conds.UserID.String()).
		Where("announcement_id = ?", conds.AnnouncementID.String())

	db = dbhelper.ApplyLock(db, lock)

	dbAlreadyRead := new(announcementdbmodel.AlreadyRead)
	if err := db.Take(dbAlreadyRead).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*announcementdomain.AlreadyRead](err)
	}

	return base.SomeWithErr(announcementdbmodel.FromDBAlreadyRead(dbAlreadyRead))
}

func (r *impl) ListAlreadyReads(ctx context.Context, conds announcementport.ListAlreadyReadsConds, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error) {
	db := r.db.WithContext(ctx)

	if userID, ok := conds.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if announcementIDs, ok := conds.AnnouncementIDs.Get(); ok {
		db = db.Where("announcement_id IN ?", base.MapByString(announcementIDs))
	}

	db = dbhelper.ApplyLock(db, lock)

	var dbAlreadyReads []*announcementdbmodel.AlreadyRead
	if err := db.Find(&dbAlreadyReads).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbAlreadyReads, announcementdbmodel.FromDBAlreadyRead)
}

func (r *impl) CreateAlreadyReads(ctx context.Context, alreadyReads ...*announcementdomain.AlreadyRead) error {
	dbAlreadyReads := base.Map(alreadyReads, announcementdbmodel.ToDBAlreadyRead)
	return r.db.WithContext(ctx).Create(dbAlreadyReads).Error
}

func (r *impl) DeleteAlreadyReads(ctx context.Context, conds announcementport.DeleteAlreadyReadsConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if userID, ok := conds.UserID.Get(); ok {
		db = db.Where("user_id = ?", userID.String())
	}

	if announcementID, ok := conds.AnnouncementID.Get(); ok {
		db = db.Where("announcement_id = ?", announcementID.String())
	}

	return int(db.Delete(&announcementdbmodel.AlreadyRead{}).RowsAffected), db.Error
}
