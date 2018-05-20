import { Component, OnInit } from '@angular/core';
import { PlayerConnectionService } from '../../services/player-connection.service';
import { Message } from '../../models/game';

@Component({
  selector: 'app-game-admin',
  templateUrl: './game-admin.component.html',
  styleUrls: ['./game-admin.component.css']
})
export class GameAdminComponent implements OnInit {

  constructor(private playerConnection: PlayerConnectionService) { }

  ngOnInit() {
  }

  get numPlayers() {
    return this.playerConnection.gameInfo.numPlayers;
  }

  get gameName() {
    return this.playerConnection.gameInfo.id;
  }

  startGame() {
    const startMessage = new Message('start-game');
    this.playerConnection.sendMessage(startMessage);
  }
}
