package schoolcalendardomain

import (
	"fmt"

	"cloud.google.com/go/civil"
	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Module -trimprefix=Module -output=module_string.gen.go
type Module int

func (m Module) IsZero() bool {
	return m == 0
}

const (
	ModuleSpringA Module = iota + 1
	ModuleSpringB
	ModuleSpringC
	ModuleSummerVacation
	ModuleFallA
	ModuleFallB
	ModuleWinterVacation
	ModuleFallC
	ModuleSpringVacation
)

var AllModules = []Module{
	ModuleSpringA,
	ModuleSpringB,
	ModuleSpringC,
	ModuleSummerVacation,
	ModuleFallA,
	ModuleFallB,
	ModuleWinterVacation,
	ModuleFallC,
	ModuleSpringVacation,
}

func ParseModule(s string) (Module, error) {
	ret, ok := base.FindByString(AllModules, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("invalid module %s", s)
}

type ModuleDetail struct {
	ID     idtype.ModuleDetailID
	Year   shareddomain.AcademicYear
	Module Module
	Start  civil.Date
	End    civil.Date
}

func (md *ModuleDetail) Clone() *ModuleDetail {
	ret := lo.ToPtr(*md)
	return ret
}

func ConstructModuleDetail(fn func(md *ModuleDetail) (err error)) (*ModuleDetail, error) {
	md := new(ModuleDetail)
	if err := fn(md); err != nil {
		return nil, err
	}

	if md.ID.IsZero() ||
		md.Year.IsZero() ||
		md.Module.IsZero() ||
		md.Start.IsZero() ||
		md.End.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", md)
	}

	return md, nil
}

func ParseModuleDetail(id, year int, module, start, end string) (moduleDetail *ModuleDetail, err error) {
	return ConstructModuleDetail(func(md *ModuleDetail) (err error) {
		md.ID, err = idtype.ParseModuleDetailID(id)
		if err != nil {
			return
		}

		md.Year, err = shareddomain.ParseAcademicYear(year)
		if err != nil {
			return
		}

		md.Module, err = ParseModule(module)
		if err != nil {
			return
		}

		md.Start, err = civil.ParseDate(start)
		if err != nil {
			return
		}

		md.End, err = civil.ParseDate(end)
		if err != nil {
			return
		}

		return
	})
}
