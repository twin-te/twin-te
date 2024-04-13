// @generated by protoc-gen-es v1.8.0 with parameter "target=ts,import_extension=none"
// @generated from file donation/v1/service.proto (package donation.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Contributor, PaymentHistory, PaymentUser, Subscription, SubscriptionPlan } from "./type_pb";

/**
 * @generated from message donation.v1.CreateOneTimeCheckoutSessionRequest
 */
export class CreateOneTimeCheckoutSessionRequest extends Message<CreateOneTimeCheckoutSessionRequest> {
  /**
   * @generated from field: int32 amount = 1;
   */
  amount = 0;

  constructor(data?: PartialMessage<CreateOneTimeCheckoutSessionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.CreateOneTimeCheckoutSessionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "amount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateOneTimeCheckoutSessionRequest {
    return new CreateOneTimeCheckoutSessionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateOneTimeCheckoutSessionRequest {
    return new CreateOneTimeCheckoutSessionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateOneTimeCheckoutSessionRequest {
    return new CreateOneTimeCheckoutSessionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateOneTimeCheckoutSessionRequest | PlainMessage<CreateOneTimeCheckoutSessionRequest> | undefined, b: CreateOneTimeCheckoutSessionRequest | PlainMessage<CreateOneTimeCheckoutSessionRequest> | undefined): boolean {
    return proto3.util.equals(CreateOneTimeCheckoutSessionRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.CreateOneTimeCheckoutSessionResponse
 */
export class CreateOneTimeCheckoutSessionResponse extends Message<CreateOneTimeCheckoutSessionResponse> {
  /**
   * @generated from field: string checkout_session_id = 1;
   */
  checkoutSessionId = "";

  constructor(data?: PartialMessage<CreateOneTimeCheckoutSessionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.CreateOneTimeCheckoutSessionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "checkout_session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateOneTimeCheckoutSessionResponse {
    return new CreateOneTimeCheckoutSessionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateOneTimeCheckoutSessionResponse {
    return new CreateOneTimeCheckoutSessionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateOneTimeCheckoutSessionResponse {
    return new CreateOneTimeCheckoutSessionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateOneTimeCheckoutSessionResponse | PlainMessage<CreateOneTimeCheckoutSessionResponse> | undefined, b: CreateOneTimeCheckoutSessionResponse | PlainMessage<CreateOneTimeCheckoutSessionResponse> | undefined): boolean {
    return proto3.util.equals(CreateOneTimeCheckoutSessionResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.CreateSubscriptionCheckoutSessionRequest
 */
export class CreateSubscriptionCheckoutSessionRequest extends Message<CreateSubscriptionCheckoutSessionRequest> {
  /**
   * @generated from field: string plan_id = 1;
   */
  planId = "";

  constructor(data?: PartialMessage<CreateSubscriptionCheckoutSessionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.CreateSubscriptionCheckoutSessionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "plan_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSubscriptionCheckoutSessionRequest {
    return new CreateSubscriptionCheckoutSessionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSubscriptionCheckoutSessionRequest {
    return new CreateSubscriptionCheckoutSessionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSubscriptionCheckoutSessionRequest {
    return new CreateSubscriptionCheckoutSessionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSubscriptionCheckoutSessionRequest | PlainMessage<CreateSubscriptionCheckoutSessionRequest> | undefined, b: CreateSubscriptionCheckoutSessionRequest | PlainMessage<CreateSubscriptionCheckoutSessionRequest> | undefined): boolean {
    return proto3.util.equals(CreateSubscriptionCheckoutSessionRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.CreateSubscriptionCheckoutSessionResponse
 */
export class CreateSubscriptionCheckoutSessionResponse extends Message<CreateSubscriptionCheckoutSessionResponse> {
  /**
   * @generated from field: string checkout_session_id = 1;
   */
  checkoutSessionId = "";

  constructor(data?: PartialMessage<CreateSubscriptionCheckoutSessionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.CreateSubscriptionCheckoutSessionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "checkout_session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSubscriptionCheckoutSessionResponse {
    return new CreateSubscriptionCheckoutSessionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSubscriptionCheckoutSessionResponse {
    return new CreateSubscriptionCheckoutSessionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSubscriptionCheckoutSessionResponse {
    return new CreateSubscriptionCheckoutSessionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSubscriptionCheckoutSessionResponse | PlainMessage<CreateSubscriptionCheckoutSessionResponse> | undefined, b: CreateSubscriptionCheckoutSessionResponse | PlainMessage<CreateSubscriptionCheckoutSessionResponse> | undefined): boolean {
    return proto3.util.equals(CreateSubscriptionCheckoutSessionResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.GetPaymentUserRequest
 */
export class GetPaymentUserRequest extends Message<GetPaymentUserRequest> {
  constructor(data?: PartialMessage<GetPaymentUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetPaymentUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPaymentUserRequest {
    return new GetPaymentUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPaymentUserRequest {
    return new GetPaymentUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPaymentUserRequest {
    return new GetPaymentUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetPaymentUserRequest | PlainMessage<GetPaymentUserRequest> | undefined, b: GetPaymentUserRequest | PlainMessage<GetPaymentUserRequest> | undefined): boolean {
    return proto3.util.equals(GetPaymentUserRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.GetPaymentUserResponse
 */
export class GetPaymentUserResponse extends Message<GetPaymentUserResponse> {
  /**
   * @generated from field: donation.v1.PaymentUser payment_user = 1;
   */
  paymentUser?: PaymentUser;

  constructor(data?: PartialMessage<GetPaymentUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetPaymentUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "payment_user", kind: "message", T: PaymentUser },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPaymentUserResponse {
    return new GetPaymentUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPaymentUserResponse {
    return new GetPaymentUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPaymentUserResponse {
    return new GetPaymentUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetPaymentUserResponse | PlainMessage<GetPaymentUserResponse> | undefined, b: GetPaymentUserResponse | PlainMessage<GetPaymentUserResponse> | undefined): boolean {
    return proto3.util.equals(GetPaymentUserResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.UpdatePaymentUserRequest
 */
export class UpdatePaymentUserRequest extends Message<UpdatePaymentUserRequest> {
  /**
   * @generated from field: optional string display_name = 1;
   */
  displayName?: string;

  /**
   * @generated from field: optional string link = 2;
   */
  link?: string;

  constructor(data?: PartialMessage<UpdatePaymentUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.UpdatePaymentUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "link", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePaymentUserRequest {
    return new UpdatePaymentUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePaymentUserRequest {
    return new UpdatePaymentUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePaymentUserRequest {
    return new UpdatePaymentUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePaymentUserRequest | PlainMessage<UpdatePaymentUserRequest> | undefined, b: UpdatePaymentUserRequest | PlainMessage<UpdatePaymentUserRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePaymentUserRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.UpdatePaymentUserResponse
 */
export class UpdatePaymentUserResponse extends Message<UpdatePaymentUserResponse> {
  /**
   * @generated from field: donation.v1.PaymentUser payment_user = 1;
   */
  paymentUser?: PaymentUser;

  constructor(data?: PartialMessage<UpdatePaymentUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.UpdatePaymentUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "payment_user", kind: "message", T: PaymentUser },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePaymentUserResponse {
    return new UpdatePaymentUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePaymentUserResponse {
    return new UpdatePaymentUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePaymentUserResponse {
    return new UpdatePaymentUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePaymentUserResponse | PlainMessage<UpdatePaymentUserResponse> | undefined, b: UpdatePaymentUserResponse | PlainMessage<UpdatePaymentUserResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePaymentUserResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.GetPaymentHistoriesRequest
 */
export class GetPaymentHistoriesRequest extends Message<GetPaymentHistoriesRequest> {
  constructor(data?: PartialMessage<GetPaymentHistoriesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetPaymentHistoriesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPaymentHistoriesRequest {
    return new GetPaymentHistoriesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPaymentHistoriesRequest {
    return new GetPaymentHistoriesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPaymentHistoriesRequest {
    return new GetPaymentHistoriesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetPaymentHistoriesRequest | PlainMessage<GetPaymentHistoriesRequest> | undefined, b: GetPaymentHistoriesRequest | PlainMessage<GetPaymentHistoriesRequest> | undefined): boolean {
    return proto3.util.equals(GetPaymentHistoriesRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.GetPaymentHistoriesResponse
 */
export class GetPaymentHistoriesResponse extends Message<GetPaymentHistoriesResponse> {
  /**
   * @generated from field: repeated donation.v1.PaymentHistory payment_histories = 1;
   */
  paymentHistories: PaymentHistory[] = [];

  constructor(data?: PartialMessage<GetPaymentHistoriesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetPaymentHistoriesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "payment_histories", kind: "message", T: PaymentHistory, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPaymentHistoriesResponse {
    return new GetPaymentHistoriesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPaymentHistoriesResponse {
    return new GetPaymentHistoriesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPaymentHistoriesResponse {
    return new GetPaymentHistoriesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetPaymentHistoriesResponse | PlainMessage<GetPaymentHistoriesResponse> | undefined, b: GetPaymentHistoriesResponse | PlainMessage<GetPaymentHistoriesResponse> | undefined): boolean {
    return proto3.util.equals(GetPaymentHistoriesResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.GetSubscriptionPlansRequest
 */
export class GetSubscriptionPlansRequest extends Message<GetSubscriptionPlansRequest> {
  constructor(data?: PartialMessage<GetSubscriptionPlansRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetSubscriptionPlansRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubscriptionPlansRequest {
    return new GetSubscriptionPlansRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubscriptionPlansRequest {
    return new GetSubscriptionPlansRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubscriptionPlansRequest {
    return new GetSubscriptionPlansRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubscriptionPlansRequest | PlainMessage<GetSubscriptionPlansRequest> | undefined, b: GetSubscriptionPlansRequest | PlainMessage<GetSubscriptionPlansRequest> | undefined): boolean {
    return proto3.util.equals(GetSubscriptionPlansRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.GetSubscriptionPlansResponse
 */
export class GetSubscriptionPlansResponse extends Message<GetSubscriptionPlansResponse> {
  /**
   * @generated from field: repeated donation.v1.SubscriptionPlan subscription_plans = 1;
   */
  subscriptionPlans: SubscriptionPlan[] = [];

  constructor(data?: PartialMessage<GetSubscriptionPlansResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetSubscriptionPlansResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "subscription_plans", kind: "message", T: SubscriptionPlan, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSubscriptionPlansResponse {
    return new GetSubscriptionPlansResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSubscriptionPlansResponse {
    return new GetSubscriptionPlansResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSubscriptionPlansResponse {
    return new GetSubscriptionPlansResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetSubscriptionPlansResponse | PlainMessage<GetSubscriptionPlansResponse> | undefined, b: GetSubscriptionPlansResponse | PlainMessage<GetSubscriptionPlansResponse> | undefined): boolean {
    return proto3.util.equals(GetSubscriptionPlansResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.GetActiveSubscriptionRequest
 */
export class GetActiveSubscriptionRequest extends Message<GetActiveSubscriptionRequest> {
  constructor(data?: PartialMessage<GetActiveSubscriptionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetActiveSubscriptionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetActiveSubscriptionRequest {
    return new GetActiveSubscriptionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetActiveSubscriptionRequest {
    return new GetActiveSubscriptionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetActiveSubscriptionRequest {
    return new GetActiveSubscriptionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetActiveSubscriptionRequest | PlainMessage<GetActiveSubscriptionRequest> | undefined, b: GetActiveSubscriptionRequest | PlainMessage<GetActiveSubscriptionRequest> | undefined): boolean {
    return proto3.util.equals(GetActiveSubscriptionRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.GetActiveSubscriptionResponse
 */
export class GetActiveSubscriptionResponse extends Message<GetActiveSubscriptionResponse> {
  /**
   * @generated from field: donation.v1.Subscription subscription = 1;
   */
  subscription?: Subscription;

  constructor(data?: PartialMessage<GetActiveSubscriptionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetActiveSubscriptionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "subscription", kind: "message", T: Subscription },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetActiveSubscriptionResponse {
    return new GetActiveSubscriptionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetActiveSubscriptionResponse {
    return new GetActiveSubscriptionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetActiveSubscriptionResponse {
    return new GetActiveSubscriptionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetActiveSubscriptionResponse | PlainMessage<GetActiveSubscriptionResponse> | undefined, b: GetActiveSubscriptionResponse | PlainMessage<GetActiveSubscriptionResponse> | undefined): boolean {
    return proto3.util.equals(GetActiveSubscriptionResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.UnsubscribeRequest
 */
export class UnsubscribeRequest extends Message<UnsubscribeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<UnsubscribeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.UnsubscribeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnsubscribeRequest {
    return new UnsubscribeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnsubscribeRequest {
    return new UnsubscribeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnsubscribeRequest {
    return new UnsubscribeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UnsubscribeRequest | PlainMessage<UnsubscribeRequest> | undefined, b: UnsubscribeRequest | PlainMessage<UnsubscribeRequest> | undefined): boolean {
    return proto3.util.equals(UnsubscribeRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.UnsubscribeResponse
 */
export class UnsubscribeResponse extends Message<UnsubscribeResponse> {
  constructor(data?: PartialMessage<UnsubscribeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.UnsubscribeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnsubscribeResponse {
    return new UnsubscribeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnsubscribeResponse {
    return new UnsubscribeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnsubscribeResponse {
    return new UnsubscribeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UnsubscribeResponse | PlainMessage<UnsubscribeResponse> | undefined, b: UnsubscribeResponse | PlainMessage<UnsubscribeResponse> | undefined): boolean {
    return proto3.util.equals(UnsubscribeResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.GetTotalAmountRequest
 */
export class GetTotalAmountRequest extends Message<GetTotalAmountRequest> {
  constructor(data?: PartialMessage<GetTotalAmountRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetTotalAmountRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTotalAmountRequest {
    return new GetTotalAmountRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTotalAmountRequest {
    return new GetTotalAmountRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTotalAmountRequest {
    return new GetTotalAmountRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTotalAmountRequest | PlainMessage<GetTotalAmountRequest> | undefined, b: GetTotalAmountRequest | PlainMessage<GetTotalAmountRequest> | undefined): boolean {
    return proto3.util.equals(GetTotalAmountRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.GetTotalAmountResponse
 */
export class GetTotalAmountResponse extends Message<GetTotalAmountResponse> {
  /**
   * @generated from field: int32 total_amount = 1;
   */
  totalAmount = 0;

  constructor(data?: PartialMessage<GetTotalAmountResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetTotalAmountResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "total_amount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTotalAmountResponse {
    return new GetTotalAmountResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTotalAmountResponse {
    return new GetTotalAmountResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTotalAmountResponse {
    return new GetTotalAmountResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTotalAmountResponse | PlainMessage<GetTotalAmountResponse> | undefined, b: GetTotalAmountResponse | PlainMessage<GetTotalAmountResponse> | undefined): boolean {
    return proto3.util.equals(GetTotalAmountResponse, a, b);
  }
}

/**
 * @generated from message donation.v1.GetContributorsRequest
 */
export class GetContributorsRequest extends Message<GetContributorsRequest> {
  constructor(data?: PartialMessage<GetContributorsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetContributorsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetContributorsRequest {
    return new GetContributorsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetContributorsRequest {
    return new GetContributorsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetContributorsRequest {
    return new GetContributorsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetContributorsRequest | PlainMessage<GetContributorsRequest> | undefined, b: GetContributorsRequest | PlainMessage<GetContributorsRequest> | undefined): boolean {
    return proto3.util.equals(GetContributorsRequest, a, b);
  }
}

/**
 * @generated from message donation.v1.GetContributorsResponse
 */
export class GetContributorsResponse extends Message<GetContributorsResponse> {
  /**
   * @generated from field: repeated donation.v1.Contributor contributors = 1;
   */
  contributors: Contributor[] = [];

  constructor(data?: PartialMessage<GetContributorsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "donation.v1.GetContributorsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contributors", kind: "message", T: Contributor, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetContributorsResponse {
    return new GetContributorsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetContributorsResponse {
    return new GetContributorsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetContributorsResponse {
    return new GetContributorsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetContributorsResponse | PlainMessage<GetContributorsResponse> | undefined, b: GetContributorsResponse | PlainMessage<GetContributorsResponse> | undefined): boolean {
    return proto3.util.equals(GetContributorsResponse, a, b);
  }
}

