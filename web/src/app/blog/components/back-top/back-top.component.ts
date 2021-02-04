import { Component, OnInit, HostListener, ElementRef } from '@angular/core';
import { fadeIn } from '../../animations/fade-in';
@Component({
  selector: 'app-back-top',
  exportAs: 'BackTop',
  templateUrl: './back-top.component.html',
  styleUrls: ['./back-top.component.less'],
  animations: [fadeIn]
})
export class BackTopComponent implements OnInit {
  visible = false;
  constructor(private el: ElementRef) { }
  @HostListener('window:scroll', [])
  onScroll(): void {
    document.documentElement.scrollTop !== 0 ? this.visible = true : this.visible = false;

  }
  ngOnInit(): void {
  }
  btnClick(): void {
    this.visible = !this.visible;
  }
  backTop(): void {
    console.log(document.body.scrollTop, document.documentElement.scrollTop);
    // document.body.scrollTop = document.documentElement.scrollTop = 0;
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

  private scrollTo(to: number, duration: number): void {

  }
}
