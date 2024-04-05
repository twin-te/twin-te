// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: timetable/v1/type.proto

package timetablev1

import (
	sharedpb "github.com/twin-te/twin-te/back/handler/api/rpcgen/sharedpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Module int32

const (
	Module_MODULE_UNSPECIFIED     Module = 0
	Module_MODULE_SPRING_A        Module = 1
	Module_MODULE_SPRING_B        Module = 2
	Module_MODULE_SPRING_C        Module = 3
	Module_MODULE_FALL_A          Module = 4
	Module_MODULE_FALL_B          Module = 5
	Module_MODULE_FALL_C          Module = 6
	Module_MODULE_SUMMER_VACATION Module = 7
	Module_MODULE_SPRING_VACATION Module = 8
)

// Enum value maps for Module.
var (
	Module_name = map[int32]string{
		0: "MODULE_UNSPECIFIED",
		1: "MODULE_SPRING_A",
		2: "MODULE_SPRING_B",
		3: "MODULE_SPRING_C",
		4: "MODULE_FALL_A",
		5: "MODULE_FALL_B",
		6: "MODULE_FALL_C",
		7: "MODULE_SUMMER_VACATION",
		8: "MODULE_SPRING_VACATION",
	}
	Module_value = map[string]int32{
		"MODULE_UNSPECIFIED":     0,
		"MODULE_SPRING_A":        1,
		"MODULE_SPRING_B":        2,
		"MODULE_SPRING_C":        3,
		"MODULE_FALL_A":          4,
		"MODULE_FALL_B":          5,
		"MODULE_FALL_C":          6,
		"MODULE_SUMMER_VACATION": 7,
		"MODULE_SPRING_VACATION": 8,
	}
)

func (x Module) Enum() *Module {
	p := new(Module)
	*p = x
	return p
}

func (x Module) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Module) Descriptor() protoreflect.EnumDescriptor {
	return file_timetable_v1_type_proto_enumTypes[0].Descriptor()
}

func (Module) Type() protoreflect.EnumType {
	return &file_timetable_v1_type_proto_enumTypes[0]
}

func (x Module) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Module.Descriptor instead.
func (Module) EnumDescriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{0}
}

type Day int32

const (
	Day_DAY_UNSPECIFIED Day = 0
	Day_DAY_SUN         Day = 1
	Day_DAY_MON         Day = 2
	Day_DAY_TUE         Day = 3
	Day_DAY_WED         Day = 4
	Day_DAY_THU         Day = 5
	Day_DAY_FRI         Day = 6
	Day_DAY_SAT         Day = 7
	Day_DAY_INTENSIVE   Day = 8
	Day_DAY_APPOINTMENT Day = 9
	Day_DAY_ANY_TIME    Day = 10
)

// Enum value maps for Day.
var (
	Day_name = map[int32]string{
		0:  "DAY_UNSPECIFIED",
		1:  "DAY_SUN",
		2:  "DAY_MON",
		3:  "DAY_TUE",
		4:  "DAY_WED",
		5:  "DAY_THU",
		6:  "DAY_FRI",
		7:  "DAY_SAT",
		8:  "DAY_INTENSIVE",
		9:  "DAY_APPOINTMENT",
		10: "DAY_ANY_TIME",
	}
	Day_value = map[string]int32{
		"DAY_UNSPECIFIED": 0,
		"DAY_SUN":         1,
		"DAY_MON":         2,
		"DAY_TUE":         3,
		"DAY_WED":         4,
		"DAY_THU":         5,
		"DAY_FRI":         6,
		"DAY_SAT":         7,
		"DAY_INTENSIVE":   8,
		"DAY_APPOINTMENT": 9,
		"DAY_ANY_TIME":    10,
	}
)

func (x Day) Enum() *Day {
	p := new(Day)
	*p = x
	return p
}

func (x Day) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Day) Descriptor() protoreflect.EnumDescriptor {
	return file_timetable_v1_type_proto_enumTypes[1].Descriptor()
}

func (Day) Type() protoreflect.EnumType {
	return &file_timetable_v1_type_proto_enumTypes[1]
}

