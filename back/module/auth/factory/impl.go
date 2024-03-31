package authfactory

import (
	"time"

	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	authport "github.com/twin-te/twinte-back/module/auth/port"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

var _ authport.Factory = (*impl)(nil)

type impl struct {
	nowFunc func() time.Time
}

func (f *impl) NewUser(authentication authdomain.UserAuthentication) (*authdomain.User, error) {
	return authdomain.ConstructUser(func(u *authdomain.User) (err error) {
		u.ID = idtype.NewUserID()
		u.Authentications = []authdomain.UserAuthentication{authentication}
		u.CreatedAt = f.nowFunc()
		return nil
	})
}

func (f *impl) NewSession(userID idtype.UserID) (*authdomain.Session, error) {
	return authdomain.ConstructSession(func(s *authdomain.Session) error {
		s.ID = idtype.NewSessionID()
		s.UserID = userID
		s.ExpiredAt = f.nowFunc().Add(authdomain.SessionLifeTime)
		return nil
	})
}

func New(nowFunc func() time.Time) *impl {
	return &impl{
		nowFunc: nowFunc,
	}
}
