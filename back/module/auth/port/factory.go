package authport

import (
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

type Factory interface {
	NewUser(authentication authdomain.UserAuthentication) (*authdomain.User, error)
	NewSession(userID idtype.UserID) (*authdomain.Session, error)
}
