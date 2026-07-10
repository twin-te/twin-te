package authusecase

import (
	"context"
	"errors"
	"time"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

func (uc *impl) CreateAuthChallenge(ctx context.Context, provider authdomain.Provider) (*authdomain.AuthChallenge, error) {
	challenge, err := uc.f.NewAuthChallenge(provider)
	if err != nil {
		return nil, err
	}
	if err := uc.r.CreateAuthChallenge(ctx, challenge); err != nil {
		return nil, err
	}
	return challenge, nil
}

func (uc *impl) ConsumeAuthChallenge(
	ctx context.Context,
	id string,
	provider authdomain.Provider,
) (*authdomain.AuthChallenge, error) {
	challenge, err := uc.r.ConsumeAuthChallenge(ctx, id, provider, time.Now())
	if err != nil {
		return nil, err
	}
	if value, ok := challenge.Get(); ok {
		return value, nil
	}
	return nil, errors.New("authentication challenge is missing, expired, or already used")
}
