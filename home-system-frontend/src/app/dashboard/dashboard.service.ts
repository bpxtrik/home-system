import { HttpClient } from '@angular/common/http';
import { inject } from '@angular/core';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { GetResponse } from '../shared/dto';

export class DashboardService {
  private http = inject(HttpClient);
  private apiUrl = environment.apiUrl;

  getMotions(page: number, limit: number): Observable<GetResponse> {
    return this.http.get<GetResponse>(this.apiUrl + `motions?page=${page}&limit=${limit}`, {
      withCredentials: true,
    });
  }
}
