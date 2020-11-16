import { Component, OnInit } from '@angular/core';
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
  constructor() { }

  ngOnInit(): void {
  }
  btnClick(): void {
    this.visible = !this.visible;
  }
  private scrollTo(to: number, duration: number): void {

  }
}
