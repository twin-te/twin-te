package schoolcalendarusecase

import (
	authmodule "github.com/twin-te/twinte-back/module/auth"
	schoolcalendarmodule "github.com/twin-te/twinte-back/module/schoolcalendar"
	schoolcalendarport "github.com/twin-te/twinte-back/module/schoolcalendar/port"
)

var _ schoolcalendarmodule.UseCase = (*impl)(nil)

type impl struct {
	a authmodule.AccessController
	r schoolcalendarport.Repository
}

func New(a authmodule.AccessController, r schoolcalendarport.Repository) *impl {
	return &impl{a, r}
}
