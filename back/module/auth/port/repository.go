package authport

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/mo"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

type Repository interface {
	Transaction(ctx context.Context, fn func(rtx Repository) error, readOnly bool) error

	FindUser(ctx context.Context, conds FindUserConds, lock sharedport.Lock) (mo.Option[*authdomain.User], error)
	ListUsers(ctx context.Context, conds ListUsersConds, lock sharedport.Lock) ([]*authdomain.User, error)
	CreateUsers(ctx context.Context, users ...*authdomain.User) error
	UpdateUser(ctx context.Context, user *authdomain.User) error
	DeleteUsers(ctx context.Context, conds DeleteUserConds) (rowsAffected int, err error)

	FindSession(ctx context.Context, conds FindSessionConds, lock sharedport.Lock) (mo.Option[*authdomain.Session], error)
	ListSessions(ctx context.Context, conds ListSessionsConds, lock sharedport.Lock) ([]*authdomain.Session, error)
	CreateSessions(ctx context.Context, sessions ...*authdomain.Session) error
	DeleteSessions(ctx context.Context, conds DeleteSessionsConds) (rowsAffected int, err error)
}

// User

type FindUserConds struct {
	ID                 mo.Option[idtype.UserID]
	UserAuthentication mo.Option[authdomain.UserAuthentication]
}

func (conds FindUserConds) Validate() error {
	if conds.ID.IsAbsent() && conds.UserAuthentication.IsAbsent() {
		return fmt.Errorf("invalid %v", conds)
	}
	return nil
}

type ListUsersConds struct{}

type DeleteUserConds struct {
	ID mo.Option[idtype.UserID]
}

// Session

type FindSessionConds struct {
	ID             idtype.SessionID
	ExpiredAtAfter mo.Option[time.Time]
}

type ListSessionsConds struct{}

type DeleteSessionsConds struct {
	UserID mo.Option[idtype.UserID]
}
