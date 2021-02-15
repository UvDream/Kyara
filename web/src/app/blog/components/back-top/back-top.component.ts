import { Component, OnInit, HostListener, Inject, PLATFORM_ID } from '@angular/core';
import { fadeIn } from '../../animations/fade-in';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-back-top',
  exportAs: 'BackTop',
  templateUrl: './back-top.component.html',
  styleUrls: ['./back-top.component.less'],
  animations: [fadeIn]
})
export class BackTopComponent implements OnInit {
  visible = false;
  constructor(
    @Inject(PLATFORM_ID) private platformId: object
  ) { }
  @HostListener('window:scroll', [])
  onScroll(): void {
    if (isPlatformBrowser(this.platformId)) {
      document.documentElement.scrollTop !== 0 ? this.visible = true : this.visible = false;
    }
  }
  ngOnInit(): void {
  }
  btnClick(): void {
    this.visible = !this.visible;
  }
  backTop(): void {
    if (isPlatformBrowser(this.platformId)) {
      let speed = 0;
      let osTop = 0;
      const timer = setInterval(() => {
        osTop = document.documentElement.scrollTop || document.body.scrollTop;
        speed = Math.ceil(-osTop / 2);
        document.body.scrollTop = document.documentElement.scrollTop -= (osTop + speed);
        if (speed === 0) {
          clearInterval(timer);
        }
      }, 100);
    }
  }

  private scrollTo(to: number, duration: number): void {

  }
}
