import { Component, ElementRef, inject, OnInit, signal, ViewChild } from '@angular/core';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';
import { GetResponse, Motion } from '../shared/dto';
import { WsService } from './ws.service';
import { Toast } from 'bootstrap';

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
  wsService = inject(WsService);

  data = signal<GetResponse | null>(null);

  page = 0;
  limit = 10;

  currentPage = signal(0);
  totalPages = signal(0);

  loading = signal(true);

  @ViewChild('toast') toastEl!: ElementRef;

  ngOnInit(): void {
    this.getMotions();

    this.wsService.connect();

    this.wsService.messages.subscribe((m: Motion) => {
      this.showToast(m);
      this.getMotions();
    });
  }

  goToPage(page: number): void {
    if (page < 0 || page >= this.totalPages()) return;
    this.page = page;
    this.currentPage.set(page);
    this.getMotions();
  }

  getMotions(): void {
    this.loading.set(true);

    this.dashboardService.getMotions(this.page, this.limit).subscribe({
      next: (res: GetResponse) => {
        this.data.set(res);
        this.totalPages.set(res.total_pages);
        this.loading.set(false);
      },
      error: (err) => {
        console.log(err);
        this.loading.set(false);
      },
    });
  }

  showToast(motion?: Motion) {
    if (motion && this.toastEl) {
      this.toastEl.nativeElement.querySelector('.toast-body').textContent =
        `Motion detected at ${new Date(motion.Timestamp).toLocaleString()}`;
      const toast = new Toast(this.toastEl.nativeElement);
      toast.show();
    }
  }
}
