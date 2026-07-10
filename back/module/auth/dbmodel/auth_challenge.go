package authdbmodel

import (
	"time"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
)

type AuthChallenge struct {
	ID        string
	Provider  string
	Nonce     string
	ExpiredAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDBAuthChallenge(challenge *AuthChallenge) (*authdomain.AuthChallenge, error) {
	provider, err := authdomain.ParseProvider(challenge.Provider)
	if err != nil {
		return nil, err
	}
	return &authdomain.AuthChallenge{
		ID:        challenge.ID,
		Provider:  provider,
		Nonce:     challenge.Nonce,
		ExpiredAt: challenge.ExpiredAt,
	}, nil
}

func ToDBAuthChallenge(challenge *authdomain.AuthChallenge) *AuthChallenge {
	return &AuthChallenge{
		ID:        challenge.ID,
		Provider:  challenge.Provider.String(),
		Nonce:     challenge.Nonce,
		ExpiredAt: challenge.ExpiredAt,
	}
}
