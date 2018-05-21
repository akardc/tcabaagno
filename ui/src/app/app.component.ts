import { Component, OnInit } from '@angular/core';
import { RestService } from './services/rest.service';
import { Game, Message } from './models/game';
import { GamesListDataSource } from './utils/games-list-data-source';
import { PlayerConnectionService } from './services/player-connection.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(public playerConnection: PlayerConnectionService,
              private router: Router) {}

  public goToAdminPage() {
    if (this.playerConnection.connection) {
      this.router.navigate([`/games/${this.playerConnection.gameInfo.id}/admin`]);
    }
  }

  public goToGameList() {
    this.router.navigate(['/games']);
  }

  public goToCurrentStep() {
    this.playerConnection.openCurrentStep():
  }
}
