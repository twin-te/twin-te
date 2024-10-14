package authrepository

import (
	"context"
	"fmt"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
	"gorm.io/gorm"
)

func (r *impl) FindSession(ctx context.Context, filter authport.SessionFilter, lock sharedport.Lock) (mo.Option[*authdomain.Session], error) {
	if !filter.IsUniqueFilter() {
		return mo.None[*authdomain.Session](), fmt.Errorf("%v is not unique", filter)
	}

	db := r.db.WithContext(ctx)
	db = applySessionFilter(db, filter)
	db = dbhelper.ApplyLock(db, lock)

	dbSession := new(authdbmodel.Session)
	if err := db.Take(dbSession).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*authdomain.Session](err)
	}

	return base.SomeWithErr(authdbmodel.FromDBSession(dbSession))
}

func (r *impl) ListSessions(ctx context.Context, filter authport.SessionFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*authdomain.Session, error) {
	db := r.db.WithContext(ctx)
	db = applySessionFilter(db, filter)
	db = dbhelper.ApplyLimitOffset(db, limitOffset)
	db = dbhelper.ApplyLock(db, lock)

	var dbSessions []*authdbmodel.Session
	err := db.Find(&dbSessions).Error
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbSessions, authdbmodel.FromDBSession)
}

func (r *impl) CreateSessions(ctx context.Context, sessions ...*authdomain.Session) error {
	dbSessions := base.Map(sessions, authdbmodel.ToDBSession)
	return r.transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(dbSessions).Error
	}, false)
}

func (r *impl) DeleteSessions(ctx context.Context, filter authport.SessionFilter) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)
	db = applySessionFilter(db, filter)
	return int(db.Delete(&authdbmodel.Session{}).RowsAffected), db.Error
}

func applySessionFilter(db *gorm.DB, filter authport.SessionFilter) *gorm.DB {
	if id, ok := filter.ID.Get(); ok {
		db = db.Where("id = ?", id.String())
	}

	if userID, ok := filter.UserID.Get(); ok {
		db.Where("user_id = ?", userID.String())
	}

	if expiredAtAfter, ok := filter.ExpiredAtAfter.Get(); ok {
		db = db.Where("expired_at > ?", expiredAtAfter)
	}

	return db
}
