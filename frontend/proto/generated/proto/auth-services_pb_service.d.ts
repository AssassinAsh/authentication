// package: proto
// file: proto/auth-services.proto

import * as proto_auth_services_pb from "../proto/auth-services_pb";
import * as proto_auth_messages_pb from "../proto/auth-messages_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AuthenticationServiceLogin = {
  readonly methodName: string;
  readonly service: typeof AuthenticationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_auth_messages_pb.LoginRequest;
  readonly responseType: typeof proto_auth_messages_pb.LoginResponse;
};

type AuthenticationServiceRegister = {
  readonly methodName: string;
  readonly service: typeof AuthenticationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_auth_messages_pb.RegisterRequest;
  readonly responseType: typeof proto_auth_messages_pb.RegisterResponse;
};

type AuthenticationServiceVerifyOtp = {
  readonly methodName: string;
  readonly service: typeof AuthenticationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_auth_messages_pb.OtpRequest;
  readonly responseType: typeof proto_auth_messages_pb.OtpResponse;
};

export class AuthenticationService {
  static readonly serviceName: string;
  static readonly Login: AuthenticationServiceLogin;
  static readonly Register: AuthenticationServiceRegister;
  static readonly VerifyOtp: AuthenticationServiceVerifyOtp;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AuthenticationServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  login(
    requestMessage: proto_auth_messages_pb.LoginRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_auth_messages_pb.LoginResponse|null) => void
  ): UnaryResponse;
  login(
    requestMessage: proto_auth_messages_pb.LoginRequest,
    callback: (error: ServiceError|null, responseMessage: proto_auth_messages_pb.LoginResponse|null) => void
  ): UnaryResponse;
  register(
    requestMessage: proto_auth_messages_pb.RegisterRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_auth_messages_pb.RegisterResponse|null) => void
  ): UnaryResponse;
  register(
    requestMessage: proto_auth_messages_pb.RegisterRequest,
    callback: (error: ServiceError|null, responseMessage: proto_auth_messages_pb.RegisterResponse|null) => void
  ): UnaryResponse;
  verifyOtp(
    requestMessage: proto_auth_messages_pb.OtpRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_auth_messages_pb.OtpResponse|null) => void
  ): UnaryResponse;
  verifyOtp(
    requestMessage: proto_auth_messages_pb.OtpRequest,
    callback: (error: ServiceError|null, responseMessage: proto_auth_messages_pb.OtpResponse|null) => void
  ): UnaryResponse;
}

