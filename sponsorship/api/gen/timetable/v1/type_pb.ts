// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file timetable/v1/type.proto (package timetable.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AcademicYear, RFC3339DateTime, UUID } from "../../shared/type_pb";

/**
 * @generated from enum timetable.v1.Module
 */
export enum Module {
  /**
   * @generated from enum value: MODULE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: MODULE_SPRING_A = 1;
   */
  SPRING_A = 1,

  /**
   * @generated from enum value: MODULE_SPRING_B = 2;
   */
  SPRING_B = 2,

  /**
   * @generated from enum value: MODULE_SPRING_C = 3;
   */
  SPRING_C = 3,

  /**
   * @generated from enum value: MODULE_FALL_A = 4;
   */
  FALL_A = 4,

  /**
   * @generated from enum value: MODULE_FALL_B = 5;
   */
  FALL_B = 5,

  /**
   * @generated from enum value: MODULE_FALL_C = 6;
   */
  FALL_C = 6,

  /**
   * @generated from enum value: MODULE_SUMMER_VACATION = 7;
   */
  SUMMER_VACATION = 7,

  /**
   * @generated from enum value: MODULE_SPRING_VACATION = 8;
   */
  SPRING_VACATION = 8,
}
// Retrieve enum metadata with: proto3.getEnumType(Module)
proto3.util.setEnumType(Module, "timetable.v1.Module", [
  { no: 0, name: "MODULE_UNSPECIFIED" },
  { no: 1, name: "MODULE_SPRING_A" },
  { no: 2, name: "MODULE_SPRING_B" },
  { no: 3, name: "MODULE_SPRING_C" },
  { no: 4, name: "MODULE_FALL_A" },
  { no: 5, name: "MODULE_FALL_B" },
  { no: 6, name: "MODULE_FALL_C" },
  { no: 7, name: "MODULE_SUMMER_VACATION" },
  { no: 8, name: "MODULE_SPRING_VACATION" },
]);

/**
 * @generated from enum timetable.v1.Day
 */
export enum Day {
  /**
   * @generated from enum value: DAY_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: DAY_SUN = 1;
   */
  SUN = 1,

  /**
   * @generated from enum value: DAY_MON = 2;
   */
  MON = 2,

  /**
   * @generated from enum value: DAY_TUE = 3;
   */
  TUE = 3,

  /**
   * @generated from enum value: DAY_WED = 4;
   */
  WED = 4,

  /**
   * @generated from enum value: DAY_THU = 5;
   */
  THU = 5,

  /**
   * @generated from enum value: DAY_FRI = 6;
   */
  FRI = 6,

  /**
   * @generated from enum value: DAY_SAT = 7;
   */
  SAT = 7,

  /**
   * @generated from enum value: DAY_INTENSIVE = 8;
   */
  INTENSIVE = 8,

  /**
   * @generated from enum value: DAY_APPOINTMENT = 9;
   */
  APPOINTMENT = 9,

  /**
   * @generated from enum value: DAY_ANY_TIME = 10;
   */
  ANY_TIME = 10,

  /**
   * @generated from enum value: DAY_NT = 11;
   */
  NT = 11,
}
// Retrieve enum metadata with: proto3.getEnumType(Day)
proto3.util.setEnumType(Day, "timetable.v1.Day", [
  { no: 0, name: "DAY_UNSPECIFIED" },
  { no: 1, name: "DAY_SUN" },
  { no: 2, name: "DAY_MON" },
  { no: 3, name: "DAY_TUE" },
  { no: 4, name: "DAY_WED" },
  { no: 5, name: "DAY_THU" },
  { no: 6, name: "DAY_FRI" },
  { no: 7, name: "DAY_SAT" },
  { no: 8, name: "DAY_INTENSIVE" },
  { no: 9, name: "DAY_APPOINTMENT" },
  { no: 10, name: "DAY_ANY_TIME" },
  { no: 11, name: "DAY_NT" },
]);

/**
 * @generated from enum timetable.v1.CourseMethod
 */
export enum CourseMethod {
  /**
   * @generated from enum value: COURSE_METHOD_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: COURSE_METHOD_ONLINE_ASYNCHRONOUS = 1;
   */
  ONLINE_ASYNCHRONOUS = 1,

  /**
   * @generated from enum value: COURSE_METHOD_ONLINE_SYNCHRONOUS = 2;
   */
  ONLINE_SYNCHRONOUS = 2,

  /**
   * @generated from enum value: COURSE_METHOD_FACE_TO_FACE = 3;
   */
  FACE_TO_FACE = 3,

