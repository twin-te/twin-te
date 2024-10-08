package authdbmodel

import (
	"time"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type Session struct {
	ID        string
	UserID    string
	ExpiredAt time.Time
}

func FromDBSession(dbSession *Session) (*authdomain.Session, error) {
	return authdomain.ConstructSession(func(s *authdomain.Session) (err error) {
		s.ID, err = idtype.ParseSessionID(dbSession.ID)
		if err != nil {
			return err
		}

		s.UserID, err = idtype.ParseUserID(dbSession.UserID)
		if err != nil {
			return err
		}

		s.ExpiredAt = dbSession.ExpiredAt

		return nil
	})
}

func ToDBSession(session *authdomain.Session) *Session {
	return &Session{
		ID:        session.ID.String(),
		UserID:    session.UserID.String(),
		ExpiredAt: session.ExpiredAt,
	}
}
