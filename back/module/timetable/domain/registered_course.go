package timetabledomain

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/twin-te/twinte-back/base"
	shareddomain "github.com/twin-te/twinte-back/module/shared/domain"
	"github.com/twin-te/twinte-back/module/shared/domain/idtype"
	sharedhelper "github.com/twin-te/twinte-back/module/shared/helper"
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
	CourseID    *idtype.CourseID
	Name        *shareddomain.RequiredString
	Instructors *string
	Credit      *Credit
	Methods     *[]CourseMethod
	Schedules   *[]Schedule
	Memo        string
	Attendance  shareddomain.NonNegativeInt
	Absence     shareddomain.NonNegativeInt
	Late        shareddomain.NonNegativeInt
	TagIDs      []idtype.TagID

	EntityBeforeUpdated *RegisteredCourse

	CourseAssociation shareddomain.Association[*Course]
}

func (rc *RegisteredCourse) HasBasedCourse() bool {
	return rc.CourseID != nil
}

func (rc *RegisteredCourse) GetName() shareddomain.RequiredString {
	if rc.HasBasedCourse() {
		return lo.FromPtrOr(rc.Name, rc.CourseAssociation.MustGet().Name)
	}
	return *rc.Name
}

func (rc *RegisteredCourse) GetInstructors() string {
	if rc.HasBasedCourse() {
		return lo.FromPtrOr(rc.Instructors, rc.CourseAssociation.MustGet().Instructors)
	}
	return *rc.Instructors
}

func (rc *RegisteredCourse) GetCredit() Credit {
	if rc.HasBasedCourse() {
		return lo.FromPtrOr(rc.Credit, rc.CourseAssociation.MustGet().Credit)
	}
	return *rc.Credit
}

func (rc *RegisteredCourse) GetMethods() []CourseMethod {
	if rc.HasBasedCourse() {
		return lo.FromPtrOr(rc.Methods, rc.CourseAssociation.MustGet().Methods)
	}
	return *rc.Methods
}

func (rc *RegisteredCourse) GetSchedules() []Schedule {
	if rc.HasBasedCourse() {
		return lo.FromPtrOr(rc.Schedules, rc.CourseAssociation.MustGet().Schedules)
	}
	return *rc.Schedules
}

func (rc *RegisteredCourse) Clone() *RegisteredCourse {
	ret := lo.ToPtr(*rc)

	if rc.CourseID != nil {
		*ret.CourseID = *rc.CourseID
	}

	if rc.Name != nil {
		*ret.Name = *rc.Name
	}

	if rc.Instructors != nil {
		*ret.Instructors = *rc.Instructors
	}

	if rc.Credit != nil {
		*ret.Credit = *rc.Credit
	}

	if rc.Methods != nil {
		*ret.Methods = base.CopySlice(*rc.Methods)
	}

	if rc.Schedules != nil {
		*ret.Schedules = base.CopySlice(*rc.Schedules)
	}

	ret.TagIDs = base.CopySlice(rc.TagIDs)

	return ret
}

func (rc *RegisteredCourse) BeforeUpdateHook() {
	rc.EntityBeforeUpdated = rc.Clone()
}

type RegisteredCourseDataToUpdate struct {
	Name        *shareddomain.RequiredString
	Instructors *string
	Credit      *Credit
	Methods     *[]CourseMethod
	Schedules   *[]Schedule
	Memo        *string
	Attendance  *shareddomain.NonNegativeInt
	Absence     *shareddomain.NonNegativeInt
	Late        *shareddomain.NonNegativeInt
	TagIDs      *[]idtype.TagID
}

func (rc *RegisteredCourse) updateName(name shareddomain.RequiredString) {
	if rc.HasBasedCourse() && rc.Name == nil && rc.CourseAssociation.MustGet().Name == name {
		return
	}
	rc.Name = &name
}

func (rc *RegisteredCourse) updateInstructors(instructors string) {
	if rc.HasBasedCourse() && rc.Instructors == nil && rc.CourseAssociation.MustGet().Instructors == instructors {
		return
	}
	rc.Instructors = &instructors
}

func (rc *RegisteredCourse) updateCredit(credit Credit) {
	if rc.HasBasedCourse() && rc.Credit == nil && rc.CourseAssociation.MustGet().Credit == credit {
		return
	}
	rc.Credit = &credit
}

func (rc *RegisteredCourse) updateMethods(methods []CourseMethod) {
	if rc.HasBasedCourse() && rc.Methods == nil && base.HaveSameElements(rc.CourseAssociation.MustGet().Methods, methods) {
		return
	}
	rc.Methods = &methods
}

func (rc *RegisteredCourse) updateSchedules(schedules []Schedule) {
	if rc.HasBasedCourse() && rc.Schedules == nil && base.HaveSameElements(rc.CourseAssociation.MustGet().Schedules, schedules) {
		return
	}
	rc.Schedules = &schedules
}

func (rc *RegisteredCourse) Update(data RegisteredCourseDataToUpdate) error {
	if data.Name != nil {
		rc.updateName(*data.Name)
	}

	if data.Instructors != nil {
		rc.updateInstructors(*data.Instructors)
	}

	if data.Credit != nil {
		rc.updateCredit(*data.Credit)
	}

	if data.Methods != nil {
		if err := sharedhelper.ValidateDuplicates(*data.Methods); err != nil {
			return err
		}
		rc.updateMethods(*data.Methods)
	}

	if data.Schedules != nil {
		if err := sharedhelper.ValidateDuplicates(*data.Schedules); err != nil {
			return err
		}
		rc.updateSchedules(*data.Schedules)
	}

	if data.Memo != nil {
		rc.Memo = *data.Memo
	}

	if data.Attendance != nil {
		rc.Attendance = *data.Attendance
	}

	if data.Absence != nil {
		rc.Absence = *data.Absence
	}

	if data.Late != nil {
		rc.Late = *data.Late
	}

	if data.TagIDs != nil {
		if err := sharedhelper.ValidateDuplicates(*data.TagIDs); err != nil {
			return err
		}
		rc.TagIDs = *data.TagIDs
	}

	return nil
}

func ConstructRegisteredCourse(fn func(rc *RegisteredCourse) (err error)) (*RegisteredCourse, error) {
	rc := new(RegisteredCourse)
	if err := fn(rc); err != nil {
		return nil, err
	}

	if rc.CourseID == nil && (rc.Name == nil || rc.Instructors == nil || rc.Credit == nil || rc.Methods == nil || rc.Schedules == nil) {
		return nil, fmt.Errorf("the registered course, which does not have the based course, must have name, instructors, credit, methods, and schedules. %+v", rc)
	}

	if rc.ID.IsZero() || rc.UserID.IsZero() || rc.Year.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", rc)
	}

	return rc, nil
}
