package announcementusecase

import (
	announcementmodule "github.com/twin-te/twin-te/back/module/announcement"
	announcementport "github.com/twin-te/twin-te/back/module/announcement/port"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
)

var _ announcementmodule.UseCase = (*impl)(nil)

type impl struct {
	a authmodule.AccessController
	f announcementport.Factory
	r announcementport.Repository
}

func New(a authmodule.AccessController, f announcementport.Factory, r announcementport.Repository) *impl {
	return &impl{a, f, r}
}
