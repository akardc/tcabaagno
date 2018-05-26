export class Game {
    id: string;
    status: string;
}

export class RegistrationResponse {
    id: string;
    gameId: string;
}

export class GameInfo {
    id: string;
    numPlayers: number;
}

export class SubmitQuestionsPayload {
    who: string;
    what: string;
    when: string;
    where: string;
    why: string;
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
