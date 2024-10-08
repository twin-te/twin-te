// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file donation/v1/service.proto (package donation.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateOneTimeCheckoutSessionRequest, CreateOneTimeCheckoutSessionResponse, CreateSubscriptionCheckoutSessionRequest, CreateSubscriptionCheckoutSessionResponse, GetActiveSubscriptionRequest, GetActiveSubscriptionResponse, GetPaymentUserRequest, GetPaymentUserResponse, GetTotalAmountRequest, GetTotalAmountResponse, ListContributorsRequest, ListContributorsResponse, ListPaymentHistoriesRequest, ListPaymentHistoriesResponse, ListSubscriptionPlansRequest, ListSubscriptionPlansResponse, UnsubscribeRequest, UnsubscribeResponse, UpdatePaymentUserRequest, UpdatePaymentUserResponse } from "./service_pb.js";
import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * The following error codes are not stated explicitly in the each rpc, but may be returned.
 *   - shared.InvalidArgument
 *   - shared.Unauthenticated
 *   - shared.Unauthorized
 *
 * @generated from service donation.v1.DonationService
 */
export const DonationService = {
  typeName: "donation.v1.DonationService",
  methods: {
    /**
     * @generated from rpc donation.v1.DonationService.CreateOneTimeCheckoutSession
     */
    createOneTimeCheckoutSession: {
      name: "CreateOneTimeCheckoutSession",
      I: CreateOneTimeCheckoutSessionRequest,
      O: CreateOneTimeCheckoutSessionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc donation.v1.DonationService.CreateSubscriptionCheckoutSession
     */
    createSubscriptionCheckoutSession: {
      name: "CreateSubscriptionCheckoutSession",
      I: CreateSubscriptionCheckoutSessionRequest,
      O: CreateSubscriptionCheckoutSessionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc donation.v1.DonationService.GetPaymentUser
     */
    getPaymentUser: {
      name: "GetPaymentUser",
      I: GetPaymentUserRequest,
      O: GetPaymentUserResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc donation.v1.DonationService.UpdatePaymentUser
     */
    updatePaymentUser: {
      name: "UpdatePaymentUser",
      I: UpdatePaymentUserRequest,
      O: UpdatePaymentUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc donation.v1.DonationService.ListPaymentHistories
     */
    listPaymentHistories: {
      name: "ListPaymentHistories",
      I: ListPaymentHistoriesRequest,
      O: ListPaymentHistoriesResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc donation.v1.DonationService.ListSubscriptionPlans
     */
    listSubscriptionPlans: {
      name: "ListSubscriptionPlans",
      I: ListSubscriptionPlansRequest,
      O: ListSubscriptionPlansResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc donation.v1.DonationService.GetActiveSubscription
     */
    getActiveSubscription: {
      name: "GetActiveSubscription",
      I: GetActiveSubscriptionRequest,
      O: GetActiveSubscriptionResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc donation.v1.DonationService.Unsubscribe
     */
    unsubscribe: {
      name: "Unsubscribe",
      I: UnsubscribeRequest,
      O: UnsubscribeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc donation.v1.DonationService.GetTotalAmount
     */
    getTotalAmount: {
      name: "GetTotalAmount",
      I: GetTotalAmountRequest,
      O: GetTotalAmountResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc donation.v1.DonationService.ListContributors
     */
    listContributors: {
      name: "ListContributors",
      I: ListContributorsRequest,
      O: ListContributorsResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
  }
} as const;

