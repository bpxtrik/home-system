import { HttpClient } from "@angular/common/http";
import { inject } from "@angular/core";
import { environment } from "../../environments/environment";
import { LoginRequest } from "./login.dto";
import { Observable } from "rxjs";
import { Response } from "../shared/dto";

export class LoginService {
  private http = inject(HttpClient)
  private apiUrl = environment.apiUrl


  login(loginRequest: LoginRequest): Observable<Response> {
    return this.http.post<Response>(this.apiUrl + "login", loginRequest, {withCredentials: true})
  }
}
