package authdomain

import (
	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

type Actor interface {
	HasPermission(permission Permission) bool
	AuthNUser() (authNUser *AuthNUser, ok bool)
}

var _ Actor = (*Unknown)(nil)

type Unknown struct {
	Permissions []Permission
}

func (actor *Unknown) HasPermission(permission Permission) bool {
	return lo.Contains(actor.Permissions, permission)
}

func (actor *Unknown) AuthNUser() (*AuthNUser, bool) {
	return nil, false
}

var _ Actor = (*AuthNUser)(nil)

// AuthNUser represents the authenticated user.
type AuthNUser struct {
	UserID      idtype.UserID
	Permissions []Permission
}

func (actor *AuthNUser) HasPermission(permission Permission) bool {
	return lo.Contains(actor.Permissions, permission)
}

func (actor *AuthNUser) AuthNUser() (*AuthNUser, bool) {
	return actor, true
}

func NewUnknown(permissions ...Permission) *Unknown {
	return &Unknown{Permissions: permissions}
}

func NewAuthNUser(userID idtype.UserID, permissions ...Permission) *AuthNUser {
	return &AuthNUser{UserID: userID, Permissions: permissions}
}
