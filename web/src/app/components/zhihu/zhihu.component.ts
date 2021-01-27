import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';

@Component({
  selector: 'app-zhihu',
  templateUrl: './zhihu.component.html',
  styleUrls: ['./zhihu.component.less']
})
export class ZhihuComponent implements OnInit {
  @ViewChild('canvas', { static: true })
  canvas: ElementRef<HTMLCanvasElement>;

  public context: CanvasRenderingContext2D;
  constructor() { }

  ngOnInit(): void {
    this.context = this.canvas.nativeElement.getContext('2d');
    this.context.fillStyle = '#FF0000';
    this.context.fillRect(0, 0, 150, 75);
  }

}
