import { Component, OnInit, Input } from '@angular/core';
import {GrpcService} from '../grpc.service';
import {RegisterRequest, RegisterResponse, OtpRequest, OtpResponse}
 from 'proto/generated/proto/auth-messages_pb';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
 @Input() username : string;
 @Input() password : string;
 @Input() otp : string;
 otpFlag = false;
 buttonName = "Get OTP";

  constructor(private grpcService: GrpcService) { }

  ngOnInit(): void {

  }

  action() : void {
    if(this.otpFlag) this.verifyOtp();
    else this.registerUser();
  }

  registerUser() : void {
    if(this.username == null || this.password == null) {
      alert("Enter Username and Password");
    } else {
      const request = new RegisterRequest();
      request.setUsername(this.username);
      request.setPassword(this.password);
      this.grpcService.register(request).on(
        "data", (resp : RegisterResponse) => {
          this.otpFlag = true;
          this.buttonName = "Verify OTP";
          alert(resp.getRegisterresponsemapMap().get("Response"));
        }
      );
    }
  }

  verifyOtp() : void {
    const request = new OtpRequest();
    request.setUsername(this.username);
    request.setOtp(this.otp);
    this.grpcService.verifyOtp(request).on(
      "data", (resp : OtpResponse) => {
        this.otpFlag = false;
        this.buttonName = "Get OTP";
        alert(resp.getOtpresponsemapMap().get("Response"));
      }
    );
  }
}
