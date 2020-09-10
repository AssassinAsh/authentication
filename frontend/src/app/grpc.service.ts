import { Injectable } from '@angular/core';
import {LoginResponse, LoginRequest, RegisterRequest, RegisterResponse, OtpRequest, OtpResponse} from 'proto/generated/proto/auth-messages_pb';
import {AuthenticationServiceClient} from 'proto/generated/proto/Auth-servicesServiceClientPb';
import * as grpcWeb from 'grpc-web';
import {Observable} from 'rxjs';
import { grpc } from '@improbable-eng/grpc-web';

@Injectable({
  providedIn: 'root'
})
export class GrpcService {
  client : AuthenticationServiceClient;

  constructor() {
    this.client = new AuthenticationServiceClient('http://localhost:6565');
  }

  login(req: LoginRequest) : grpcWeb.ClientReadableStream<LoginResponse> {

    return this.client.login(req, {'Content-Type':'application/grpc-web+proto'},
      (err: grpcWeb.Error, response: LoginResponse) => {
        if(err != null) {
          console.log(err.message);
        } if(response != null) {
          console.log(response.getLoginresponsemapMap().get("Response Code"));
        }
    }).on("end",()=> {
      console.log("Login : Stream Complete");
    });
  }

  register(req : RegisterRequest) : grpcWeb.ClientReadableStream<RegisterResponse> {
    return this.client.register(req, {'Content-Type':'application/grpc-web+proto'},
      (err: grpcWeb.Error, response : RegisterResponse) => {
        if(err != null) {
          console.log(err.message);
        } if(response != null) {
          console.log(response.getRegisterresponsemapMap().get("Response Code"));
        }
      }).on("end", () => {
        console.log("Register : Stream Complete");
      });
  }

  verifyOtp(req : OtpRequest) : grpcWeb.ClientReadableStream<OtpResponse> {
    return this.client.verifyOtp(req, {'Content-Type':'application/grpc-web+proto'},
      (err: grpcWeb.Error, response : OtpResponse) => {
        if(err != null) {
          console.log(err.message);
        } if(response != null) {
          console.log(response.getOtpresponsemapMap().get("Response Code"));
        }
      }).on("end", () => {
        console.log("Verify Otp : Stream Complete");
      });
  }
}
