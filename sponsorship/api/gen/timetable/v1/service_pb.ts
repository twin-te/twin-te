// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file timetable/v1/service.proto (package timetable.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AcademicYear, UUID, UUIDList } from "../../shared/type_pb";
import { Course, CourseMethod, CourseMethodList, RegisteredCourse, Schedule, ScheduleList, Tag } from "./type_pb";

/**
 * @generated from message timetable.v1.GetCoursesByCodesRequest
 */
export class GetCoursesByCodesRequest extends Message<GetCoursesByCodesRequest> {
  /**
   * @generated from field: shared.AcademicYear year = 1;
   */
  year?: AcademicYear;

  /**
   * @generated from field: repeated string codes = 2;
   */
  codes: string[] = [];

  constructor(data?: PartialMessage<GetCoursesByCodesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.GetCoursesByCodesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "message", T: AcademicYear },
    { no: 2, name: "codes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCoursesByCodesRequest {
    return new GetCoursesByCodesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCoursesByCodesRequest {
    return new GetCoursesByCodesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCoursesByCodesRequest {
    return new GetCoursesByCodesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetCoursesByCodesRequest | PlainMessage<GetCoursesByCodesRequest> | undefined, b: GetCoursesByCodesRequest | PlainMessage<GetCoursesByCodesRequest> | undefined): boolean {
    return proto3.util.equals(GetCoursesByCodesRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.GetCoursesByCodesResponse
 */
export class GetCoursesByCodesResponse extends Message<GetCoursesByCodesResponse> {
  /**
   * @generated from field: repeated timetable.v1.Course courses = 1;
   */
  courses: Course[] = [];

  constructor(data?: PartialMessage<GetCoursesByCodesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.GetCoursesByCodesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courses", kind: "message", T: Course, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCoursesByCodesResponse {
    return new GetCoursesByCodesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCoursesByCodesResponse {
    return new GetCoursesByCodesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCoursesByCodesResponse {
    return new GetCoursesByCodesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetCoursesByCodesResponse | PlainMessage<GetCoursesByCodesResponse> | undefined, b: GetCoursesByCodesResponse | PlainMessage<GetCoursesByCodesResponse> | undefined): boolean {
    return proto3.util.equals(GetCoursesByCodesResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.SearchCoursesRequest
 */
export class SearchCoursesRequest extends Message<SearchCoursesRequest> {
  /**
   * @generated from field: shared.AcademicYear year = 1;
   */
  year?: AcademicYear;

  /**
   * @generated from field: repeated string keywords = 2;
   */
  keywords: string[] = [];

  /**
   * @generated from field: repeated string code_prefixes_included = 3;
   */
  codePrefixesIncluded: string[] = [];

  /**
   * @generated from field: repeated string code_prefixes_excluded = 4;
   */
  codePrefixesExcluded: string[] = [];

  /**
   * @generated from field: repeated timetable.v1.Schedule schedules_fully_included = 5;
   */
  schedulesFullyIncluded: Schedule[] = [];

  /**
   * @generated from field: repeated timetable.v1.Schedule schedules_partially_overlapped = 6;
   */
  schedulesPartiallyOverlapped: Schedule[] = [];

  /**
   * @generated from field: int32 limit = 7;
   */
  limit = 0;

  /**
   * @generated from field: int32 offset = 8;
   */
  offset = 0;

  constructor(data?: PartialMessage<SearchCoursesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.SearchCoursesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "message", T: AcademicYear },
    { no: 2, name: "keywords", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "code_prefixes_included", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "code_prefixes_excluded", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "schedules_fully_included", kind: "message", T: Schedule, repeated: true },
    { no: 6, name: "schedules_partially_overlapped", kind: "message", T: Schedule, repeated: true },
    { no: 7, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "offset", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchCoursesRequest {
    return new SearchCoursesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchCoursesRequest {
    return new SearchCoursesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchCoursesRequest {
    return new SearchCoursesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SearchCoursesRequest | PlainMessage<SearchCoursesRequest> | undefined, b: SearchCoursesRequest | PlainMessage<SearchCoursesRequest> | undefined): boolean {
    return proto3.util.equals(SearchCoursesRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.SearchCoursesResponse
 */
export class SearchCoursesResponse extends Message<SearchCoursesResponse> {
  /**
   * @generated from field: repeated timetable.v1.Course courses = 1;
   */
  courses: Course[] = [];

  constructor(data?: PartialMessage<SearchCoursesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.SearchCoursesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courses", kind: "message", T: Course, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchCoursesResponse {
    return new SearchCoursesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchCoursesResponse {
    return new SearchCoursesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchCoursesResponse {
    return new SearchCoursesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SearchCoursesResponse | PlainMessage<SearchCoursesResponse> | undefined, b: SearchCoursesResponse | PlainMessage<SearchCoursesResponse> | undefined): boolean {
    return proto3.util.equals(SearchCoursesResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.CreateRegisteredCoursesByCodesRequest
 */
export class CreateRegisteredCoursesByCodesRequest extends Message<CreateRegisteredCoursesByCodesRequest> {
  /**
   * @generated from field: shared.AcademicYear year = 1;
   */
  year?: AcademicYear;

  /**
   * @generated from field: repeated string codes = 2;
   */
  codes: string[] = [];

  constructor(data?: PartialMessage<CreateRegisteredCoursesByCodesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CreateRegisteredCoursesByCodesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "message", T: AcademicYear },
    { no: 2, name: "codes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRegisteredCoursesByCodesRequest {
    return new CreateRegisteredCoursesByCodesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRegisteredCoursesByCodesRequest {
    return new CreateRegisteredCoursesByCodesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRegisteredCoursesByCodesRequest {
    return new CreateRegisteredCoursesByCodesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRegisteredCoursesByCodesRequest | PlainMessage<CreateRegisteredCoursesByCodesRequest> | undefined, b: CreateRegisteredCoursesByCodesRequest | PlainMessage<CreateRegisteredCoursesByCodesRequest> | undefined): boolean {
    return proto3.util.equals(CreateRegisteredCoursesByCodesRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.CreateRegisteredCoursesByCodesResponse
 */
export class CreateRegisteredCoursesByCodesResponse extends Message<CreateRegisteredCoursesByCodesResponse> {
  /**
   * @generated from field: repeated timetable.v1.RegisteredCourse registered_courses = 1;
   */
  registeredCourses: RegisteredCourse[] = [];

  constructor(data?: PartialMessage<CreateRegisteredCoursesByCodesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CreateRegisteredCoursesByCodesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "registered_courses", kind: "message", T: RegisteredCourse, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRegisteredCoursesByCodesResponse {
    return new CreateRegisteredCoursesByCodesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRegisteredCoursesByCodesResponse {
    return new CreateRegisteredCoursesByCodesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRegisteredCoursesByCodesResponse {
    return new CreateRegisteredCoursesByCodesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRegisteredCoursesByCodesResponse | PlainMessage<CreateRegisteredCoursesByCodesResponse> | undefined, b: CreateRegisteredCoursesByCodesResponse | PlainMessage<CreateRegisteredCoursesByCodesResponse> | undefined): boolean {
    return proto3.util.equals(CreateRegisteredCoursesByCodesResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.CreateRegisteredCourseManuallyRequest
 */
export class CreateRegisteredCourseManuallyRequest extends Message<CreateRegisteredCourseManuallyRequest> {
  /**
   * @generated from field: shared.AcademicYear year = 1;
   */
  year?: AcademicYear;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string instructors = 3;
   */
  instructors = "";

  /**
   * @generated from field: string credit = 4;
   */
  credit = "";

  /**
   * @generated from field: repeated timetable.v1.CourseMethod methods = 5;
   */
  methods: CourseMethod[] = [];

  /**
   * @generated from field: repeated timetable.v1.Schedule schedules = 6;
   */
  schedules: Schedule[] = [];

  constructor(data?: PartialMessage<CreateRegisteredCourseManuallyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CreateRegisteredCourseManuallyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "message", T: AcademicYear },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "instructors", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "credit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "methods", kind: "enum", T: proto3.getEnumType(CourseMethod), repeated: true },
    { no: 6, name: "schedules", kind: "message", T: Schedule, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRegisteredCourseManuallyRequest {
    return new CreateRegisteredCourseManuallyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRegisteredCourseManuallyRequest {
    return new CreateRegisteredCourseManuallyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRegisteredCourseManuallyRequest {
    return new CreateRegisteredCourseManuallyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRegisteredCourseManuallyRequest | PlainMessage<CreateRegisteredCourseManuallyRequest> | undefined, b: CreateRegisteredCourseManuallyRequest | PlainMessage<CreateRegisteredCourseManuallyRequest> | undefined): boolean {
    return proto3.util.equals(CreateRegisteredCourseManuallyRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.CreateRegisteredCourseManuallyResponse
 */
export class CreateRegisteredCourseManuallyResponse extends Message<CreateRegisteredCourseManuallyResponse> {
  /**
   * @generated from field: timetable.v1.RegisteredCourse registered_course = 1;
   */
  registeredCourse?: RegisteredCourse;

  constructor(data?: PartialMessage<CreateRegisteredCourseManuallyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CreateRegisteredCourseManuallyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "registered_course", kind: "message", T: RegisteredCourse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRegisteredCourseManuallyResponse {
    return new CreateRegisteredCourseManuallyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRegisteredCourseManuallyResponse {
    return new CreateRegisteredCourseManuallyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRegisteredCourseManuallyResponse {
    return new CreateRegisteredCourseManuallyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRegisteredCourseManuallyResponse | PlainMessage<CreateRegisteredCourseManuallyResponse> | undefined, b: CreateRegisteredCourseManuallyResponse | PlainMessage<CreateRegisteredCourseManuallyResponse> | undefined): boolean {
    return proto3.util.equals(CreateRegisteredCourseManuallyResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.GetRegisteredCoursesRequest
 */
export class GetRegisteredCoursesRequest extends Message<GetRegisteredCoursesRequest> {
  /**
   * @generated from field: optional shared.AcademicYear year = 1;
   */
  year?: AcademicYear;

  constructor(data?: PartialMessage<GetRegisteredCoursesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.GetRegisteredCoursesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "message", T: AcademicYear, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRegisteredCoursesRequest {
    return new GetRegisteredCoursesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRegisteredCoursesRequest {
    return new GetRegisteredCoursesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRegisteredCoursesRequest {
    return new GetRegisteredCoursesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRegisteredCoursesRequest | PlainMessage<GetRegisteredCoursesRequest> | undefined, b: GetRegisteredCoursesRequest | PlainMessage<GetRegisteredCoursesRequest> | undefined): boolean {
    return proto3.util.equals(GetRegisteredCoursesRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.GetRegisteredCoursesResponse
 */
export class GetRegisteredCoursesResponse extends Message<GetRegisteredCoursesResponse> {
  /**
   * @generated from field: repeated timetable.v1.RegisteredCourse registered_courses = 1;
   */
  registeredCourses: RegisteredCourse[] = [];

  constructor(data?: PartialMessage<GetRegisteredCoursesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.GetRegisteredCoursesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "registered_courses", kind: "message", T: RegisteredCourse, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRegisteredCoursesResponse {
    return new GetRegisteredCoursesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRegisteredCoursesResponse {
    return new GetRegisteredCoursesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRegisteredCoursesResponse {
    return new GetRegisteredCoursesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetRegisteredCoursesResponse | PlainMessage<GetRegisteredCoursesResponse> | undefined, b: GetRegisteredCoursesResponse | PlainMessage<GetRegisteredCoursesResponse> | undefined): boolean {
    return proto3.util.equals(GetRegisteredCoursesResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.UpdateRegisteredCourseRequest
 */
export class UpdateRegisteredCourseRequest extends Message<UpdateRegisteredCourseRequest> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  /**
   * @generated from field: optional string name = 2;
   */
  name?: string;

  /**
   * @generated from field: optional string instructors = 3;
   */
  instructors?: string;

  /**
   * @generated from field: optional string credit = 4;
   */
  credit?: string;

  /**
   * @generated from field: optional timetable.v1.CourseMethodList methods = 5;
   */
  methods?: CourseMethodList;

  /**
   * @generated from field: optional timetable.v1.ScheduleList schedules = 6;
   */
  schedules?: ScheduleList;

  /**
   * @generated from field: optional string memo = 7;
   */
  memo?: string;

  /**
   * @generated from field: optional int32 attendance = 8;
   */
  attendance?: number;

  /**
   * @generated from field: optional int32 absence = 9;
   */
  absence?: number;

  /**
   * @generated from field: optional int32 late = 10;
   */
  late?: number;

  /**
   * @generated from field: optional shared.UUIDList tag_ids = 11;
   */
  tagIds?: UUIDList;

  constructor(data?: PartialMessage<UpdateRegisteredCourseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.UpdateRegisteredCourseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "instructors", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "credit", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "methods", kind: "message", T: CourseMethodList, opt: true },
    { no: 6, name: "schedules", kind: "message", T: ScheduleList, opt: true },
    { no: 7, name: "memo", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 8, name: "attendance", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 9, name: "absence", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 10, name: "late", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 11, name: "tag_ids", kind: "message", T: UUIDList, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRegisteredCourseRequest {
    return new UpdateRegisteredCourseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRegisteredCourseRequest {
    return new UpdateRegisteredCourseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRegisteredCourseRequest {
    return new UpdateRegisteredCourseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateRegisteredCourseRequest | PlainMessage<UpdateRegisteredCourseRequest> | undefined, b: UpdateRegisteredCourseRequest | PlainMessage<UpdateRegisteredCourseRequest> | undefined): boolean {
    return proto3.util.equals(UpdateRegisteredCourseRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.UpdateRegisteredCourseResponse
 */
export class UpdateRegisteredCourseResponse extends Message<UpdateRegisteredCourseResponse> {
  /**
   * @generated from field: timetable.v1.RegisteredCourse registered_course = 1;
   */
  registeredCourse?: RegisteredCourse;

  constructor(data?: PartialMessage<UpdateRegisteredCourseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.UpdateRegisteredCourseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "registered_course", kind: "message", T: RegisteredCourse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRegisteredCourseResponse {
    return new UpdateRegisteredCourseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRegisteredCourseResponse {
    return new UpdateRegisteredCourseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRegisteredCourseResponse {
    return new UpdateRegisteredCourseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateRegisteredCourseResponse | PlainMessage<UpdateRegisteredCourseResponse> | undefined, b: UpdateRegisteredCourseResponse | PlainMessage<UpdateRegisteredCourseResponse> | undefined): boolean {
    return proto3.util.equals(UpdateRegisteredCourseResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.DeleteRegisteredCourseRequest
 */
export class DeleteRegisteredCourseRequest extends Message<DeleteRegisteredCourseRequest> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  constructor(data?: PartialMessage<DeleteRegisteredCourseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.DeleteRegisteredCourseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRegisteredCourseRequest {
    return new DeleteRegisteredCourseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRegisteredCourseRequest {
    return new DeleteRegisteredCourseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRegisteredCourseRequest {
    return new DeleteRegisteredCourseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRegisteredCourseRequest | PlainMessage<DeleteRegisteredCourseRequest> | undefined, b: DeleteRegisteredCourseRequest | PlainMessage<DeleteRegisteredCourseRequest> | undefined): boolean {
    return proto3.util.equals(DeleteRegisteredCourseRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.DeleteRegisteredCourseResponse
 */
export class DeleteRegisteredCourseResponse extends Message<DeleteRegisteredCourseResponse> {
  constructor(data?: PartialMessage<DeleteRegisteredCourseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.DeleteRegisteredCourseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRegisteredCourseResponse {
    return new DeleteRegisteredCourseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRegisteredCourseResponse {
    return new DeleteRegisteredCourseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRegisteredCourseResponse {
    return new DeleteRegisteredCourseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRegisteredCourseResponse | PlainMessage<DeleteRegisteredCourseResponse> | undefined, b: DeleteRegisteredCourseResponse | PlainMessage<DeleteRegisteredCourseResponse> | undefined): boolean {
    return proto3.util.equals(DeleteRegisteredCourseResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.CreateTagRequest
 */
export class CreateTagRequest extends Message<CreateTagRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<CreateTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CreateTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTagRequest {
    return new CreateTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTagRequest {
    return new CreateTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTagRequest {
    return new CreateTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTagRequest | PlainMessage<CreateTagRequest> | undefined, b: CreateTagRequest | PlainMessage<CreateTagRequest> | undefined): boolean {
    return proto3.util.equals(CreateTagRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.CreateTagResponse
 */
export class CreateTagResponse extends Message<CreateTagResponse> {
  /**
   * @generated from field: timetable.v1.Tag tag = 1;
   */
  tag?: Tag;

  constructor(data?: PartialMessage<CreateTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.CreateTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tag", kind: "message", T: Tag },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTagResponse {
    return new CreateTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTagResponse {
    return new CreateTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTagResponse {
    return new CreateTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTagResponse | PlainMessage<CreateTagResponse> | undefined, b: CreateTagResponse | PlainMessage<CreateTagResponse> | undefined): boolean {
    return proto3.util.equals(CreateTagResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.GetTagsRequest
 */
export class GetTagsRequest extends Message<GetTagsRequest> {
  constructor(data?: PartialMessage<GetTagsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.GetTagsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTagsRequest {
    return new GetTagsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTagsRequest {
    return new GetTagsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTagsRequest {
    return new GetTagsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTagsRequest | PlainMessage<GetTagsRequest> | undefined, b: GetTagsRequest | PlainMessage<GetTagsRequest> | undefined): boolean {
    return proto3.util.equals(GetTagsRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.GetTagsResponse
 */
export class GetTagsResponse extends Message<GetTagsResponse> {
  /**
   * @generated from field: repeated timetable.v1.Tag tags = 1;
   */
  tags: Tag[] = [];

  constructor(data?: PartialMessage<GetTagsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.GetTagsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tags", kind: "message", T: Tag, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTagsResponse {
    return new GetTagsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTagsResponse {
    return new GetTagsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTagsResponse {
    return new GetTagsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTagsResponse | PlainMessage<GetTagsResponse> | undefined, b: GetTagsResponse | PlainMessage<GetTagsResponse> | undefined): boolean {
    return proto3.util.equals(GetTagsResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.UpdateTagRequest
 */
export class UpdateTagRequest extends Message<UpdateTagRequest> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  /**
   * @generated from field: optional string name = 2;
   */
  name?: string;

  constructor(data?: PartialMessage<UpdateTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.UpdateTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTagRequest {
    return new UpdateTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTagRequest {
    return new UpdateTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTagRequest {
    return new UpdateTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTagRequest | PlainMessage<UpdateTagRequest> | undefined, b: UpdateTagRequest | PlainMessage<UpdateTagRequest> | undefined): boolean {
    return proto3.util.equals(UpdateTagRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.UpdateTagResponse
 */
export class UpdateTagResponse extends Message<UpdateTagResponse> {
  /**
   * @generated from field: timetable.v1.Tag tag = 1;
   */
  tag?: Tag;

  constructor(data?: PartialMessage<UpdateTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.UpdateTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tag", kind: "message", T: Tag },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTagResponse {
    return new UpdateTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTagResponse {
    return new UpdateTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTagResponse {
    return new UpdateTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTagResponse | PlainMessage<UpdateTagResponse> | undefined, b: UpdateTagResponse | PlainMessage<UpdateTagResponse> | undefined): boolean {
    return proto3.util.equals(UpdateTagResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.DeleteTagRequest
 */
export class DeleteTagRequest extends Message<DeleteTagRequest> {
  /**
   * @generated from field: shared.UUID id = 1;
   */
  id?: UUID;

  constructor(data?: PartialMessage<DeleteTagRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.DeleteTagRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: UUID },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTagRequest {
    return new DeleteTagRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTagRequest {
    return new DeleteTagRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTagRequest {
    return new DeleteTagRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTagRequest | PlainMessage<DeleteTagRequest> | undefined, b: DeleteTagRequest | PlainMessage<DeleteTagRequest> | undefined): boolean {
    return proto3.util.equals(DeleteTagRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.DeleteTagResponse
 */
export class DeleteTagResponse extends Message<DeleteTagResponse> {
  constructor(data?: PartialMessage<DeleteTagResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.DeleteTagResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTagResponse {
    return new DeleteTagResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTagResponse {
    return new DeleteTagResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTagResponse {
    return new DeleteTagResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTagResponse | PlainMessage<DeleteTagResponse> | undefined, b: DeleteTagResponse | PlainMessage<DeleteTagResponse> | undefined): boolean {
    return proto3.util.equals(DeleteTagResponse, a, b);
  }
}

/**
 * @generated from message timetable.v1.RearrangeTagsRequest
 */
export class RearrangeTagsRequest extends Message<RearrangeTagsRequest> {
  /**
   * Please specify all tag ids that the user have.
   *
   * @generated from field: repeated shared.UUID ids = 1;
   */
  ids: UUID[] = [];

  constructor(data?: PartialMessage<RearrangeTagsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.RearrangeTagsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ids", kind: "message", T: UUID, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RearrangeTagsRequest {
    return new RearrangeTagsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RearrangeTagsRequest {
    return new RearrangeTagsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RearrangeTagsRequest {
    return new RearrangeTagsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RearrangeTagsRequest | PlainMessage<RearrangeTagsRequest> | undefined, b: RearrangeTagsRequest | PlainMessage<RearrangeTagsRequest> | undefined): boolean {
    return proto3.util.equals(RearrangeTagsRequest, a, b);
  }
}

/**
 * @generated from message timetable.v1.RearrangeTagsResponse
 */
export class RearrangeTagsResponse extends Message<RearrangeTagsResponse> {
  /**
   * @generated from field: repeated timetable.v1.Tag tags = 1;
   */
  tags: Tag[] = [];

  constructor(data?: PartialMessage<RearrangeTagsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "timetable.v1.RearrangeTagsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tags", kind: "message", T: Tag, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RearrangeTagsResponse {
    return new RearrangeTagsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RearrangeTagsResponse {
    return new RearrangeTagsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RearrangeTagsResponse {
    return new RearrangeTagsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RearrangeTagsResponse | PlainMessage<RearrangeTagsResponse> | undefined, b: RearrangeTagsResponse | PlainMessage<RearrangeTagsResponse> | undefined): boolean {
    return proto3.util.equals(RearrangeTagsResponse, a, b);
  }
}

