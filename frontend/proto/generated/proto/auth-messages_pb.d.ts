import * as jspb from "google-protobuf"

export class LoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class RegisterRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): RegisterRequest;

  getPassword(): string;
  setPassword(value: string): RegisterRequest;

  getPhone(): string;
  setPhone(value: string): RegisterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    username: string,
    password: string,
    phone: string,
  }
}

export class LoginResponse extends jspb.Message {
  getLoginresponsemapMap(): jspb.Map<string, string>;
  clearLoginresponsemapMap(): LoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    loginresponsemapMap: Array<[string, string]>,
  }
}

export class RegisterResponse extends jspb.Message {
  getRegisterresponsemapMap(): jspb.Map<string, string>;
  clearRegisterresponsemapMap(): RegisterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterResponse): RegisterResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterResponse;
  static deserializeBinaryFromReader(message: RegisterResponse, reader: jspb.BinaryReader): RegisterResponse;
}

export namespace RegisterResponse {
  export type AsObject = {
    registerresponsemapMap: Array<[string, string]>,
  }
}

export class OtpRequest extends jspb.Message {
  getOtp(): string;
  setOtp(value: string): OtpRequest;

  getUsername(): string;
  setUsername(value: string): OtpRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtpRequest.AsObject;
  static toObject(includeInstance: boolean, msg: OtpRequest): OtpRequest.AsObject;
  static serializeBinaryToWriter(message: OtpRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtpRequest;
  static deserializeBinaryFromReader(message: OtpRequest, reader: jspb.BinaryReader): OtpRequest;
}

export namespace OtpRequest {
  export type AsObject = {
    otp: string,
    username: string,
  }
}

export class OtpResponse extends jspb.Message {
  getOtpresponsemapMap(): jspb.Map<string, string>;
  clearOtpresponsemapMap(): OtpResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtpResponse.AsObject;
  static toObject(includeInstance: boolean, msg: OtpResponse): OtpResponse.AsObject;
  static serializeBinaryToWriter(message: OtpResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtpResponse;
  static deserializeBinaryFromReader(message: OtpResponse, reader: jspb.BinaryReader): OtpResponse;
}

export namespace OtpResponse {
  export type AsObject = {
    otpresponsemapMap: Array<[string, string]>,
  }
}