  /**
   * @generated from enum value: COURSE_METHOD_OTHERS = 4;
   */
  OTHERS = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(CourseMethod)
proto3.util.setEnumType(CourseMethod, "timetable.v1.CourseMethod", [
  { no: 0, name: "COURSE_METHOD_UNSPECIFIED" },
  { no: 1, name: "COURSE_METHOD_ONLINE_ASYNCHRONOUS" },
  { no: 2, name: "COURSE_METHOD_ONLINE_SYNCHRONOUS" },
  { no: 3, name: "COURSE_METHOD_FACE_TO_FACE" },
  { no: 4, name: "COURSE_METHOD_OTHERS" },
]);

/**
 * @generated from message timetable.v1.Schedule
 */
export class Schedule extends Message<Schedule> {
  /**
   * @generated from field: timetable.v1.Module module = 1;
   */
  module = Module.UNSPECIFIED;

  /**
   * @generated from field: timetable.v1.Day day = 2;
   */
  day = Day.UNSPECIFIED;

  /**
   * @generated from field: int32 period = 3;
   */
  period = 0;

  /**
   * @generated from field: string locations = 4;
   */
  locations = "";

  constructor(data?: PartialMessage<Schedule>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.Schedule";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "module", kind: "enum", T: proto3.getEnumType(Module) },
    { no: 2, name: "day", kind: "enum", T: proto3.getEnumType(Day) },
    { no: 3, name: "period", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "locations", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Schedule {
    return new Schedule().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Schedule {
    return new Schedule().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Schedule {
    return new Schedule().fromJsonString(jsonString, options);
  }

  static equals(a: Schedule | PlainMessage<Schedule> | undefined, b: Schedule | PlainMessage<Schedule> | undefined): boolean {
    return proto3.util.equals(Schedule, a, b);
  }
}

/**
 * @generated from message timetable.v1.CourseMethodList
 */
export class CourseMethodList extends Message<CourseMethodList> {
  /**
   * @generated from field: repeated timetable.v1.CourseMethod values = 1;
   */
  values: CourseMethod[] = [];

  constructor(data?: PartialMessage<CourseMethodList>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CourseMethodList";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "values", kind: "enum", T: proto3.getEnumType(CourseMethod), repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CourseMethodList {
    return new CourseMethodList().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CourseMethodList {
    return new CourseMethodList().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CourseMethodList {
    return new CourseMethodList().fromJsonString(jsonString, options);
  }

  static equals(a: CourseMethodList | PlainMessage<CourseMethodList> | undefined, b: CourseMethodList | PlainMessage<CourseMethodList> | undefined): boolean {
    return proto3.util.equals(CourseMethodList, a, b);
  }
}

/**
 * @generated from message timetable.v1.ScheduleList
 */
export class ScheduleList extends Message<ScheduleList> {
  /**
   * @generated from field: repeated timetable.v1.Schedule values = 1;
   */
  values: Schedule[] = [];

  constructor(data?: PartialMessage<ScheduleList>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.ScheduleList";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "values", kind: "message", T: Schedule, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ScheduleList {
    return new ScheduleList().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ScheduleList {
    return new ScheduleList().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ScheduleList {
    return new ScheduleList().fromJsonString(jsonString, options);
  }

  static equals(a: ScheduleList | PlainMessage<ScheduleList> | undefined, b: ScheduleList | PlainMessage<ScheduleList> | undefined): boolean {
    return proto3.util.equals(ScheduleList, a, b);
  }
}

/**
 * @generated from message timetable.v1.Course
 */
export class Course extends Message<Course> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  /**
   * @generated from field: shared.AcademicYear year = 2;
   */
  year?: AcademicYear;

  /**
   * @generated from field: string code = 3;
   */
  code = "";

  /**
   * @generated from field: string name = 4;
   */
  name = "";

  /**
   * @generated from field: string instructors = 5;
   */
  instructors = "";

  /**
   * @generated from field: string credit = 6;
   */
  credit = "";

  /**
   * @generated from field: string overview = 7;
   */
  overview = "";

  /**
   * @generated from field: string remarks = 8;
   */
  remarks = "";

  /**
   * @generated from field: shared.RFC3339DateTime last_updated_at = 9;
   */
  lastUpdatedAt?: RFC3339DateTime;

  /**
   * @generated from field: repeated int32 recommended_grades = 10;
   */
  recommendedGrades: number[] = [];

  /**
   * @generated from field: repeated timetable.v1.CourseMethod methods = 11;
   */
  methods: CourseMethod[] = [];

  /**
   * @generated from field: repeated timetable.v1.Schedule schedules = 12;
   */
  schedules: Schedule[] = [];

  /**
   * @generated from field: bool has_parse_error = 13;
   */
  hasParseError = false;

  /**
   * @generated from field: bool is_annual = 14;
   */
  isAnnual = false;

  constructor(data?: PartialMessage<Course>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.Course";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
    { no: 2, name: "year", kind: "message", T: AcademicYear },
    { no: 3, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "instructors", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "credit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "overview", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "remarks", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "last_updated_at", kind: "message", T: RFC3339DateTime },
    { no: 10, name: "recommended_grades", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 11, name: "methods", kind: "enum", T: proto3.getEnumType(CourseMethod), repeated: true },
    { no: 12, name: "schedules", kind: "message", T: Schedule, repeated: true },
    { no: 13, name: "has_parse_error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 14, name: "is_annual", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Course {
    return new Course().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Course {
    return new Course().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Course {
    return new Course().fromJsonString(jsonString, options);
  }

  static equals(a: Course | PlainMessage<Course> | undefined, b: Course | PlainMessage<Course> | undefined): boolean {
    return proto3.util.equals(Course, a, b);
  }
}

/**
 * If it has the based course, code is present.
 *
 * @generated from message timetable.v1.RegisteredCourse
 */
export class RegisteredCourse extends Message<RegisteredCourse> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  /**
   * @generated from field: shared.UUID user_id = 2;
   */
  userId?: UUID;

  /**
   * @generated from field: shared.AcademicYear year = 3;
   */
  year?: AcademicYear;

  /**
   * @generated from field: optional string code = 4;
   */
  code?: string;

  /**
   * @generated from field: string name = 5;
   */
  name = "";

  /**
   * @generated from field: string instructors = 6;
   */
  instructors = "";

  /**
   * @generated from field: string credit = 7;
   */
  credit = "";

  /**
   * @generated from field: repeated timetable.v1.CourseMethod methods = 8;
   */
  methods: CourseMethod[] = [];

  /**
   * @generated from field: repeated timetable.v1.Schedule schedules = 9;
   */
  schedules: Schedule[] = [];

  /**
   * @generated from field: string memo = 10;
   */
  memo = "";

  /**
   * @generated from field: int32 attendance = 11;
   */
  attendance = 0;

  /**
   * @generated from field: int32 absence = 12;
   */
  absence = 0;

  /**
   * @generated from field: int32 late = 13;
   */
  late = 0;

  /**
   * @generated from field: repeated shared.UUID tag_ids = 14;
   */
  tagIds: UUID[] = [];

  constructor(data?: PartialMessage<RegisteredCourse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.RegisteredCourse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
    { no: 2, name: "user_id", kind: "message", T: UUID },
    { no: 3, name: "year", kind: "message", T: AcademicYear },
    { no: 4, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "instructors", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "credit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "methods", kind: "enum", T: proto3.getEnumType(CourseMethod), repeated: true },
    { no: 9, name: "schedules", kind: "message", T: Schedule, repeated: true },
    { no: 10, name: "memo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "attendance", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 12, name: "absence", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 13, name: "late", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 14, name: "tag_ids", kind: "message", T: UUID, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisteredCourse {
    return new RegisteredCourse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisteredCourse {
    return new RegisteredCourse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisteredCourse {
    return new RegisteredCourse().fromJsonString(jsonString, options);
  }

  static equals(a: RegisteredCourse | PlainMessage<RegisteredCourse> | undefined, b: RegisteredCourse | PlainMessage<RegisteredCourse> | undefined): boolean {
    return proto3.util.equals(RegisteredCourse, a, b);
  }
}

/**
 * @generated from message timetable.v1.Tag
 */
export class Tag extends Message<Tag> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  /**
   * @generated from field: shared.UUID user_id = 2;
   */
  userId?: UUID;

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: int32 position = 4;
   */
  position = 0;

  constructor(data?: PartialMessage<Tag>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.Tag";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
    { no: 2, name: "user_id", kind: "message", T: UUID },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "position", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Tag {
    return new Tag().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Tag {
    return new Tag().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Tag {
    return new Tag().fromJsonString(jsonString, options);
  }

  static equals(a: Tag | PlainMessage<Tag> | undefined, b: Tag | PlainMessage<Tag> | undefined): boolean {
    return proto3.util.equals(Tag, a, b);
  }
}

