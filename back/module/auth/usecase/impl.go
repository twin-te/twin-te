package authusecase

import (
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	authport "github.com/twin-te/twin-te/back/module/auth/port"
)

var _ authmodule.UseCase = (*impl)(nil)

type impl struct {
	a                      authmodule.AccessController
	f                      authport.Factory
	r                      authport.Repository
	appleCredentialRevoker authport.AppleCredentialRevoker
}

func New(
	a authmodule.AccessController,
	f authport.Factory,
	r authport.Repository,
	appleCredentialRevokers ...authport.AppleCredentialRevoker,
) *impl {
	var appleCredentialRevoker authport.AppleCredentialRevoker
	if len(appleCredentialRevokers) != 0 {
		appleCredentialRevoker = appleCredentialRevokers[0]
	}
	return &impl{
		a:                      a,
		f:                      f,
		r:                      r,
		appleCredentialRevoker: appleCredentialRevoker,
	}
}
