import { Component, OnInit } from '@angular/core';
import { RestService } from './services/rest.service';
import { Game } from './models/game';
import { GamesListDataSource } from './utils/games-list-data-source';
import { PlayerConnectionService } from './services/player-connection.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  dataSource: GamesListDataSource;
  displayedColumns = ['name', 'join'];

  constructor(private restService: RestService,
              private playerConnection: PlayerConnectionService) {}

  ngOnInit() {
    this.dataSource = new GamesListDataSource(this.restService);
  }

  public startNewGame() {
    this.restService.createNewGame().subscribe((game: Game) => {
      console.log(game);
      this.dataSource.refresh();
    });
  }

  public joinGame(gameName: string) {
    this.playerConnection.joinGame(gameName);
  }

  public sendMessage() {
    this.playerConnection.sendMessage('hello!');
  }
}
