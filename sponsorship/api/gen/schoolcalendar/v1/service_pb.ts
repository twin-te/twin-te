// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file schoolcalendar/v1/service.proto (package schoolcalendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { RFC3339FullDate } from "../../shared/type_pb";
import { Event, Module } from "./type_pb";

/**
 * @generated from message schoolcalendar.v1.ListEventsByDateRequest
 */
export class ListEventsByDateRequest extends Message<ListEventsByDateRequest> {
  /**
   * @generated from field: shared.RFC3339FullDate date = 1;
   */
  date?: RFC3339FullDate;

  constructor(data?: PartialMessage<ListEventsByDateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "schoolcalendar.v1.ListEventsByDateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "message", T: RFC3339FullDate },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventsByDateRequest {
    return new ListEventsByDateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventsByDateRequest {
    return new ListEventsByDateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventsByDateRequest {
    return new ListEventsByDateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListEventsByDateRequest | PlainMessage<ListEventsByDateRequest> | undefined, b: ListEventsByDateRequest | PlainMessage<ListEventsByDateRequest> | undefined): boolean {
    return proto3.util.equals(ListEventsByDateRequest, a, b);
  }
}

/**
 * @generated from message schoolcalendar.v1.ListEventsByDateResponse
 */
export class ListEventsByDateResponse extends Message<ListEventsByDateResponse> {
  /**
   * @generated from field: repeated schoolcalendar.v1.Event events = 1;
   */
  events: Event[] = [];

  constructor(data?: PartialMessage<ListEventsByDateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "schoolcalendar.v1.ListEventsByDateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "events", kind: "message", T: Event, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventsByDateResponse {
    return new ListEventsByDateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventsByDateResponse {
    return new ListEventsByDateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventsByDateResponse {
    return new ListEventsByDateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListEventsByDateResponse | PlainMessage<ListEventsByDateResponse> | undefined, b: ListEventsByDateResponse | PlainMessage<ListEventsByDateResponse> | undefined): boolean {
    return proto3.util.equals(ListEventsByDateResponse, a, b);
  }
}

/**
 * @generated from message schoolcalendar.v1.GetModuleByDateRequest
 */
export class GetModuleByDateRequest extends Message<GetModuleByDateRequest> {
  /**
   * @generated from field: shared.RFC3339FullDate date = 1;
   */
  date?: RFC3339FullDate;

  constructor(data?: PartialMessage<GetModuleByDateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "schoolcalendar.v1.GetModuleByDateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "message", T: RFC3339FullDate },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetModuleByDateRequest {
    return new GetModuleByDateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetModuleByDateRequest {
    return new GetModuleByDateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetModuleByDateRequest {
    return new GetModuleByDateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetModuleByDateRequest | PlainMessage<GetModuleByDateRequest> | undefined, b: GetModuleByDateRequest | PlainMessage<GetModuleByDateRequest> | undefined): boolean {
    return proto3.util.equals(GetModuleByDateRequest, a, b);
  }
}

/**
 * @generated from message schoolcalendar.v1.GetModuleByDateResponse
 */
export class GetModuleByDateResponse extends Message<GetModuleByDateResponse> {
  /**
   * @generated from field: schoolcalendar.v1.Module module = 1;
   */
  module = Module.UNSPECIFIED;

  constructor(data?: PartialMessage<GetModuleByDateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "schoolcalendar.v1.GetModuleByDateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "module", kind: "enum", T: proto3.getEnumType(Module) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetModuleByDateResponse {
    return new GetModuleByDateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetModuleByDateResponse {
    return new GetModuleByDateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetModuleByDateResponse {
    return new GetModuleByDateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetModuleByDateResponse | PlainMessage<GetModuleByDateResponse> | undefined, b: GetModuleByDateResponse | PlainMessage<GetModuleByDateResponse> | undefined): boolean {
    return proto3.util.equals(GetModuleByDateResponse, a, b);
  }
}

