package authport

import (
	"context"
	"time"

	"github.com/samber/mo"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindUser(ctx context.Context, filter UserFilter, lock sharedport.Lock) (mo.Option[*authdomain.User], error)
	ListUsers(ctx context.Context, filter UserFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*authdomain.User, error)
	CreateUsers(ctx context.Context, users ...*authdomain.User) error
	UpdateUser(ctx context.Context, user *authdomain.User) error
	DeleteUsers(ctx context.Context, filter UserFilter) (rowsAffected int, err error)

	FindSession(ctx context.Context, filter SessionFilter, lock sharedport.Lock) (mo.Option[*authdomain.Session], error)
	ListSessions(ctx context.Context, filter SessionFilter, limitOffset sharedport.LimitOffset, lock sharedport.Lock) ([]*authdomain.Session, error)
	CreateSessions(ctx context.Context, sessions ...*authdomain.Session) error
	DeleteSessions(ctx context.Context, filter SessionFilter) (rowsAffected int, err error)
}

type UserFilter struct {
	ID                 mo.Option[idtype.UserID]
	UserAuthentication mo.Option[authdomain.UserAuthentication]
}

func (f *UserFilter) IsUniqueFilter() bool {
	return f.ID.IsPresent() || f.UserAuthentication.IsPresent()
}

type SessionFilter struct {
	ID             mo.Option[idtype.SessionID]
	UserID         mo.Option[idtype.UserID]
	ExpiredAtAfter mo.Option[time.Time]
}

func (f *SessionFilter) IsUniqueFilter() bool {
	return f.ID.IsPresent()
}