func (x Day) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Day.Descriptor instead.
func (Day) EnumDescriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{1}
}

type CourseMethod int32

const (
	CourseMethod_COURSE_METHOD_UNSPECIFIED         CourseMethod = 0
	CourseMethod_COURSE_METHOD_ONLINE_ASYNCHRONOUS CourseMethod = 1
	CourseMethod_COURSE_METHOD_ONLINE_SYNCHRONOUS  CourseMethod = 2
	CourseMethod_COURSE_METHOD_FACE_TO_FACE        CourseMethod = 3
	CourseMethod_COURSE_METHOD_OTHERS              CourseMethod = 4
)

// Enum value maps for CourseMethod.
var (
	CourseMethod_name = map[int32]string{
		0: "COURSE_METHOD_UNSPECIFIED",
		1: "COURSE_METHOD_ONLINE_ASYNCHRONOUS",
		2: "COURSE_METHOD_ONLINE_SYNCHRONOUS",
		3: "COURSE_METHOD_FACE_TO_FACE",
		4: "COURSE_METHOD_OTHERS",
	}
	CourseMethod_value = map[string]int32{
		"COURSE_METHOD_UNSPECIFIED":         0,
		"COURSE_METHOD_ONLINE_ASYNCHRONOUS": 1,
		"COURSE_METHOD_ONLINE_SYNCHRONOUS":  2,
		"COURSE_METHOD_FACE_TO_FACE":        3,
		"COURSE_METHOD_OTHERS":              4,
	}
)

func (x CourseMethod) Enum() *CourseMethod {
	p := new(CourseMethod)
	*p = x
	return p
}

func (x CourseMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourseMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_timetable_v1_type_proto_enumTypes[2].Descriptor()
}

func (CourseMethod) Type() protoreflect.EnumType {
	return &file_timetable_v1_type_proto_enumTypes[2]
}

func (x CourseMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourseMethod.Descriptor instead.
func (CourseMethod) EnumDescriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{2}
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module Module `protobuf:"varint,1,opt,name=module,proto3,enum=timetable.v1.Module" json:"module,omitempty"`
	Day    Day    `protobuf:"varint,2,opt,name=day,proto3,enum=timetable.v1.Day" json:"day,omitempty"`
	Period int32  `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	Rooms  string `protobuf:"bytes,4,opt,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timetable_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_timetable_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *Schedule) GetModule() Module {
	if x != nil {
		return x.Module
	}
	return Module_MODULE_UNSPECIFIED
}

func (x *Schedule) GetDay() Day {
	if x != nil {
		return x.Day
	}
	return Day_DAY_UNSPECIFIED
}

func (x *Schedule) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *Schedule) GetRooms() string {
	if x != nil {
		return x.Rooms
	}
	return ""
}

type CourseMethodList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []CourseMethod `protobuf:"varint,1,rep,packed,name=values,proto3,enum=timetable.v1.CourseMethod" json:"values,omitempty"`
}

func (x *CourseMethodList) Reset() {
	*x = CourseMethodList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timetable_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseMethodList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseMethodList) ProtoMessage() {}

func (x *CourseMethodList) ProtoReflect() protoreflect.Message {
	mi := &file_timetable_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseMethodList.ProtoReflect.Descriptor instead.
func (*CourseMethodList) Descriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *CourseMethodList) GetValues() []CourseMethod {
	if x != nil {
		return x.Values
	}
	return nil
}

type ScheduleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Schedule `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ScheduleList) Reset() {
	*x = ScheduleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timetable_v1_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleList) ProtoMessage() {}

func (x *ScheduleList) ProtoReflect() protoreflect.Message {
	mi := &file_timetable_v1_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleList.ProtoReflect.Descriptor instead.
func (*ScheduleList) Descriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleList) GetValues() []*Schedule {
	if x != nil {
		return x.Values
	}
	return nil
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                *sharedpb.UUID            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Year              *sharedpb.AcademicYear    `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Code              string                    `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name              string                    `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Instructors       string                    `protobuf:"bytes,5,opt,name=instructors,proto3" json:"instructors,omitempty"`
	Credit            string                    `protobuf:"bytes,6,opt,name=credit,proto3" json:"credit,omitempty"`
	Overview          string                    `protobuf:"bytes,7,opt,name=overview,proto3" json:"overview,omitempty"`
	Remarks           string                    `protobuf:"bytes,8,opt,name=remarks,proto3" json:"remarks,omitempty"`
	LastUpdatedAt     *sharedpb.RFC3339DateTime `protobuf:"bytes,9,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
	RecommendedGrades []int32                   `protobuf:"varint,10,rep,packed,name=recommended_grades,json=recommendedGrades,proto3" json:"recommended_grades,omitempty"`
	Methods           []CourseMethod            `protobuf:"varint,11,rep,packed,name=methods,proto3,enum=timetable.v1.CourseMethod" json:"methods,omitempty"`
	Schedules         []*Schedule               `protobuf:"bytes,12,rep,name=schedules,proto3" json:"schedules,omitempty"`
	HasParseError     bool                      `protobuf:"varint,13,opt,name=has_parse_error,json=hasParseError,proto3" json:"has_parse_error,omitempty"`
	IsAnnual          bool                      `protobuf:"varint,14,opt,name=is_annual,json=isAnnual,proto3" json:"is_annual,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timetable_v1_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_timetable_v1_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{3}
}

func (x *Course) GetId() *sharedpb.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Course) GetYear() *sharedpb.AcademicYear {
	if x != nil {
		return x.Year
	}
	return nil
}

func (x *Course) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetInstructors() string {
	if x != nil {
		return x.Instructors
	}
	return ""
}

func (x *Course) GetCredit() string {
	if x != nil {
		return x.Credit
	}
	return ""
}

func (x *Course) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *Course) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *Course) GetLastUpdatedAt() *sharedpb.RFC3339DateTime {
	if x != nil {
		return x.LastUpdatedAt
	}
	return nil
}

func (x *Course) GetRecommendedGrades() []int32 {
	if x != nil {
		return x.RecommendedGrades
	}
	return nil
}

func (x *Course) GetMethods() []CourseMethod {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Course) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *Course) GetHasParseError() bool {
	if x != nil {
		return x.HasParseError
	}
	return false
}

func (x *Course) GetIsAnnual() bool {
	if x != nil {
		return x.IsAnnual
	}
	return false
}

// If it has the based course, code is present.
type RegisteredCourse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *sharedpb.UUID         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      *sharedpb.UUID         `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Year        *sharedpb.AcademicYear `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Code        *string                `protobuf:"bytes,4,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Name        string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Instructors string                 `protobuf:"bytes,6,opt,name=instructors,proto3" json:"instructors,omitempty"`
	Credit      string                 `protobuf:"bytes,7,opt,name=credit,proto3" json:"credit,omitempty"`
	Methods     []CourseMethod         `protobuf:"varint,8,rep,packed,name=methods,proto3,enum=timetable.v1.CourseMethod" json:"methods,omitempty"`
	Schedules   []*Schedule            `protobuf:"bytes,9,rep,name=schedules,proto3" json:"schedules,omitempty"`
	Memo        string                 `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo,omitempty"`
	Attendance  int32                  `protobuf:"varint,11,opt,name=attendance,proto3" json:"attendance,omitempty"`
	Absence     int32                  `protobuf:"varint,12,opt,name=absence,proto3" json:"absence,omitempty"`
	Late        int32                  `protobuf:"varint,13,opt,name=late,proto3" json:"late,omitempty"`
	TagIds      []*sharedpb.UUID       `protobuf:"bytes,14,rep,name=tag_ids,json=tagIds,proto3" json:"tag_ids,omitempty"`
}

func (x *RegisteredCourse) Reset() {
	*x = RegisteredCourse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timetable_v1_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredCourse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredCourse) ProtoMessage() {}

