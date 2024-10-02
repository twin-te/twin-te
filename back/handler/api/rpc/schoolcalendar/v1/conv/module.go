package schoolcalendarv1conv

import (
	"fmt"

	schoolcalendarv1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/schoolcalendar/v1"
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
)

func FromPBModule(pbModule schoolcalendarv1.Module) (schoolcalendardomain.Module, error) {
	switch pbModule {
	case schoolcalendarv1.Module_MODULE_SPRING_A:
		return schoolcalendardomain.ModuleSpringA, nil
	case schoolcalendarv1.Module_MODULE_SPRING_B:
		return schoolcalendardomain.ModuleSpringB, nil
	case schoolcalendarv1.Module_MODULE_SPRING_C:
		return schoolcalendardomain.ModuleSpringC, nil
	case schoolcalendarv1.Module_MODULE_SUMMER_VACATION:
		return schoolcalendardomain.ModuleSummerVacation, nil
	case schoolcalendarv1.Module_MODULE_FALL_A:
		return schoolcalendardomain.ModuleFallA, nil
	case schoolcalendarv1.Module_MODULE_FALL_B:
		return schoolcalendardomain.ModuleFallB, nil
	case schoolcalendarv1.Module_MODULE_WINTER_VACATION:
		return schoolcalendardomain.ModuleWinterVacation, nil
	case schoolcalendarv1.Module_MODULE_FALL_C:
		return schoolcalendardomain.ModuleFallC, nil
	case schoolcalendarv1.Module_MODULE_SPRING_VACATION:
		return schoolcalendardomain.ModuleSpringVacation, nil
	}
	return 0, fmt.Errorf("invalid %v", pbModule)
}

func ToPBModule(module schoolcalendardomain.Module) (schoolcalendarv1.Module, error) {
	switch module {
	case schoolcalendardomain.ModuleSpringA:
		return schoolcalendarv1.Module_MODULE_SPRING_A, nil
	case schoolcalendardomain.ModuleSpringB:
		return schoolcalendarv1.Module_MODULE_SPRING_B, nil
	case schoolcalendardomain.ModuleSpringC:
		return schoolcalendarv1.Module_MODULE_SPRING_C, nil
	case schoolcalendardomain.ModuleSummerVacation:
		return schoolcalendarv1.Module_MODULE_SUMMER_VACATION, nil
	case schoolcalendardomain.ModuleFallA:
		return schoolcalendarv1.Module_MODULE_FALL_A, nil
	case schoolcalendardomain.ModuleFallB:
		return schoolcalendarv1.Module_MODULE_FALL_B, nil
	case schoolcalendardomain.ModuleWinterVacation:
		return schoolcalendarv1.Module_MODULE_WINTER_VACATION, nil
	case schoolcalendardomain.ModuleFallC:
		return schoolcalendarv1.Module_MODULE_FALL_C, nil
	case schoolcalendardomain.ModuleSpringVacation:
		return schoolcalendarv1.Module_MODULE_SPRING_VACATION, nil
	}
	return 0, fmt.Errorf("invalid %v", module)
}
