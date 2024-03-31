package authusecase

import (
	authmodule "github.com/twin-te/twinte-back/module/auth"
	authport "github.com/twin-te/twinte-back/module/auth/port"
)

var _ authmodule.UseCase = (*impl)(nil)

type impl struct {
	a authmodule.AccessController
	f authport.Factory
	r authport.Repository
}

func New(a authmodule.AccessController, f authport.Factory, r authport.Repository) *impl {
	return &impl{a, f, r}
}
