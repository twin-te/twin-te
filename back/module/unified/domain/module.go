package unifieddomain

import (
	schoolcalendardomain "github.com/twin-te/twin-te/back/module/schoolcalendar/domain"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

var TimetableModuleToSchoolCalendarModule = map[timetabledomain.Module]schoolcalendardomain.Module{
	timetabledomain.ModuleSpringA:        schoolcalendardomain.ModuleSpringA,
	timetabledomain.ModuleSpringB:        schoolcalendardomain.ModuleSpringB,
	timetabledomain.ModuleSpringC:        schoolcalendardomain.ModuleSpringC,
	timetabledomain.ModuleSummerVacation: schoolcalendardomain.ModuleSummerVacation,
	timetabledomain.ModuleFallA:          schoolcalendardomain.ModuleFallA,
	timetabledomain.ModuleFallB:          schoolcalendardomain.ModuleFallB,
	timetabledomain.ModuleFallC:          schoolcalendardomain.ModuleFallC,
	timetabledomain.ModuleSpringVacation: schoolcalendardomain.ModuleSpringVacation,
}
