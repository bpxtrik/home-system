import { Component, inject, signal } from '@angular/core';
import { LoginService } from './login.service';
import { LoginRequest } from './login.dto';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [],
  templateUrl: './login.html',
  styleUrl: './login.css',
  providers: [LoginService],
})
export class Login {
  loginService = inject(LoginService);
  username = signal('');
  password = signal('');
  errorMessage = signal('');
  router = inject(Router);

  login() {
    if (this.usernameNull() || this.passwordNull()) {
      this.errorMessage.set('Fill out the fields!');
      return;
    } else {
      this.errorMessage.set('');
    }

    let request: LoginRequest = { username: this.username(), password: this.password() };
    this.loginService.login(request).subscribe({
      next: (res) => {
        this.errorMessage.set('');
        this.router.navigate(["/dashboard"])
      },
      error: (error) => {
        if (error.status == 401) {
          this.errorMessage.set('Login failed. Try again!');
        }
      },
    });
  }

  private usernameNull(): boolean {
    return this.username() == null || this.username() == '';
  }

  private passwordNull(): boolean {
    return this.password() == null || this.password() == '';
  }
}
