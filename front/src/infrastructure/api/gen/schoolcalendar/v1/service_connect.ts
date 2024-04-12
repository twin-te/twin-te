// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file schoolcalendar/v1/service.proto (package schoolcalendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetEventsByDateRequest, GetEventsByDateResponse, GetModuleByDateRequest, GetModuleByDateResponse } from "./service_pb.js";
import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * The following error codes are not stated explicitly in the each rpc, but may be returned.
 *   - shared.InvalidArgument
 *   - shared.Unauthenticated
 *   - shared.Unauthorized
 *
 * @generated from service schoolcalendar.v1.SchoolCalendarService
 */
export const SchoolCalendarService = {
  typeName: "schoolcalendar.v1.SchoolCalendarService",
  methods: {
    /**
     * @generated from rpc schoolcalendar.v1.SchoolCalendarService.GetEventsByDate
     */
    getEventsByDate: {
      name: "GetEventsByDate",
      I: GetEventsByDateRequest,
      O: GetEventsByDateResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc schoolcalendar.v1.SchoolCalendarService.GetModuleByDate
     */
    getModuleByDate: {
      name: "GetModuleByDate",
      I: GetModuleByDateRequest,
      O: GetModuleByDateResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
  }
} as const;