func (x *RegisteredCourse) ProtoReflect() protoreflect.Message {
	mi := &file_timetable_v1_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredCourse.ProtoReflect.Descriptor instead.
func (*RegisteredCourse) Descriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{4}
}

func (x *RegisteredCourse) GetId() *sharedpb.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RegisteredCourse) GetUserId() *sharedpb.UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *RegisteredCourse) GetYear() *sharedpb.AcademicYear {
	if x != nil {
		return x.Year
	}
	return nil
}

func (x *RegisteredCourse) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *RegisteredCourse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisteredCourse) GetInstructors() string {
	if x != nil {
		return x.Instructors
	}
	return ""
}

func (x *RegisteredCourse) GetCredit() string {
	if x != nil {
		return x.Credit
	}
	return ""
}

func (x *RegisteredCourse) GetMethods() []CourseMethod {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *RegisteredCourse) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *RegisteredCourse) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *RegisteredCourse) GetAttendance() int32 {
	if x != nil {
		return x.Attendance
	}
	return 0
}

func (x *RegisteredCourse) GetAbsence() int32 {
	if x != nil {
		return x.Absence
	}
	return 0
}

func (x *RegisteredCourse) GetLate() int32 {
	if x != nil {
		return x.Late
	}
	return 0
}

func (x *RegisteredCourse) GetTagIds() []*sharedpb.UUID {
	if x != nil {
		return x.TagIds
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *sharedpb.UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   *sharedpb.UUID `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name     string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Position int32          `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timetable_v1_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_timetable_v1_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_timetable_v1_type_proto_rawDescGZIP(), []int{5}
}

func (x *Tag) GetId() *sharedpb.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Tag) GetUserId() *sharedpb.UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

var File_timetable_v1_type_proto protoreflect.FileDescriptor

var file_timetable_v1_type_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x08, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x79, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x3e, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x89, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x46,
	0x43, 0x33, 0x33, 0x33, 0x39, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0d, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x12,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x5f, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x68, 0x61, 0x73, 0x50, 0x61, 0x72, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x22, 0xe6, 0x03, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x41, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x34, 0x0a,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x62, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x67, 0x49, 0x64, 0x73, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x7a, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0xd0, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53,
	0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x10, 0x02, 0x12, 0x13,
	0x0a, 0x0f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f,
	0x43, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x46, 0x41,
	0x4c, 0x4c, 0x5f, 0x41, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45,
	0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x5f, 0x42, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16,
	0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x52, 0x5f, 0x56, 0x41,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x44, 0x55,
	0x4c, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x41, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x08, 0x2a, 0xaf, 0x01, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x12, 0x13, 0x0a, 0x0f,
	0x44, 0x41, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x59, 0x5f, 0x53, 0x55, 0x4e, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x41, 0x59, 0x5f, 0x4d, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x41, 0x59, 0x5f, 0x54, 0x55, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x59, 0x5f,
	0x57, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x59, 0x5f, 0x54, 0x48, 0x55,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x59, 0x5f, 0x46, 0x52, 0x49, 0x10, 0x06, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x41, 0x59, 0x5f, 0x53, 0x41, 0x54, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d,
	0x44, 0x41, 0x59, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x56, 0x45, 0x10, 0x08, 0x12,
	0x13, 0x0a, 0x0f, 0x44, 0x41, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x41, 0x59, 0x5f, 0x41, 0x4e, 0x59, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x10, 0x0a, 0x2a, 0xb4, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x55, 0x52, 0x53,
	0x45, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45,
	0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x41,
	0x53, 0x59, 0x4e, 0x43, 0x48, 0x52, 0x4f, 0x4e, 0x4f, 0x55, 0x53, 0x10, 0x01, 0x12, 0x24, 0x0a,
	0x20, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4f,
	0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x48, 0x52, 0x4f, 0x4e, 0x4f, 0x55,
	0x53, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x41, 0x43,
	0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x53, 0x10, 0x04, 0x42, 0x4d, 0x5a,
	0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x69, 0x6e,
	0x2d, 0x74, 0x65, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63,
	0x67, 0x65, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timetable_v1_type_proto_rawDescOnce sync.Once
	file_timetable_v1_type_proto_rawDescData = file_timetable_v1_type_proto_rawDesc
)

