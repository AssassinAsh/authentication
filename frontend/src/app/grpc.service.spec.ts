import { TestBed } from '@angular/core/testing';

import { GrpcService } from './grpc.service';

describe('GrpcService', () => {
  let service: GrpcService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GrpcService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
