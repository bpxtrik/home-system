import { Component, ElementRef, inject, OnInit, signal, ViewChild } from '@angular/core';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';
import { Motion } from '../shared/dto';
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
  data = signal<Motion[]>([]);
  wsService = inject(WsService);
  @ViewChild('toast') toastEl!: ElementRef;

  ngOnInit(): void {
    this.wsService.connect();

    this.wsService.messages.subscribe((_) => {
      console.log('WS');
      this.showToast();
      this.fetchMotions();
    });
    this.fetchMotions();
  }

  fetchMotions(): void {
    this.dashboardService.getMotions().subscribe({
      next: (res) => {
        this.data.set(res);
      },
      error: (err) => {
        console.log(err);
      },
    });
  }

  showToast() {
    const toast = new Toast(this.toastEl.nativeElement);
    toast.show();
  }
}
