import { HttpClient } from "@angular/common/http";
import { environment } from "../../environments/environment";
import { inject, Injectable } from "@angular/core";

@Injectable({providedIn : 'root'})
export class AuthService {
  private http = inject(HttpClient)
  private apiUrl = environment.apiUrl

  check() {
    return this.http.get(this.apiUrl + "me", {withCredentials: true})
  }
}
