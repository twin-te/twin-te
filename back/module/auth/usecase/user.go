package authusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/twin-te/twinte-back/apperr"
	authdomain "github.com/twin-te/twinte-back/module/auth/domain"
	autherr "github.com/twin-te/twinte-back/module/auth/err"
	authport "github.com/twin-te/twinte-back/module/auth/port"
	sharedport "github.com/twin-te/twinte-back/module/shared/port"
)

func (uc *impl) SignUpOrLogin(ctx context.Context, userAuthentication authdomain.UserAuthentication) (*authdomain.Session, error) {
	user, err := uc.r.FindUser(ctx, authport.FindUserConds{
		UserAuthentication: &userAuthentication,
	}, sharedport.LockNone)

	if errors.Is(err, sharedport.ErrNotFound) {
		user, err = uc.f.NewUser(userAuthentication)
		if err != nil {
			return nil, err
		}
		err = uc.r.CreateUsers(ctx, user)
	}

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

	return uc.r.FindUser(ctx, authport.FindUserConds{
		ID: &userID,
	}, sharedport.LockNone)
}

func (uc *impl) AddUserAuthentication(ctx context.Context, userAuthentication authdomain.UserAuthentication) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	_, err = uc.r.FindUser(ctx, authport.FindUserConds{
		UserAuthentication: &userAuthentication,
	}, sharedport.LockNone)
	if !errors.Is(err, sharedport.ErrNotFound) {
		if err != nil {
			return err
		}
		return apperr.New(autherr.CodeUserAuthenticationAlreadyExists, fmt.Sprintf("the given user authentication already exists, %+v", userAuthentication))
	}

	return uc.r.Transaction(ctx, func(rtx authport.Repository) error {
		user, err := rtx.FindUser(ctx, authport.FindUserConds{
			ID: &userID,
		}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		user.BeforeUpdateHook()
		if err := user.AddAuthentication(userAuthentication); err != nil {
			return err
		}
		return rtx.UpdateUser(ctx, user)
	})
}

func (uc *impl) DeleteUserAuthentication(ctx context.Context, provider authdomain.Provider) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	return uc.r.Transaction(ctx, func(rtx authport.Repository) error {
		user, err := rtx.FindUser(ctx, authport.FindUserConds{
			ID: &userID,
		}, sharedport.LockExclusive)
		if err != nil {
			return err
		}

		user.BeforeUpdateHook()
		if err := user.DeleteAuthentication(provider); err != nil {
			return err
		}
		return rtx.UpdateUser(ctx, user)
	})
}

func (uc *impl) Logout(ctx context.Context) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	_, err = uc.r.DeleteSessions(ctx, authport.DeleteSessionsConds{
		UserID: &userID,
	})
	return err
}

func (uc *impl) DeleteAccount(ctx context.Context) error {
	userID, err := uc.a.Authenticate(ctx)
	if err != nil {
		return err
	}

	_, err = uc.r.DeleteUsers(ctx, authport.DeleteUserConds{ID: &userID})
	if err != nil {
		return err
	}

	_, err = uc.r.DeleteSessions(ctx, authport.DeleteSessionsConds{
		UserID: &userID,
	})
	return err
}
