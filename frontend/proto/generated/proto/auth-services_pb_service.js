// package: proto
// file: proto/auth-services.proto

var proto_auth_services_pb = require("../proto/auth-services_pb");
var proto_auth_messages_pb = require("../proto/auth-messages_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var AuthenticationService = (function () {
  function AuthenticationService() {}
  AuthenticationService.serviceName = "proto.AuthenticationService";
  return AuthenticationService;
}());

AuthenticationService.Login = {
  methodName: "Login",
  service: AuthenticationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_auth_messages_pb.LoginRequest,
  responseType: proto_auth_messages_pb.LoginResponse
};

AuthenticationService.Register = {
  methodName: "Register",
  service: AuthenticationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_auth_messages_pb.RegisterRequest,
  responseType: proto_auth_messages_pb.RegisterResponse
};

AuthenticationService.VerifyOtp = {
  methodName: "VerifyOtp",
  service: AuthenticationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_auth_messages_pb.OtpRequest,
  responseType: proto_auth_messages_pb.OtpResponse
};

exports.AuthenticationService = AuthenticationService;

function AuthenticationServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AuthenticationServiceClient.prototype.login = function login(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthenticationService.Login, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthenticationServiceClient.prototype.register = function register(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthenticationService.Register, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthenticationServiceClient.prototype.verifyOtp = function verifyOtp(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthenticationService.VerifyOtp, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.AuthenticationServiceClient = AuthenticationServiceClient;

