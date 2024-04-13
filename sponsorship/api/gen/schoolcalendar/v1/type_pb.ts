// @generated by protoc-gen-es v1.8.0 with parameter "target=ts,import_extension=none"
// @generated from file schoolcalendar/v1/type.proto (package schoolcalendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { RFC3339FullDate, Weekday } from "../../shared/type_pb";

/**
 * @generated from enum schoolcalendar.v1.EventType
 */
export enum EventType {
  /**
   * @generated from enum value: EVENT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: EVENT_TYPE_HOLIDAY = 1;
   */
  HOLIDAY = 1,

  /**
   * @generated from enum value: EVENT_TYPE_PUBLIC_HOLIDAY = 2;
   */
  PUBLIC_HOLIDAY = 2,

  /**
   * @generated from enum value: EVENT_TYPE_EXAM = 3;
   */
  EXAM = 3,

  /**
   * @generated from enum value: EVENT_TYPE_SUBSTITUTE_DAY = 4;
   */
  SUBSTITUTE_DAY = 4,

  /**
   * @generated from enum value: EVENT_TYPE_OTHER = 5;
   */
  OTHER = 5,
}
// Retrieve enum metadata with: proto3.getEnumType(EventType)
proto3.util.setEnumType(EventType, "schoolcalendar.v1.EventType", [
  { no: 0, name: "EVENT_TYPE_UNSPECIFIED" },
  { no: 1, name: "EVENT_TYPE_HOLIDAY" },
  { no: 2, name: "EVENT_TYPE_PUBLIC_HOLIDAY" },
  { no: 3, name: "EVENT_TYPE_EXAM" },
  { no: 4, name: "EVENT_TYPE_SUBSTITUTE_DAY" },
  { no: 5, name: "EVENT_TYPE_OTHER" },
]);

/**
 * @generated from enum schoolcalendar.v1.Module
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
   * @generated from enum value: MODULE_SUMMER_VACATION = 4;
   */
  SUMMER_VACATION = 4,

  /**
   * @generated from enum value: MODULE_FALL_A = 5;
   */
  FALL_A = 5,

  /**
   * @generated from enum value: MODULE_FALL_B = 6;
   */
  FALL_B = 6,

  /**
   * @generated from enum value: MODULE_WINTER_VACATION = 7;
   */
  WINTER_VACATION = 7,

  /**
   * @generated from enum value: MODULE_FALL_C = 8;
   */
  FALL_C = 8,

  /**
   * @generated from enum value: MODULE_SPRING_VACATION = 9;
   */
  SPRING_VACATION = 9,
}
// Retrieve enum metadata with: proto3.getEnumType(Module)
proto3.util.setEnumType(Module, "schoolcalendar.v1.Module", [
  { no: 0, name: "MODULE_UNSPECIFIED" },
  { no: 1, name: "MODULE_SPRING_A" },
  { no: 2, name: "MODULE_SPRING_B" },
  { no: 3, name: "MODULE_SPRING_C" },
  { no: 4, name: "MODULE_SUMMER_VACATION" },
  { no: 5, name: "MODULE_FALL_A" },
  { no: 6, name: "MODULE_FALL_B" },
  { no: 7, name: "MODULE_WINTER_VACATION" },
  { no: 8, name: "MODULE_FALL_C" },
  { no: 9, name: "MODULE_SPRING_VACATION" },
]);

/**
 * @generated from message schoolcalendar.v1.Event
 */
export class Event extends Message<Event> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: schoolcalendar.v1.EventType type = 2;
   */
  type = EventType.UNSPECIFIED;

  /**
   * @generated from field: shared.RFC3339FullDate date = 3;
   */
  date?: RFC3339FullDate;

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * @generated from field: optional shared.Weekday change_to = 5;
   */
  changeTo?: Weekday;

  constructor(data?: PartialMessage<Event>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "schoolcalendar.v1.Event";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 3, name: "date", kind: "message", T: RFC3339FullDate },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "change_to", kind: "enum", T: proto3.getEnumType(Weekday), opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event {
    return new Event().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event {
    return new Event().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event {
    return new Event().fromJsonString(jsonString, options);
  }

  static equals(a: Event | PlainMessage<Event> | undefined, b: Event | PlainMessage<Event> | undefined): boolean {
    return proto3.util.equals(Event, a, b);
  }
}
