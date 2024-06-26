// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=none"
// @generated from file announcement/v1/service.proto (package announcement.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetAnnouncementsRequest, GetAnnouncementsResponse, ReadAnnouncementsRequest, ReadAnnouncementsResponse } from "./service_pb";
import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * The following error codes are not stated explicitly in the each rpc, but may be returned.
 *   - shared.InvalidArgument
 *   - shared.Unauthenticated
 *   - shared.Unauthorized
 *
 * @generated from service announcement.v1.AnnouncementService
 */
export const AnnouncementService = {
  typeName: "announcement.v1.AnnouncementService",
  methods: {
    /**
     * @generated from rpc announcement.v1.AnnouncementService.GetAnnouncements
     */
    getAnnouncements: {
      name: "GetAnnouncements",
      I: GetAnnouncementsRequest,
      O: GetAnnouncementsResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc announcement.v1.AnnouncementService.ReadAnnouncements
     */
    readAnnouncements: {
      name: "ReadAnnouncements",
      I: ReadAnnouncementsRequest,
      O: ReadAnnouncementsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

