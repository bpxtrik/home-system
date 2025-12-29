import { Routes } from '@angular/router';
import { Login } from './login/login';
import { Dashboard } from './dashboard/dashboard';
import { AuthGuard } from './shared/auth.guard';

export const routes: Routes = [
  {path: "", redirectTo: "login", pathMatch: "full"},
  {path: "dashboard", component: Dashboard, canActivate: [AuthGuard]},
  {path: "login", component: Login},
  {path: "**", redirectTo: "login"}
];
