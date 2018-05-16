import { Injectable } from '@angular/core';
import { WebSocketSubject } from 'rxjs/webSocket';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PlayerConnectionService {

  connection: WebSocketSubject<any>;

  constructor() { }

  public joinGame(name: string) {
    if (this.connection) {
      return;
    }
    const baseUrl = environment.serverUrl.replace('http', 'ws');
    const url = `${baseUrl}/games/${name}/join`;
    this.connection = new WebSocketSubject(url);
    this.connection.subscribe((msg) => {
      console.log(`incoming message: ${msg}`);
    });
  }

  public sendMessage(message: string) {
    if (!this.connection) {
      return;
    }

    this.connection.next(message);
  }
}
