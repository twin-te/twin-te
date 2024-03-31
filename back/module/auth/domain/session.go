package authdomain

import (
	"fmt"
	"time"

	"github.com/twin-te/twinte-back/appenv"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
)

var SessionLifeTime time.Duration

func init() {
	SessionLifeTime = time.Duration(appenv.SESSION_LIFE_TIME_DAYS) * 24 * time.Hour
}

// Session is identified by one of the following fields.
//   - ID
type Session struct {
	ID        idtype.SessionID
	UserID    idtype.UserID
	ExpiredAt time.Time
}

func ConstructSession(fn func(s *Session) (err error)) (*Session, error) {
	s := new(Session)
	if err := fn(s); err != nil {
		return nil, err
	}

	if s.ID.IsZero() || s.UserID.IsZero() || s.ExpiredAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", s)
	}

	return s, nil
}
