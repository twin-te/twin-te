package timetablegateway

import timetableport "github.com/twin-te/twin-te/back/module/timetable/port"

var _ timetableport.Gateway = (*impl)(nil)

type impl struct {
	kdbJSONFilePath string
}

func New(kdbJSONFilePath string) *impl {
	return &impl{
		kdbJSONFilePath: kdbJSONFilePath,
	}
}
