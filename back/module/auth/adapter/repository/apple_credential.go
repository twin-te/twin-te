package authrepository

import (
	"context"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	dbhelper "github.com/twin-te/twin-te/back/db/helper"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	"gorm.io/gorm/clause"
)

func (r *impl) FindAppleCredential(
	ctx context.Context,
	userID idtype.UserID,
) (mo.Option[*authdomain.AppleCredential], error) {
	dbCredential := new(authdbmodel.AppleCredential)
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID.String()).Take(dbCredential).Error; err != nil {
		return dbhelper.ConvertErrRecordNotFound[*authdomain.AppleCredential](err)
	}
	return base.SomeWithErr(authdbmodel.FromDBAppleCredential(dbCredential))
}

func (r *impl) SaveAppleCredential(ctx context.Context, credential *authdomain.AppleCredential) error {
	dbCredential := authdbmodel.ToDBAppleCredential(credential)
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"client_id", "refresh_token", "updated_at"}),
	}).Create(dbCredential).Error
}

func (r *impl) DeleteAppleCredential(ctx context.Context, userID idtype.UserID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID.String()).
		Delete(&authdbmodel.AppleCredential{}).
		Error
}
