// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=none"
// @generated from file timetable/v1/service.proto (package timetable.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateRegisteredCourseManuallyRequest, CreateRegisteredCourseManuallyResponse, CreateRegisteredCoursesByCodesRequest, CreateRegisteredCoursesByCodesResponse, CreateTagRequest, CreateTagResponse, DeleteRegisteredCourseRequest, DeleteRegisteredCourseResponse, DeleteTagRequest, DeleteTagResponse, ListCoursesByCodesRequest, ListCoursesByCodesResponse, ListRegisteredCoursesRequest, ListRegisteredCoursesResponse, ListTagsRequest, ListTagsResponse, RearrangeTagsRequest, RearrangeTagsResponse, SearchCoursesRequest, SearchCoursesResponse, UpdateRegisteredCourseRequest, UpdateRegisteredCourseResponse, UpdateTagRequest, UpdateTagResponse } from "./service_pb";
import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * The following error codes are not stated explicitly in the each rpc, but may be returned.
 *   - shared.InvalidArgument
 *   - shared.Unauthenticated
 *   - shared.Unauthorized
 *
 * @generated from service timetable.v1.TimetableService
 */
export const TimetableService = {
  typeName: "timetable.v1.TimetableService",
  methods: {
    /**
     * @generated from rpc timetable.v1.TimetableService.ListCoursesByCodes
     */
    listCoursesByCodes: {
      name: "ListCoursesByCodes",
      I: ListCoursesByCodesRequest,
      O: ListCoursesByCodesResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.SearchCourses
     */
    searchCourses: {
      name: "SearchCourses",
      I: SearchCoursesRequest,
      O: SearchCoursesResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.CreateRegisteredCoursesByCodes
     */
    createRegisteredCoursesByCodes: {
      name: "CreateRegisteredCoursesByCodes",
      I: CreateRegisteredCoursesByCodesRequest,
      O: CreateRegisteredCoursesByCodesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.CreateRegisteredCourseManually
     */
    createRegisteredCourseManually: {
      name: "CreateRegisteredCourseManually",
      I: CreateRegisteredCourseManuallyRequest,
      O: CreateRegisteredCourseManuallyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.ListRegisteredCourses
     */
    listRegisteredCourses: {
      name: "ListRegisteredCourses",
      I: ListRegisteredCoursesRequest,
      O: ListRegisteredCoursesResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.UpdateRegisteredCourse
     */
    updateRegisteredCourse: {
      name: "UpdateRegisteredCourse",
      I: UpdateRegisteredCourseRequest,
      O: UpdateRegisteredCourseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.DeleteRegisteredCourse
     */
    deleteRegisteredCourse: {
      name: "DeleteRegisteredCourse",
      I: DeleteRegisteredCourseRequest,
      O: DeleteRegisteredCourseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.CreateTag
     */
    createTag: {
      name: "CreateTag",
      I: CreateTagRequest,
      O: CreateTagResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.ListTags
     */
    listTags: {
      name: "ListTags",
      I: ListTagsRequest,
      O: ListTagsResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.UpdateTag
     */
    updateTag: {
      name: "UpdateTag",
      I: UpdateTagRequest,
      O: UpdateTagResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.DeleteTag
     */
    deleteTag: {
      name: "DeleteTag",
      I: DeleteTagRequest,
      O: DeleteTagResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc timetable.v1.TimetableService.RearrangeTags
     */
    rearrangeTags: {
      name: "RearrangeTags",
      I: RearrangeTagsRequest,
      O: RearrangeTagsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

