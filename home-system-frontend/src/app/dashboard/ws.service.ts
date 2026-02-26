import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Motion } from '../shared/dto';
import { environment } from '../../environments/environment';

@Injectable({ providedIn: 'root' })
export class WsService {
  private ws?: WebSocket;
  private subject = new Subject<Motion>();
  private wsUrl = environment.wsUrl;

  connect() {
    this.ws = new WebSocket(`${this.wsUrl}`);

    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data) as Motion;
      this.subject.next(data);
    };
  }

  get messages() {
    return this.subject.asObservable();
  }
}
