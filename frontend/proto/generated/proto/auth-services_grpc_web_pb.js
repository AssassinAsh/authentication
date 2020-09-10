/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var proto_auth$messages_pb = require('../proto/auth-messages_pb.js')
const proto = {};
proto.proto = require('./auth-services_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AuthenticationServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AuthenticationServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.LoginRequest,
 *   !proto.proto.LoginResponse>}
 */
const methodDescriptor_AuthenticationService_Login = new grpc.web.MethodDescriptor(
  '/proto.AuthenticationService/Login',
  grpc.web.MethodType.UNARY,
  proto_auth$messages_pb.LoginRequest,
  proto_auth$messages_pb.LoginResponse,
  /**
   * @param {!proto.proto.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto_auth$messages_pb.LoginResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.LoginRequest,
 *   !proto.proto.LoginResponse>}
 */
const methodInfo_AuthenticationService_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto_auth$messages_pb.LoginResponse,
  /**
   * @param {!proto.proto.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto_auth$messages_pb.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.proto.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AuthenticationServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthenticationService/Login',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_Login,
      callback);
};


/**
 * @param {!proto.proto.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.LoginResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.AuthenticationServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthenticationService/Login',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.RegisterRequest,
 *   !proto.proto.RegisterResponse>}
 */
const methodDescriptor_AuthenticationService_Register = new grpc.web.MethodDescriptor(
  '/proto.AuthenticationService/Register',
  grpc.web.MethodType.UNARY,
  proto_auth$messages_pb.RegisterRequest,
  proto_auth$messages_pb.RegisterResponse,
  /**
   * @param {!proto.proto.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto_auth$messages_pb.RegisterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.RegisterRequest,
 *   !proto.proto.RegisterResponse>}
 */
const methodInfo_AuthenticationService_Register = new grpc.web.AbstractClientBase.MethodInfo(
  proto_auth$messages_pb.RegisterResponse,
  /**
   * @param {!proto.proto.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto_auth$messages_pb.RegisterResponse.deserializeBinary
);


/**
 * @param {!proto.proto.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.RegisterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.RegisterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AuthenticationServiceClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthenticationService/Register',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_Register,
      callback);
};


/**
 * @param {!proto.proto.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.RegisterResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.AuthenticationServicePromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthenticationService/Register',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.OtpRequest,
 *   !proto.proto.OtpResponse>}
 */
const methodDescriptor_AuthenticationService_VerifyOtp = new grpc.web.MethodDescriptor(
  '/proto.AuthenticationService/VerifyOtp',
  grpc.web.MethodType.UNARY,
  proto_auth$messages_pb.OtpRequest,
  proto_auth$messages_pb.OtpResponse,
  /**
   * @param {!proto.proto.OtpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto_auth$messages_pb.OtpResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.OtpRequest,
 *   !proto.proto.OtpResponse>}
 */
const methodInfo_AuthenticationService_VerifyOtp = new grpc.web.AbstractClientBase.MethodInfo(
  proto_auth$messages_pb.OtpResponse,
  /**
   * @param {!proto.proto.OtpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto_auth$messages_pb.OtpResponse.deserializeBinary
);


/**
 * @param {!proto.proto.OtpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.OtpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.OtpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AuthenticationServiceClient.prototype.verifyOtp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AuthenticationService/VerifyOtp',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_VerifyOtp,
      callback);
};


/**
 * @param {!proto.proto.OtpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.OtpResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.AuthenticationServicePromiseClient.prototype.verifyOtp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AuthenticationService/VerifyOtp',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_VerifyOtp);
};


module.exports = proto.proto;

