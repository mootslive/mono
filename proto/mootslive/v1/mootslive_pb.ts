// @generated by protoc-gen-es v0.3.0 with parameter "target=ts"
// @generated from file mootslive/v1/mootslive.proto (package mootslive.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message mootslive.v1.GetStatusRequest
 */
export class GetStatusRequest extends Message<GetStatusRequest> {
  constructor(data?: PartialMessage<GetStatusRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.GetStatusRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStatusRequest {
    return new GetStatusRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStatusRequest {
    return new GetStatusRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStatusRequest {
    return new GetStatusRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetStatusRequest | PlainMessage<GetStatusRequest> | undefined, b: GetStatusRequest | PlainMessage<GetStatusRequest> | undefined): boolean {
    return proto3.util.equals(GetStatusRequest, a, b);
  }
}

/**
 * @generated from message mootslive.v1.GetStatusResponse
 */
export class GetStatusResponse extends Message<GetStatusResponse> {
  /**
   * @generated from field: string x_clacks_overhead = 1;
   */
  xClacksOverhead = "";

  constructor(data?: PartialMessage<GetStatusResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.GetStatusResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "x_clacks_overhead", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStatusResponse {
    return new GetStatusResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStatusResponse {
    return new GetStatusResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStatusResponse {
    return new GetStatusResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetStatusResponse | PlainMessage<GetStatusResponse> | undefined, b: GetStatusResponse | PlainMessage<GetStatusResponse> | undefined): boolean {
    return proto3.util.equals(GetStatusResponse, a, b);
  }
}

/**
 * @generated from message mootslive.v1.GetMeRequest
 */
export class GetMeRequest extends Message<GetMeRequest> {
  constructor(data?: PartialMessage<GetMeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.GetMeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMeRequest {
    return new GetMeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMeRequest {
    return new GetMeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMeRequest {
    return new GetMeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetMeRequest | PlainMessage<GetMeRequest> | undefined, b: GetMeRequest | PlainMessage<GetMeRequest> | undefined): boolean {
    return proto3.util.equals(GetMeRequest, a, b);
  }
}

/**
 * @generated from message mootslive.v1.GetMeResponse
 */
export class GetMeResponse extends Message<GetMeResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 2;
   */
  createdAt?: Timestamp;

  constructor(data?: PartialMessage<GetMeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.GetMeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "created_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMeResponse {
    return new GetMeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMeResponse {
    return new GetMeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMeResponse {
    return new GetMeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetMeResponse | PlainMessage<GetMeResponse> | undefined, b: GetMeResponse | PlainMessage<GetMeResponse> | undefined): boolean {
    return proto3.util.equals(GetMeResponse, a, b);
  }
}

/**
 * @generated from message mootslive.v1.BeginTwitterAuthRequest
 */
export class BeginTwitterAuthRequest extends Message<BeginTwitterAuthRequest> {
  constructor(data?: PartialMessage<BeginTwitterAuthRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.BeginTwitterAuthRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginTwitterAuthRequest {
    return new BeginTwitterAuthRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginTwitterAuthRequest {
    return new BeginTwitterAuthRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginTwitterAuthRequest {
    return new BeginTwitterAuthRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BeginTwitterAuthRequest | PlainMessage<BeginTwitterAuthRequest> | undefined, b: BeginTwitterAuthRequest | PlainMessage<BeginTwitterAuthRequest> | undefined): boolean {
    return proto3.util.equals(BeginTwitterAuthRequest, a, b);
  }
}

/**
 * @generated from message mootslive.v1.BeginTwitterAuthResponse
 */
export class BeginTwitterAuthResponse extends Message<BeginTwitterAuthResponse> {
  /**
   * @generated from field: string redirect_url = 1;
   */
  redirectUrl = "";

  constructor(data?: PartialMessage<BeginTwitterAuthResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.BeginTwitterAuthResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "redirect_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginTwitterAuthResponse {
    return new BeginTwitterAuthResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginTwitterAuthResponse {
    return new BeginTwitterAuthResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginTwitterAuthResponse {
    return new BeginTwitterAuthResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BeginTwitterAuthResponse | PlainMessage<BeginTwitterAuthResponse> | undefined, b: BeginTwitterAuthResponse | PlainMessage<BeginTwitterAuthResponse> | undefined): boolean {
    return proto3.util.equals(BeginTwitterAuthResponse, a, b);
  }
}

/**
 * @generated from message mootslive.v1.FinishTwitterAuthRequest
 */
export class FinishTwitterAuthRequest extends Message<FinishTwitterAuthRequest> {
  constructor(data?: PartialMessage<FinishTwitterAuthRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.FinishTwitterAuthRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FinishTwitterAuthRequest {
    return new FinishTwitterAuthRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FinishTwitterAuthRequest {
    return new FinishTwitterAuthRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FinishTwitterAuthRequest {
    return new FinishTwitterAuthRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FinishTwitterAuthRequest | PlainMessage<FinishTwitterAuthRequest> | undefined, b: FinishTwitterAuthRequest | PlainMessage<FinishTwitterAuthRequest> | undefined): boolean {
    return proto3.util.equals(FinishTwitterAuthRequest, a, b);
  }
}

/**
 * @generated from message mootslive.v1.FinishTwitterAuthResponse
 */
export class FinishTwitterAuthResponse extends Message<FinishTwitterAuthResponse> {
  constructor(data?: PartialMessage<FinishTwitterAuthResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "mootslive.v1.FinishTwitterAuthResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FinishTwitterAuthResponse {
    return new FinishTwitterAuthResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FinishTwitterAuthResponse {
    return new FinishTwitterAuthResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FinishTwitterAuthResponse {
    return new FinishTwitterAuthResponse().fromJsonString(jsonString, options);
  }

  static equals(a: FinishTwitterAuthResponse | PlainMessage<FinishTwitterAuthResponse> | undefined, b: FinishTwitterAuthResponse | PlainMessage<FinishTwitterAuthResponse> | undefined): boolean {
    return proto3.util.equals(FinishTwitterAuthResponse, a, b);
  }
}
