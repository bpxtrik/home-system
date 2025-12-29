import { HttpClient } from "@angular/common/http";
import { inject } from "@angular/core";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";
import { Motion } from "../shared/dto";

export class DashboardService {
  private http = inject(HttpClient)
  private apiUrl = environment.apiUrl

  getMotions(): Observable<Motion[]> {
    return this.http.get<Motion[]>(this.apiUrl + "motions", {withCredentials: true})
  }
}
