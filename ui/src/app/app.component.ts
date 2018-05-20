import { Component, OnInit } from '@angular/core';
import { RestService } from './services/rest.service';
import { Game, Message } from './models/game';
import { GamesListDataSource } from './utils/games-list-data-source';
import { PlayerConnectionService } from './services/player-connection.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
}
