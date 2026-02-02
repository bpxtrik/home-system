import { Component, inject, OnInit, signal } from '@angular/core';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';
import { GetResponse } from '../shared/dto';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
  providers: [DashboardService],
})
export class Dashboard implements OnInit {
  dashboardService = inject(DashboardService);

  data = signal<GetResponse | null>(null);

  page = 0;
  limit = 10;

  currentPage = signal(0);
  totalPages = signal(0);

  loading = signal(true)

  ngOnInit(): void {
    this.getMotions();
  }

  goToPage(page: number): void {

    if (page < 0 || page >= this.totalPages()) return;

    this.page = page;
    this.currentPage.set(page);
    this.getMotions();
  }

  getMotions(): void {
    this.loading.set(true)
    this.data()?.data.splice(0, this.data()?.data.length)

    this.dashboardService.getMotions(this.page, this.limit).subscribe({
      next: (res) => {
        this.data.set(res);
        this.totalPages.set(res.total_pages);
        this.loading.set(false)
      },
      error: (err) => {
        console.log(err);
        this.loading.set(false)
      },
    });
  }
}
