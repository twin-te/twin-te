package timetableusecase

import (
	authmodule "github.com/twin-te/twinte-back/module/auth"
	timetablemodule "github.com/twin-te/twinte-back/module/timetable"
	timetableport "github.com/twin-te/twinte-back/module/timetable/port"
)

var _ timetablemodule.UseCase = (*impl)(nil)

type impl struct {
	a authmodule.AccessController
	f timetableport.Factory
	g timetableport.Gateway
	r timetableport.Repository
}

func New(a authmodule.AccessController, f timetableport.Factory, g timetableport.Gateway, r timetableport.Repository) *impl {
	return &impl{a, f, g, r}
}