func file_timetable_v1_type_proto_rawDescGZIP() []byte {
	file_timetable_v1_type_proto_rawDescOnce.Do(func() {
		file_timetable_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_timetable_v1_type_proto_rawDescData)
	})
	return file_timetable_v1_type_proto_rawDescData
}

var file_timetable_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_timetable_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_timetable_v1_type_proto_goTypes = []interface{}{
	(Module)(0),                      // 0: timetable.v1.Module
	(Day)(0),                         // 1: timetable.v1.Day
	(CourseMethod)(0),                // 2: timetable.v1.CourseMethod
	(*Schedule)(nil),                 // 3: timetable.v1.Schedule
	(*CourseMethodList)(nil),         // 4: timetable.v1.CourseMethodList
	(*ScheduleList)(nil),             // 5: timetable.v1.ScheduleList
	(*Course)(nil),                   // 6: timetable.v1.Course
	(*RegisteredCourse)(nil),         // 7: timetable.v1.RegisteredCourse
	(*Tag)(nil),                      // 8: timetable.v1.Tag
	(*sharedpb.UUID)(nil),            // 9: shared.UUID
	(*sharedpb.AcademicYear)(nil),    // 10: shared.AcademicYear
	(*sharedpb.RFC3339DateTime)(nil), // 11: shared.RFC3339DateTime
}
var file_timetable_v1_type_proto_depIdxs = []int32{
	0,  // 0: timetable.v1.Schedule.module:type_name -> timetable.v1.Module
	1,  // 1: timetable.v1.Schedule.day:type_name -> timetable.v1.Day
	2,  // 2: timetable.v1.CourseMethodList.values:type_name -> timetable.v1.CourseMethod
	3,  // 3: timetable.v1.ScheduleList.values:type_name -> timetable.v1.Schedule
	9,  // 4: timetable.v1.Course.id:type_name -> shared.UUID
	10, // 5: timetable.v1.Course.year:type_name -> shared.AcademicYear
	11, // 6: timetable.v1.Course.last_updated_at:type_name -> shared.RFC3339DateTime
	2,  // 7: timetable.v1.Course.methods:type_name -> timetable.v1.CourseMethod
	3,  // 8: timetable.v1.Course.schedules:type_name -> timetable.v1.Schedule
	9,  // 9: timetable.v1.RegisteredCourse.id:type_name -> shared.UUID
	9,  // 10: timetable.v1.RegisteredCourse.user_id:type_name -> shared.UUID
	10, // 11: timetable.v1.RegisteredCourse.year:type_name -> shared.AcademicYear
	2,  // 12: timetable.v1.RegisteredCourse.methods:type_name -> timetable.v1.CourseMethod
	3,  // 13: timetable.v1.RegisteredCourse.schedules:type_name -> timetable.v1.Schedule
	9,  // 14: timetable.v1.RegisteredCourse.tag_ids:type_name -> shared.UUID
	9,  // 15: timetable.v1.Tag.id:type_name -> shared.UUID
	9,  // 16: timetable.v1.Tag.user_id:type_name -> shared.UUID
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_timetable_v1_type_proto_init() }
func file_timetable_v1_type_proto_init() {
	if File_timetable_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timetable_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_timetable_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseMethodList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_timetable_v1_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_timetable_v1_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_timetable_v1_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredCourse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_timetable_v1_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_timetable_v1_type_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_timetable_v1_type_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timetable_v1_type_proto_goTypes,
		DependencyIndexes: file_timetable_v1_type_proto_depIdxs,
		EnumInfos:         file_timetable_v1_type_proto_enumTypes,
		MessageInfos:      file_timetable_v1_type_proto_msgTypes,
	}.Build()
	File_timetable_v1_type_proto = out.File
	file_timetable_v1_type_proto_rawDesc = nil
	file_timetable_v1_type_proto_goTypes = nil
	file_timetable_v1_type_proto_depIdxs = nil
}
