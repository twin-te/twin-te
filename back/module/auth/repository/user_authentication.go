package authrepository

import (
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	"github.com/twin-te/twin-te/back/db/gen/model"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	"gorm.io/gorm"
)

func (r *impl) updateUserAuthentications(db *gorm.DB, user *authdomain.User) error {
	toCreate, toDelete := lo.Difference(user.Authentications, user.EntityBeforeUpdated.Authentications)

	if len(toCreate) != 0 {
		dbUserAuthentications := base.MapWithArg(toCreate, user.ID, toDBUserAuthentication)

		if err := db.Create(dbUserAuthentications).Error; err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {
		dbUserAuthentications := base.MapWithArg(toDelete, user.ID, toDBUserAuthentication)

		return db.Where("user_id = ?", user.ID.String()).
			Where("(provider,social_id) IN ?", base.Map(dbUserAuthentications, func(dbUserAuthentication model.UserAuthentication) []any {
				return []any{dbUserAuthentication.Provider, dbUserAuthentication.SocialID}
			})).
			Delete(&model.UserAuthentication{}).
			Error
	}

	return nil
}

func fromDBUserAuthentication(dbUserAuthentication model.UserAuthentication) (userAuthentication authdomain.UserAuthentication, err error) {
	provider, err := authdomain.ParseProvider(dbUserAuthentication.Provider)
	if err != nil {
		return
	}

	socialID, err := authdomain.ParseSocialID(dbUserAuthentication.SocialID)
	if err != nil {
		return
	}

	userAuthentication = authdomain.NewUserAuthentication(provider, socialID)

	return
}

func toDBUserAuthentication(userAuthentication authdomain.UserAuthentication, userID idtype.UserID) model.UserAuthentication {
	return model.UserAuthentication{
		UserID:   lo.ToPtr(userID.String()),
		Provider: userAuthentication.Provider.String(),
		SocialID: userAuthentication.SocialID.String(),
	}
}
