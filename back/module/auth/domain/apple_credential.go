package authdomain

import "github.com/twin-te/twin-te/back/module/shared/domain/idtype"

type AppleCredential struct {
	UserID       idtype.UserID
	ClientID     string
	RefreshToken string
}
