package schoolcalendarrepository

import (
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	schoolcalendarport "github.com/twin-te/twin-te/back/module/schoolcalendar/port"
)

var _ schoolcalendarport.Repository = (*impl)(nil)

type impl struct {
	events        []*schoolcalendardomain.Event
	moduleDetails []*schoolcalendardomain.ModuleDetail
}

func New() *impl {
	return &impl{}
}
