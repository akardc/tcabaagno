import { DataSource } from '@angular/cdk/table';
import { Game } from '../models/game';
import { Observable, BehaviorSubject } from 'rxjs';
import { RestService } from '../services/rest.service';

export class GamesListDataSource implements DataSource<Game> {

    private gamesSubject: BehaviorSubject<Game[]>;

    constructor(private restService: RestService) {}

    public connect(): Observable<Game[]> {
        if (this.gamesSubject) {
            return this.gamesSubject.asObservable();
        }

        this.gamesSubject = new BehaviorSubject<Game[]>([]);
        this.getActiveGames();
        return this.gamesSubject.asObservable();
    }

    public disconnect() {
        // no disconnect logic required
    }

    public refresh() {
        this.getActiveGames();
    }

    private getActiveGames() {
        this.restService.getActiveGames().subscribe((games) => {
            this.gamesSubject.next(games);
        });
    }
}
