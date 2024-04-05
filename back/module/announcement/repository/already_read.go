package announcementrepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	announcementdomain "github.com/twin-te/twin-te/back/module/announcement/domain"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindAlreadyRead(ctx context.Context, conds announcementport.FindAlreadyReadConds, lock sharedport.Lock) (*announcementdomain.AlreadyRead, error) {
	db := r.db.
		WithContext(ctx).
		Where("read_user = ?", conds.UserID.String()).
		Where("information_id = ?", conds.AnnouncementID.String())

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbAlreadyRead := new(model.AlreadyRead)
	if err := db.Take(dbAlreadyRead).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return fromDBAlreadyRead(dbAlreadyRead)
}

func (r *impl) ListAlreadyReads(ctx context.Context, conds announcementport.ListAlreadyReadsConds, lock sharedport.Lock) ([]*announcementdomain.AlreadyRead, error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db = db.Where("read_user = ?", conds.UserID.String())
	}

	if conds.AnnouncementIDs != nil {
		db = db.Where("information_id IN ?", base.MapByString(*conds.AnnouncementIDs))
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	var dbAlreadyReads []*model.AlreadyRead
	if err := db.Find(&dbAlreadyReads).Error; err != nil {
		return nil, err
	}

	return base.MapWithErr(dbAlreadyReads, fromDBAlreadyRead)
}

func (r *impl) CreateAlreadyReads(ctx context.Context, alreadyReads ...*announcementdomain.AlreadyRead) error {
	dbAlreadyReads := base.Map(alreadyReads, toDBAlreadyRead)
	return r.db.WithContext(ctx).Create(dbAlreadyReads).Error
}

func (r *impl) DeleteAlreadyReads(ctx context.Context, conds announcementport.DeleteAlreadyReadsConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db = db.Where("read_user = ?", conds.UserID.String())
	}

	if conds.AnnouncementID != nil {
		db = db.Where("information_id = ?", conds.AnnouncementID.String())
	}

	return int(db.Delete(&model.AlreadyRead{}).RowsAffected), db.Error
}

func fromDBAlreadyRead(dbAlreadyRead *model.AlreadyRead) (*announcementdomain.AlreadyRead, error) {
	return announcementdomain.ConstructAlreadyRead(func(ar *announcementdomain.AlreadyRead) (err error) {
		ar.ID, err = idtype.ParseAlreadyReadID(dbAlreadyRead.ID)
		if err != nil {
			return
		}

		ar.UserID, err = idtype.ParseUserID(dbAlreadyRead.ReadUser)
		if err != nil {
			return
		}

		ar.AnnouncementID, err = idtype.ParseAnnouncementID(dbAlreadyRead.InformationID)
		if err != nil {
			return
		}

		ar.ReadAt = dbAlreadyRead.ReadAt

		return
	})
}

func toDBAlreadyRead(alreadyRead *announcementdomain.AlreadyRead) *model.AlreadyRead {
	return &model.AlreadyRead{
		ID:            alreadyRead.ID.String(),
		InformationID: alreadyRead.AnnouncementID.String(),
		ReadUser:      alreadyRead.UserID.String(),
		ReadAt:        alreadyRead.ReadAt,
	}
}
