import { Component, OnInit } from '@angular/core';
import { RestService } from '../../services/rest.service';
import { Game, Message } from '../../models/game';
import { GamesListDataSource } from '../../utils/games-list-data-source';
import { PlayerConnectionService } from '../../services/player-connection.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-games-list',
  templateUrl: './games-list.component.html',
  styleUrls: ['./games-list.component.css']
})
export class GamesListComponent implements OnInit {

  dataSource: GamesListDataSource;
  displayedColumns = ['name', 'join'];

  constructor(private restService: RestService,
              private playerConnection: PlayerConnectionService,
              private router: Router) {}

  ngOnInit() {
    this.dataSource = new GamesListDataSource(this.restService);
  }

  public createNewGame() {
    this.restService.createNewGame().subscribe((game: Game) => {
      console.log(game);
      this.joinGame(game.id);
      this.router.navigate([`/games/${game.id}/admin`]);
    });
  }

  public joinGame(gameName: string) {
    this.playerConnection.joinGame(gameName);
  }
}
