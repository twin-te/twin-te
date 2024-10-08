package timetabledomain

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

var (
	ParseAttendance = shareddomain.NewNonNegativeIntParser("attendance")
	ParseAbsence    = shareddomain.NewNonNegativeIntParser("absence")
	ParseLate       = shareddomain.NewNonNegativeIntParser("late")
)

// RegisteredCourse is identified by one of the following fields.
//   - ID
//   - UserID and CourseID ( if it has based course )
//
// There are two types of RegisteredCourse.
//   - RegisteredCourse created manually
//   - RegisteredCourse that has the based course
//
// If RegisteredCourse has the based course, the following fields are always present.
//   - CourseID
//
// And the following fields are present only if overwritten.
//   - Name
//   - Instructors
//   - Credit
//   - Methods
//   - Schedules
//
// If RegisteredCourse is created manually, the following fields are always present.
//   - Name
//   - Instructors
//   - Credit
//   - Methods
//   - Schedules
type RegisteredCourse struct {
	ID          idtype.RegisteredCourseID
	UserID      idtype.UserID
	Year        shareddomain.AcademicYear
	CourseID    mo.Option[idtype.CourseID]
	Name        mo.Option[shareddomain.RequiredString]
	Instructors mo.Option[string]
	Credit      mo.Option[Credit]
	Methods     mo.Option[[]CourseMethod]
	Schedules   mo.Option[[]Schedule]
	Memo        string
	Attendance  shareddomain.NonNegativeInt
	Absence     shareddomain.NonNegativeInt
	Late        shareddomain.NonNegativeInt
	TagIDs      []idtype.TagID

	BeforeUpdated mo.Option[*RegisteredCourse]

	CourseAssociation shareddomain.Association[*Course]
}

func (rc *RegisteredCourse) HasBasedCourse() bool {
	return rc.CourseID.IsPresent()
}

func (rc *RegisteredCourse) GetName() shareddomain.RequiredString {
	if rc.HasBasedCourse() {
		return rc.Name.OrElse(rc.CourseAssociation.MustGet().Name)
	}
	return rc.Name.MustGet()
}

func (rc *RegisteredCourse) GetInstructors() string {
	if rc.HasBasedCourse() {
		return rc.Instructors.OrElse(rc.CourseAssociation.MustGet().Instructors)
	}
	return rc.Instructors.MustGet()
}

func (rc *RegisteredCourse) GetCredit() Credit {
	if rc.HasBasedCourse() {
		return rc.Credit.OrElse(rc.CourseAssociation.MustGet().Credit)
	}
	return rc.Credit.MustGet()
}

func (rc *RegisteredCourse) GetMethods() []CourseMethod {
	if rc.HasBasedCourse() {
		return rc.Methods.OrElse(rc.CourseAssociation.MustGet().Methods)
	}
	return rc.Methods.MustGet()
}

func (rc *RegisteredCourse) GetSchedules() []Schedule {
	if rc.HasBasedCourse() {
		return rc.Schedules.OrElse(rc.CourseAssociation.MustGet().Schedules)
	}
	return rc.Schedules.MustGet()
}

func (rc *RegisteredCourse) Clone() *RegisteredCourse {
	ret := lo.ToPtr(*rc)
	ret.Methods = base.OptionCloneBy(rc.Methods, base.CopySlice)
	ret.Schedules = base.OptionCloneBy(rc.Schedules, base.CopySlice)
	ret.TagIDs = base.CopySlice(rc.TagIDs)
	return ret
}

func (rc *RegisteredCourse) BeforeUpdateHook() {
	rc.BeforeUpdated = mo.Some(rc.Clone())
}

type RegisteredCourseDataToUpdate struct {
	Name        mo.Option[shareddomain.RequiredString]
	Instructors mo.Option[string]
	Credit      mo.Option[Credit]
	Methods     mo.Option[[]CourseMethod]
	Schedules   mo.Option[[]Schedule]
	Memo        mo.Option[string]
	Attendance  mo.Option[shareddomain.NonNegativeInt]
	Absence     mo.Option[shareddomain.NonNegativeInt]
	Late        mo.Option[shareddomain.NonNegativeInt]
	TagIDs      mo.Option[[]idtype.TagID]
}

func (rc *RegisteredCourse) updateName(name shareddomain.RequiredString) {
	if rc.HasBasedCourse() && rc.Name.IsAbsent() && rc.CourseAssociation.MustGet().Name == name {
		return
	}
	rc.Name = mo.Some(name)
}

func (rc *RegisteredCourse) updateInstructors(instructors string) {
	if rc.HasBasedCourse() && rc.Instructors.IsAbsent() && rc.CourseAssociation.MustGet().Instructors == instructors {
		return
	}
	rc.Instructors = mo.Some(instructors)
}

func (rc *RegisteredCourse) updateCredit(credit Credit) {
	if rc.HasBasedCourse() && rc.Credit.IsAbsent() && rc.CourseAssociation.MustGet().Credit == credit {
		return
	}
	rc.Credit = mo.Some(credit)
}

func (rc *RegisteredCourse) updateMethods(methods []CourseMethod) {
	if rc.HasBasedCourse() && rc.Methods.IsAbsent() && base.HaveSameElements(rc.CourseAssociation.MustGet().Methods, methods) {
		return
	}
	rc.Methods = mo.Some(methods)
}

func (rc *RegisteredCourse) updateSchedules(schedules []Schedule) {
	if rc.HasBasedCourse() && rc.Schedules.IsAbsent() && base.HaveSameElements(rc.CourseAssociation.MustGet().Schedules, schedules) {
		return
	}
	rc.Schedules = mo.Some(schedules)
}

func (rc *RegisteredCourse) Update(data RegisteredCourseDataToUpdate) error {
	if name, ok := data.Name.Get(); ok {
		rc.updateName(name)
	}

	if instructors, ok := data.Instructors.Get(); ok {
		rc.updateInstructors(instructors)
	}

	if credit, ok := data.Credit.Get(); ok {
		rc.updateCredit(credit)
	}

	if methods, ok := data.Methods.Get(); ok {
		rc.updateMethods(methods)
	}

	if schedules, ok := data.Schedules.Get(); ok {
		rc.updateSchedules(schedules)
	}

	if memo, ok := data.Memo.Get(); ok {
		rc.Memo = memo
	}

	if attendance, ok := data.Attendance.Get(); ok {
		rc.Attendance = attendance
	}

	if absence, ok := data.Absence.Get(); ok {
		rc.Absence = absence
	}

	if late, ok := data.Late.Get(); ok {
		rc.Late = late
	}

	if tagIDs, ok := data.TagIDs.Get(); ok {
		rc.TagIDs = tagIDs
	}

	return nil
}

func ConstructRegisteredCourse(fn func(rc *RegisteredCourse) (err error)) (*RegisteredCourse, error) {
	rc := new(RegisteredCourse)
	if err := fn(rc); err != nil {
		return nil, err
	}

	if rc.CourseID.IsAbsent() && (rc.Name.IsAbsent() || rc.Instructors.IsAbsent() || rc.Credit.IsAbsent() || rc.Methods.IsAbsent() || rc.Schedules.IsAbsent()) {
		return nil, fmt.Errorf("the registered course, which does not have the based course, must have name, instructors, credit, methods, and schedules. %+v", rc)
	}

	if rc.ID.IsZero() || rc.UserID.IsZero() || rc.Year.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", rc)
	}

	return rc, nil
}
