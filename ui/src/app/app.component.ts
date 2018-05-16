import { Component, OnInit } from '@angular/core';
import {WebSocketSubject} from 'rxjs/webSocket';
import { RestService } from './services/rest.service';
import { Game } from './models/game';
import { GamesListDataSource } from './utils/games-list-data-source';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'app';
  ws: WebSocketSubject<any>;

  games: Game[] = [];
  dataSource: GamesListDataSource;
  displayedColumns = ['name'];

  constructor(private restService: RestService) {}

  ngOnInit() {
    // this.ws = new WebSocketSubject('ws://localhost:8080/games/Forgerdiamond/join');
    // console.log('connecting to ws', this.ws);
    // this.ws.subscribe(
    //   (msg) => console.log('incoming msg', msg),
    //   (err) => console.log('incoming err', err)
    // );
    this.dataSource = new GamesListDataSource(this.restService);
  }

  public send() {
    this.ws.next('hello server');
  }

  public startNewGame() {
    this.restService.createNewGame().subscribe((game: Game) => {
      console.log(game);
      this.dataSource.refresh();
    });
  }
}
