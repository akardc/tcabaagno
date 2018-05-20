import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { MatButtonModule, MatToolbarModule, MatTableModule } from '@angular/material';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { ServicesModule } from './services/services.module';
import { HttpClientModule } from '@angular/common/http';
import { GamesListComponent } from './components/games-list/games-list.component';
import { GameAdminComponent } from './components/game-admin/game-admin.component';

const appRoutes: Routes = [
  {path: '', redirectTo: '/games', pathMatch: 'full'},
  {path: 'games', component: GamesListComponent},
  {path: 'games/:gameId/admin', component: GameAdminComponent}
];

@NgModule({
  declarations: [
    AppComponent,
    GamesListComponent,
    GameAdminComponent
  ],
  imports: [
    RouterModule.forRoot(appRoutes),
    ServicesModule,
    HttpClientModule,
    BrowserModule,
    MatButtonModule,
    MatToolbarModule,
    MatTableModule
  ],
  providers: [
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
