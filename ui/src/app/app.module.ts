import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { MatButtonModule, MatToolbarModule, MatTableModule, MatMenuModule, MatIconModule, MatStepperModule, MatInputModule } from '@angular/material';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { ServicesModule } from './services/services.module';
import { HttpClientModule } from '@angular/common/http';
import { GamesListComponent } from './components/games-list/games-list.component';
import { GameAdminComponent } from './components/game-admin/game-admin.component';
import { QuestionsFormComponent } from './components/questions-form/questions-form.component';

const appRoutes: Routes = [
  { path: '', redirectTo: '/games', pathMatch: 'full' },
  { path: 'games', component: GamesListComponent },
  { path: 'games/:gameId/admin', component: GameAdminComponent },
  { path: 'games/:gameId/questions', component: QuestionsFormComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    GamesListComponent,
    GameAdminComponent,
    QuestionsFormComponent
  ],
  imports: [
    RouterModule.forRoot(appRoutes),
    ServicesModule,
    HttpClientModule,
    BrowserModule,
    MatButtonModule,
    MatToolbarModule,
    MatTableModule,
    MatMenuModule,
    MatIconModule,
    MatStepperModule,
    MatInputModule,
    BrowserAnimationsModule
  ],
  providers: [
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
