package announcementrepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	announcementdbmodel "github.com/twin-te/twin-te/back/module/announcement/dbmodel"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindAlreadyRead(ctx context.Context, conds announcementport.FindAlreadyReadConds, lock sharedport.Lock) (*announcementdomain.AlreadyRead, error) {
	db := r.db.
		WithContext(ctx).
		Where("user_id = ?", conds.UserID.String()).
		Where("announcement_id = ?", conds.AnnouncementID.String())

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbAlreadyRead := new(announcementdbmodel.AlreadyRead)
	if err := db.Take(dbAlreadyRead).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return announcementdbmodel.FromDBAlreadyRead(dbAlreadyRead)
}

func (r *impl) ListAlreadyReads(ctx context.Context, conds announcementport.ListAlreadyReadsConds, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	if conds.AnnouncementIDs != nil {
		db = db.Where("announcement_id IN ?", base.MapByString(*conds.AnnouncementIDs))
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

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

	if conds.UserID != nil {
		db = db.Where("user_id = ?", conds.UserID.String())
	}

	if conds.AnnouncementID != nil {
		db = db.Where("announcement_id = ?", conds.AnnouncementID.String())
	}

	return int(db.Delete(&announcementdbmodel.AlreadyRead{}).RowsAffected), db.Error
}
