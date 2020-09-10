import { Component, OnInit, Input } from '@angular/core';
import {GrpcService} from '../grpc.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
 @Input() username : string;
 @Input() password : string;
 @Input() otp : string;

  constructor(private grpcService: GrpcService) { }

  ngOnInit(): void {
  }

  registerUser() : void {
    if(this.username == null || this.password == null || this.otp == null) {
      
    }
  }

}
