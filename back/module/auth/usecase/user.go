package authusecase

import (
	"context"
	"fmt"

	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/apperr"
	"github.com/twin-te/twin-te/back/base"
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	autherr "github.com/twin-te/twin-te/back/module/auth/err"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
	sharedport "github.com/twin-te/twin-te/back/module/shared/port"
)

func (uc *impl) SignUpOrLogin(ctx context.Context, userAuthentication authdomain.UserAuthentication) (*authdomain.Session, error) {
	userOption, err := uc.r.FindUser(ctx, authport.UserFilter{
		UserAuthentication: mo.Some(userAuthentication),
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	user, err := base.OptionOrElseByWithErr(userOption, func() (*authdomain.User, error) {
		user, err := uc.f.NewUser(userAuthentication)
		if err != nil {
			return nil, err
		}
		return user, uc.r.CreateUsers(ctx, user)
	})
	if err != nil {
		return nil, err
	}

	session, err := uc.f.NewSession(user.ID)
	if err != nil {
		return nil, err
	}
	return session, uc.r.CreateSessions(ctx, session)
}

func (uc *impl) GetMe(ctx context.Context) (*authdomain.User, error) {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	userOption, err := uc.r.FindUser(ctx, authport.UserFilter{
		ID: mo.Some(userID),
	}, sharedport.LockNone)
	if err != nil {
		return nil, err
	}

	return userOption.MustGet(), nil
}

func (uc *impl) AddUserAuthentication(ctx context.Context, userAuthentication authdomain.UserAuthentication) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	userOption, err := uc.r.FindUser(ctx, authport.UserFilter{
		UserAuthentication: mo.Some(userAuthentication),
	}, sharedport.LockNone)
	if err != nil {
		return err
	}
	if userOption.IsPresent() {
		return apperr.New(autherr.CodeUserAuthenticationAlreadyExists, fmt.Sprintf("the given user authentication already exists, %+v", userAuthentication))
	}

	return uc.r.Transaction(ctx, func(rtx authport.Repository) error {
		userOption, err := rtx.FindUser(ctx, authport.UserFilter{
			ID: mo.Some(userID),
		}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		user := userOption.MustGet()
		user.BeforeUpdateHook()
		if err := user.AddAuthentication(userAuthentication); err != nil {
			return err
		}
		return rtx.UpdateUser(ctx, user)
	}, false)
}

func (uc *impl) DeleteUserAuthentication(ctx context.Context, provider authdomain.Provider) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	return uc.r.Transaction(ctx, func(rtx authport.Repository) error {
		userOption, err := rtx.FindUser(ctx, authport.UserFilter{
			ID: mo.Some(userID),
		}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		user := userOption.MustGet()
		user.BeforeUpdateHook()
		if err := user.DeleteAuthentication(provider); err != nil {
			return err
		}
		return rtx.UpdateUser(ctx, user)
	}, true)
}

func (uc *impl) Logout(ctx context.Context) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	_, err = uc.r.DeleteSessions(ctx, authport.SessionFilter{
		UserID: mo.Some(userID),
	})
	return err
}

func (uc *impl) DeleteAccount(ctx context.Context) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	_, err = uc.r.DeleteUsers(ctx, authport.UserFilter{ID: mo.Some(userID)})
	if err != nil {
		return err
	}

	_, err = uc.r.DeleteSessions(ctx, authport.SessionFilter{
		UserID: mo.Some(userID),
	})
	return err
}
