import { Injectable } from '@angular/core';
import { WebSocketSubject } from 'rxjs/webSocket';
import { environment } from '../../environments/environment';
import { RegistrationResponse, Message, GameInfo } from '../models/game';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class PlayerConnectionService {

  connection: WebSocketSubject<any>;

  private playerInfo: RegistrationResponse;
  private gameInfoData: GameInfo;

  constructor(private router: Router) { }

  public get gameInfo() {
    if (this.gameInfoData) {
      return this.gameInfoData;
    }
    return new GameInfo();
  }

  public joinGame(name: string) {
    if (this.connection) {
      return;
    }
    const baseUrl = environment.serverUrl.replace('http', 'ws');
    const url = `${baseUrl}/games/${name}/join`;
    this.connection = new WebSocketSubject(url);
    this.connection.subscribe((msg) => {
      console.log('incoming message', msg);
      this.handleIncomingMessage(msg);
    });
  }

  public sendMessage(message: Message<any>) {
    if (!this.connection) {
      return;
    }

    this.connection.next(message);
  }

  private handleIncomingMessage(message: Message<any>) {
    switch (message.type) {
      case 'player-joined':
        this.setPlayerInfo(message);
        break;
      case 'joined-game':
        this.joinedGame(message);
        break;
      case 'game-updated':
        this.gameUpdated(message);
        break;
    }
  }

  private setPlayerInfo(message: Message<RegistrationResponse>) {
    this.playerInfo = message.payload;
    console.log('player info set', this.playerInfo);
  }

  private joinedGame(message: Message<RegistrationResponse>) {
    this.playerInfo = message.payload;
    this.router.navigate([`/games/${message.payload.gameId}/admin`]);
  }

  private gameUpdated(message: Message<GameInfo>) {
    this.gameInfoData = message.payload;
  }
}
