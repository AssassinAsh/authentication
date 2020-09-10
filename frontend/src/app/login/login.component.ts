import { Component, OnInit, Input } from '@angular/core';
import {LoginRequest, LoginResponse} from 'proto/generated/proto/auth-messages_pb';
import {GrpcService} from '../grpc.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  @Input() username : string;
  @Input() password : string;

  constructor(private grpcService : GrpcService) { }

  ngOnInit(): void {
  }

  userLogin() : void {
   const request = new LoginRequest();
   request.setUsername(this.username);
   request.setPassword(this.password);

    this.grpcService.login(request).on(
      "data", (resp: LoginResponse) => {
        if(resp != null ){
          alert(resp.getLoginresponsemapMap().get("Response"));
        }
      }
    );

  }
}
