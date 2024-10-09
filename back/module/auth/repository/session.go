package authrepository

import (
	"context"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (r *impl) FindSession(ctx context.Context, conds authport.FindSessionConds, lock sharedport.Lock) (mo.Option[*authdomain.Session], error) {
	db := r.db.WithContext(ctx).Where("id = ?", conds.ID.String())

	if expiredAtAfter, ok := conds.ExpiredAtAfter.Get(); ok {
		db = db.Where("expired_at > ?", expiredAtAfter)
	}

	db = dbhelper.ApplyLock(db, lock)

	dbSession := new(authdbmodel.Session)
	if err := db.Take(dbSession).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*authdomain.Session](err)
	}

	return base.SomeWithErr(authdbmodel.FromDBSession(dbSession))
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

	if userID, ok := conds.UserID.Get(); ok {
		db.Where("user_id = ?", userID.String())
	}

	return int(db.Delete(&authdbmodel.Session{}).RowsAffected), db.Error
}
