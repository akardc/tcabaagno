import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Game } from '../models/game';
import { Response } from '../models/response';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RestService {

  constructor(private httpClient: HttpClient) { }

  public createNewGame(): Observable<Game> {
    const url = environment.serverUrl + '/games';
    return this.httpClient.post(url, null).pipe(map((res: Response<Game>) => {
      return res.payload;
    }));
  }

  public getActiveGames(): Observable<Game[]> {
    const url = environment.serverUrl + '/games';
    return this.httpClient.get(url).pipe(map((res: Response<Game[]>) => {
      return res.payload;
    }));
  }
}
