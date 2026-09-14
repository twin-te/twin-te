package authdomain_test

import (
	authdomain "github.com/twin-te/twin-te/back/module/auth/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func newTestUserAuthentication() authdomain.UserAuthentication {
	return authdomain.NewUserAuthentication(authdomain.ProviderGoogle, "social-id")
}

func newTestUser() *authdomain.User {
	return &authdomain.User{
		ID:              idtype.NewUserID(),
		Authentications: []authdomain.UserAuthentication{newTestUserAuthentication()},
	}
}
