package authdomain

import (
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/apperr"
	"github.com/twin-te/twin-te/back/base"
	autherr "github.com/twin-te/twin-te/back/module/auth/err"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
)

// User is identified by one of the following fields.
//   - ID
//   - one of the Authentications
type User struct {
	ID              idtype.UserID
	Authentications []UserAuthentication
	CreatedAt       time.Time

	EntityBeforeUpdated *User
}

func (u *User) Clone() *User {
	ret := lo.ToPtr(*u)
	ret.Authentications = base.CopySlice(u.Authentications)
	return ret
}

func (u *User) BeforeUpdateHook() {
	u.EntityBeforeUpdated = u.Clone()
}

// AddAuthentication adds the given authentication.
//
// [Error Code]
//   - auth.UserHasAtMostOneAuthenticationFromSameProvider
func (u *User) AddAuthentication(newAuthentication UserAuthentication) error {
	for _, authentication := range u.Authentications {
		if newAuthentication.Provider == authentication.Provider {
			return apperr.New(
				autherr.CodeUserHasAtMostOneAuthenticationFromSameProvider,
				fmt.Sprintf("the authentication whose provider is %s is already registered", newAuthentication.Provider),
			)
		}
	}
	u.Authentications = append(u.Authentications, newAuthentication)
	return nil
}

// DeleteAuthentication deletes the authentication specified by the given provider.
//
// [Error Code]
//   - auth.UserHasAtLeastOneAuthentication
func (u *User) DeleteAuthentication(provider Provider) error {
	if len(u.Authentications) == 1 {
		return apperr.New(autherr.CodeUserHasAtLeastOneAuthentication, "")
	}

	for i, authentication := range u.Authentications {
		if provider == authentication.Provider {
			u.Authentications = append(u.Authentications[:i], u.Authentications[i+1:]...)
			return nil
		}
	}

	return apperr.New(
		sharederr.CodeInvalidArgument,
		fmt.Sprintf("the authentication whose provider is %s is not associated with the user whose id is %s", provider, u.ID),
	)
}

func ConstructUser(fn func(u *User) (err error)) (*User, error) {
	u := new(User)
	if err := fn(u); err != nil {
		return nil, err
	}

	if len(u.Authentications) == 0 {
		return nil, errors.New("user must have one or more authentications")
	}

	if u.ID.IsZero() || u.CreatedAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", u)
	}

	return u, nil
}
