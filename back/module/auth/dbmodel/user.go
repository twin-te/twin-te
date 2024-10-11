package authdbmodel

import (
	"time"

	"github.com/twin-te/twin-te/back/base"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type User struct {
	ID                  string
	UserAuthentications []UserAuthentication

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAuthentication struct {
	UserID   string
	Provider string
	SocialID string
}

func FromDBUser(dbUser *User) (*authdomain.User, error) {
	return authdomain.ConstructUser(func(u *authdomain.User) (err error) {
		u.ID, err = idtype.ParseUserID(dbUser.ID)
		if err != nil {
			return err
		}

		u.Authentications, err = base.MapWithErr(dbUser.UserAuthentications, FromDBUserAuthentication)
		if err != nil {
			return err
		}

		return nil
	})
}

func ToDBUser(user *authdomain.User, withAssociations bool) *User {
	dbUser := &User{
		ID: user.ID.String(),
	}

	if withAssociations {
		dbUser.UserAuthentications = base.MapWithArg(user.Authentications, user.ID, ToDBUserAuthentication)
	}

	return dbUser
}

func FromDBUserAuthentication(dbUserAuthentication UserAuthentication) (userAuthentication authdomain.UserAuthentication, err error) {
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

func ToDBUserAuthentication(userAuthentication authdomain.UserAuthentication, userID idtype.UserID) UserAuthentication {
	return UserAuthentication{
		UserID:   userID.String(),
		Provider: userAuthentication.Provider.String(),
		SocialID: userAuthentication.SocialID.String(),
	}
}
