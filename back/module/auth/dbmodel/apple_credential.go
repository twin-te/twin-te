package authdbmodel

import (
	"time"

	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

type AppleCredential struct {
	UserID       string
	ClientID     string
	RefreshToken string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDBAppleCredential(value *AppleCredential) (*authdomain.AppleCredential, error) {
	userID, err := idtype.ParseUserID(value.UserID)
	if err != nil {
		return nil, err
	}
	return &authdomain.AppleCredential{
		UserID:       userID,
		ClientID:     value.ClientID,
		RefreshToken: value.RefreshToken,
	}, nil
}

func ToDBAppleCredential(value *authdomain.AppleCredential) *AppleCredential {
	return &AppleCredential{
		UserID:       value.UserID.String(),
		ClientID:     value.ClientID,
		RefreshToken: value.RefreshToken,
	}
}
