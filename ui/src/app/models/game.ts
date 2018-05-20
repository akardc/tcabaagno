export class Game {
    id: string;
}

export class RegistrationResponse {
    id: string;
    gameId: string;
}

export class GameInfo {
    id: string;
    numPlayers: number;
}

export class Message<T> {
    type: string;
    payload: T;

    constructor(type: string, payload?: T) {
        this.type = type;
        if (payload) {
            this.payload = payload;
        }
    }
}
