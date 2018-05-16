import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RestService } from './rest.service';
import { PlayerConnectionService } from './player-connection.service';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [],
  providers: [
    RestService,
    PlayerConnectionService
  ]
})
export class ServicesModule { }
