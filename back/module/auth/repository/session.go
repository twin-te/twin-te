package authrepository

import (
	"context"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
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

	dbSession := new(model.Session)
	if err := db.Take(dbSession).Error; err != nil {
		return nil, dbhelper.ConvertErrRecordNotFound(err)
	}

	return fromDBSession(dbSession)
}

func (r *impl) ListSessions(ctx context.Context, conds authport.ListSessionsConds, lock sharedport.Lock) ([]*authdomain.Session, error) {
	var dbSessions []*model.Session

	err := r.db.WithContext(ctx).Find(&dbSessions).Error
	if err != nil {
		return nil, err
	}

	return base.MapWithErr(dbSessions, fromDBSession)
}

func (r *impl) CreateSessions(ctx context.Context, sessions ...*authdomain.Session) error {
	dbSessions := base.Map(sessions, toDBSession)
	return r.db.WithContext(ctx).Create(dbSessions).Error
}

func (r *impl) DeleteSessions(ctx context.Context, conds authport.DeleteSessionsConds) (rowsAffected int, err error) {
	db := r.db.WithContext(ctx)

	if conds.UserID != nil {
		db.Where("user_id = ?", conds.UserID.String())
	}

	return int(db.Delete(&model.Session{}).RowsAffected), db.Error
}

func fromDBSession(dbSession *model.Session) (*authdomain.Session, error) {
	return authdomain.ConstructSession(func(s *authdomain.Session) (err error) {
		s.ID, err = idtype.ParseSessionID(dbSession.ID)
		if err != nil {
			return err
		}

		s.UserID, err = idtype.ParseUserID(dbSession.UserID)
		if err != nil {
			return err
		}

		s.ExpiredAt = dbSession.ExpiredAt

		return nil
	})
}

func toDBSession(session *authdomain.Session) *model.Session {
	return &model.Session{
		ID:        session.ID.String(),
		UserID:    session.UserID.String(),
		ExpiredAt: session.ExpiredAt,
	}
}
