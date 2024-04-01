package timetablev1conv

import (
	"fmt"

	timetablev1 "github.com/twin-te/twinte-back/handler/api/rpcgen/timetable/v1"
	timetabledomain "github.com/twin-te/twinte-back/module/timetable/domain"
)

func FromPBModule(pbModule timetablev1.Module) (timetabledomain.Module, error) {
	switch pbModule {
	case timetablev1.Module_MODULE_SPRING_A:
		return timetabledomain.ModuleSpringA, nil
	case timetablev1.Module_MODULE_SPRING_B:
		return timetabledomain.ModuleSpringB, nil
	case timetablev1.Module_MODULE_SPRING_C:
		return timetabledomain.ModuleSpringC, nil
	case timetablev1.Module_MODULE_FALL_A:
		return timetabledomain.ModuleFallA, nil
	case timetablev1.Module_MODULE_FALL_B:
		return timetabledomain.ModuleFallB, nil
	case timetablev1.Module_MODULE_FALL_C:
		return timetabledomain.ModuleFallC, nil
	case timetablev1.Module_MODULE_SUMMER_VACATION:
		return timetabledomain.ModuleSummerVacation, nil
	case timetablev1.Module_MODULE_SPRING_VACATION:
		return timetabledomain.ModuleSpringVacation, nil
	}
	return 0, fmt.Errorf("invalid %#v", pbModule)
}

func ToPBModule(module timetabledomain.Module) (timetablev1.Module, error) {
	switch module {
	case timetabledomain.ModuleSpringA:
		return timetablev1.Module_MODULE_SPRING_A, nil
	case timetabledomain.ModuleSpringB:
		return timetablev1.Module_MODULE_SPRING_B, nil
	case timetabledomain.ModuleSpringC:
		return timetablev1.Module_MODULE_SPRING_C, nil
	case timetabledomain.ModuleFallA:
		return timetablev1.Module_MODULE_FALL_A, nil
	case timetabledomain.ModuleFallB:
		return timetablev1.Module_MODULE_FALL_B, nil
	case timetabledomain.ModuleFallC:
		return timetablev1.Module_MODULE_FALL_C, nil
	case timetabledomain.ModuleSummerVacation:
		return timetablev1.Module_MODULE_SUMMER_VACATION, nil
	case timetabledomain.ModuleSpringVacation:
		return timetablev1.Module_MODULE_SPRING_VACATION, nil
	}
	return 0, fmt.Errorf("invalid %#v", module)
}
