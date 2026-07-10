package authrepository

import (
	"context"
	"time"

	"github.com/samber/mo"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *impl) CreateAuthChallenge(ctx context.Context, challenge *authdomain.AuthChallenge) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("expired_at <= ?", time.Now()).Delete(&authdbmodel.AuthChallenge{}).Error; err != nil {
			return err
		}
		return tx.Create(authdbmodel.ToDBAuthChallenge(challenge)).Error
	})
}

func (r *impl) ConsumeAuthChallenge(
	ctx context.Context,
	id string,
	provider authdomain.Provider,
	now time.Time,
) (mo.Option[*authdomain.AuthChallenge], error) {
	dbChallenge := new(authdbmodel.AuthChallenge)
	result := r.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Where("provider = ?", provider.String()).
		Where("expired_at > ?", now).
		Delete(dbChallenge)
	if result.Error != nil {
		return mo.None[*authdomain.AuthChallenge](), result.Error
	}
	if result.RowsAffected == 0 {
		return mo.None[*authdomain.AuthChallenge](), nil
	}
	challenge, err := authdbmodel.FromDBAuthChallenge(dbChallenge)
	if err != nil {
		return mo.None[*authdomain.AuthChallenge](), err
	}
	return mo.Some(challenge), nil
}
