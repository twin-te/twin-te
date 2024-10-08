package timetableusecase

import (
	authmodule "github.com/twin-te/twin-te/back/module/auth"
	timetablemodule "github.com/twin-te/twin-te/back/module/timetable"
	timetableport "github.com/twin-te/twin-te/back/module/timetable/port"
)

var _ timetablemodule.UseCase = (*impl)(nil)

type impl struct {
	a authmodule.AccessController
	f timetableport.Factory
	i timetableport.Integrator
	q timetableport.Query
	r timetableport.Repository
}

func New(a authmodule.AccessController, f timetableport.Factory, i timetableport.Integrator, q timetableport.Query, r timetableport.Repository) *impl {
	return &impl{a, f, i, q, r}
}
