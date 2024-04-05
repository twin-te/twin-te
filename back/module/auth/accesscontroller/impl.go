package accesscontroller

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/appctx"
	"github.com/twin-te/twin-te/back/apperr"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
	sharederr "github.com/twin-te/twin-te/back/module/shared/err"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

var _ authmodule.AccessController = (*impl)(nil)

type impl struct {
	r authport.Repository
}

func (a *impl) WithActor(ctx context.Context, id *idtype.SessionID) (context.Context, error) {
	if id == nil {
		return appctx.SetActor(ctx, authdomain.NewUnknown()), nil
	}

	session, err := a.r.FindSession(ctx, authport.FindSessionConds{
		ID:             *id,
		ExpiredAtAfter: lo.ToPtr(time.Now()),
	}, sharedport.LockNone)
	if err != nil {
		if errors.Is(err, sharedport.ErrNotFound) {
			return appctx.SetActor(ctx, authdomain.NewUnknown()), nil
		}
		return nil, err
	}

	user, err := a.r.FindUser(ctx, authport.FindUserConds{
		ID: &session.UserID,
	}, sharedport.LockNone)
	if err != nil {
		if errors.Is(err, sharedport.ErrNotFound) {
			return appctx.SetActor(ctx, authdomain.NewUnknown()), nil
		}
		return nil, err
	}

	return appctx.SetActor(ctx, authdomain.NewAuthNUser(user.ID)), nil
}

func (a *impl) Authenticate(ctx context.Context) (idtype.UserID, error) {
	actor, ok := appctx.GetActor(ctx)
	if !ok {
		return idtype.UserID{}, fmt.Errorf("failed to retrieve actor from the context")
	}

	if authNUser, ok := actor.AuthNUser(); ok {
		return authNUser.UserID, nil
	}

	return idtype.UserID{}, apperr.New(sharederr.CodeUnauthenticated, "")
}

func (*impl) Authorize(ctx context.Context, permission authdomain.Permission) error {
	actor, ok := appctx.GetActor(ctx)
	if !ok {
		return errors.New("failed to retrieve actor from the context")
	}

	if actor.HasPermission(permission) {
		return nil
	}

	return apperr.New(sharederr.CodeUnauthorized, fmt.Sprintf("required permission is %s", permission))
}

func New(r authport.Repository) *impl {
	return &impl{
		r: r,
	}
}
