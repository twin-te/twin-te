package accesscontroller

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/mo"
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

func (a *impl) WithActor(ctx context.Context, id mo.Option[idtype.SessionID]) (context.Context, error) {
	if id.IsAbsent() {
		return appctx.SetActor(ctx, authdomain.NewUnknown()), nil
	}

	sessionOption, err := a.r.FindSession(ctx, authport.SessionFilter{
		ID:             id,
		ExpiredAtAfter: mo.Some(time.Now()),
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	session, found := sessionOption.Get()
	if !found {
		return appctx.SetActor(ctx, authdomain.NewUnknown()), nil
	}

	userOption, err := a.r.FindUser(ctx, authport.UserFilter{
		ID: mo.Some(session.UserID),
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	user, found := userOption.Get()
	if !found {
		return appctx.SetActor(ctx, authdomain.NewUnknown()), nil
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
