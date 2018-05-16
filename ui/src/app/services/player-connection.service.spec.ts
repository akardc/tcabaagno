import { TestBed, inject } from '@angular/core/testing';

import { PlayerConnectionService } from './player-connection.service';

describe('PlayerConnectionService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [PlayerConnectionService]
    });
  });

  it('should be created', inject([PlayerConnectionService], (service: PlayerConnectionService) => {
    expect(service).toBeTruthy();
  }));
});
