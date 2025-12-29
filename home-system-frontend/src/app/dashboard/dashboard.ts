import { Component, inject, OnInit, signal } from '@angular/core';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';
import { Motion } from '../shared/dto';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
  providers: [DashboardService]
})
export class Dashboard implements OnInit {
  dashboardService = inject(DashboardService);
  data = signal<Motion[]>([])

  ngOnInit(): void {
    this.dashboardService.getMotions().subscribe({
      next: (res) => {
        this.data.set(res)
      },
      error: (err) => {
        console.log(err)
      }
    })
  }
}
