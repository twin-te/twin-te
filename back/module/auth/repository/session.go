package authrepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm/clause"
)

func (r *impl) FindSession(ctx context.Context, conds authport.FindSessionConds, lock sharedport.Lock) (*authdomain.Session, error) {
	db := r.db.WithContext(ctx).Where("id = ?", conds.ID.String())

	if conds.ExpiredAtAfter != nil {
		db = db.Where("expired_at > ?", *conds.ExpiredAtAfter)
	}

	if lock != sharedport.LockNone {
		db = db.Clauses(clause.Locking{
			Strength: lo.Ternary(lock == sharedport.LockExclusive, "UPDATE", "SHARE"),
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	dbSession := new(authdbmodel.Session)
	if err := db.Take(dbSession).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return authdbmodel.FromDBSession(dbSession)
}

func (r *impl) ListSessions(ctx context.Context, conds authport.ListSessionsConds, lock sharedport.Lock) ([]*authdomain.Session, error) {
	var dbSessions []*authdbmodel.Session

	err := r.db.WithContext(ctx).Find(&dbSessions).Error
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbSessions, authdbmodel.FromDBSession)
}

func (r *impl) CreateSessions(ctx context.Context, sessions ...*authdomain.Session) error {
	dbSessions := base.Map(sessions, authdbmodel.ToDBSession)
	return r.db.WithContext(ctx).Create(dbSessions).Error
}

func (r *impl) DeleteSessions(ctx context.Context, conds authport.DeleteSessionsConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db.Where("user_id = ?", conds.UserID.String())
	}

	return int(db.Delete(&authdbmodel.Session{}).RowsAffected), db.Error
}
