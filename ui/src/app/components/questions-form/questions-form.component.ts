import { Component, OnInit } from '@angular/core';
import { PlayerConnectionService } from '../../services/player-connection.service';
import { Message, SubmitQuestionsPayload } from '../../models/game';
import { FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-questions-form',
  templateUrl: './questions-form.component.html',
  styleUrls: ['./questions-form.component.css']
})
export class QuestionsFormComponent implements OnInit {

  public form: FormGroup;

  constructor(private playerConnection: PlayerConnectionService) { }

  ngOnInit() {
    this.form = new FormGroup({
      'who': new FormControl(),
      'what': new FormControl(),
      'when': new FormControl(),
      'where': new FormControl(),
      'why': new FormControl()
    });
  }

  submitQuestions() {
    const questions = new SubmitQuestionsPayload();
    questions.who = this.form.get('who').value;
    questions.what = this.form.get('what').value;
    questions.when = this.form.get('when').value;
    questions.where = this.form.get('where').value;
    questions.why = this.form.get('why').value;
    this.playerConnection.sendMessage(new Message('submit-questions', questions));
  }
}
