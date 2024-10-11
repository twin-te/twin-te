package authrepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	authdbmodel "github.com/twin-te/twin-te/back/module/auth/dbmodel"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"gorm.io/gorm"
)

func (r *impl) updateUserAuthentications(db *gorm.DB, user *authdomain.User) error {
	before := user.BeforeUpdated.MustGet()
	toCreate, toDelete := lo.Difference(user.Authentications, before.Authentications)

	if len(toCreate) != 0 {
		dbUserAuthentications := base.MapWithArg(toCreate, user.ID, authdbmodel.ToDBUserAuthentication)

		if err := db.Create(dbUserAuthentications).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbUserAuthentications := base.MapWithArg(toDelete, user.ID, authdbmodel.ToDBUserAuthentication)

		return db.
			Where("user_id = ?", user.ID.String()).
			Where("(provider, social_id) IN ?", base.Map(dbUserAuthentications, func(dbUserAuthentication authdbmodel.UserAuthentication) []any {
				return []any{dbUserAuthentication.Provider, dbUserAuthentication.SocialID}
			})).
			Delete(&authdbmodel.UserAuthentication{}).
			Error
	}

	return nil
}
