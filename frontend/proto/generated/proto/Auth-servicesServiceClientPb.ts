/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_auth$messages_pb from '../proto/auth-messages_pb';

export class AuthenticationServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth$messages_pb.LoginResponse,
    (request: proto_auth$messages_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    proto_auth$messages_pb.LoginResponse.deserializeBinary
  );

  login(
    request: proto_auth$messages_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth$messages_pb.LoginResponse>;

  login(
    request: proto_auth$messages_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth$messages_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<proto_auth$messages_pb.LoginResponse>;

  login(
    request: proto_auth$messages_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth$messages_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        new URL('/proto.AuthenticationService/Login', this.hostname_).toString(),
        request,
        metadata || {},
        this.methodInfoLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.AuthenticationService/Login',
    request,
    metadata || {},
    this.methodInfoLogin);
  }

  methodInfoRegister = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth$messages_pb.RegisterResponse,
    (request: proto_auth$messages_pb.RegisterRequest) => {
      return request.serializeBinary();
    },
    proto_auth$messages_pb.RegisterResponse.deserializeBinary
  );

  register(
    request: proto_auth$messages_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth$messages_pb.RegisterResponse>;

  register(
    request: proto_auth$messages_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth$messages_pb.RegisterResponse) => void): grpcWeb.ClientReadableStream<proto_auth$messages_pb.RegisterResponse>;

  register(
    request: proto_auth$messages_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth$messages_pb.RegisterResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        new URL('/proto.AuthenticationService/Register', this.hostname_).toString(),
        request,
        metadata || {},
        this.methodInfoRegister,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.AuthenticationService/Register',
    request,
    metadata || {},
    this.methodInfoRegister);
  }

  methodInfoVerifyOtp = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth$messages_pb.OtpResponse,
    (request: proto_auth$messages_pb.OtpRequest) => {
      return request.serializeBinary();
    },
    proto_auth$messages_pb.OtpResponse.deserializeBinary
  );

  verifyOtp(
    request: proto_auth$messages_pb.OtpRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth$messages_pb.OtpResponse>;

  verifyOtp(
    request: proto_auth$messages_pb.OtpRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth$messages_pb.OtpResponse) => void): grpcWeb.ClientReadableStream<proto_auth$messages_pb.OtpResponse>;

  verifyOtp(
    request: proto_auth$messages_pb.OtpRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth$messages_pb.OtpResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        new URL('/proto.AuthenticationService/VerifyOtp', this.hostname_).toString(),
        request,
        metadata || {},
        this.methodInfoVerifyOtp,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.AuthenticationService/VerifyOtp',
    request,
    metadata || {},
    this.methodInfoVerifyOtp);
  }

}

