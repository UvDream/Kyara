import { Component, Input, OnInit, Output, EventEmitter, Inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
@Component({
  selector: 'app-comment-component',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.less']
})
export class CommentComponent implements OnInit {
  @Input() data: Array<any>;
  @Input() isChildren = false;
  @Output() ActiveIDOut = new EventEmitter();
  constructor(
    @Inject(PLATFORM_ID) private platformId: object,
  ) { }
  public activeId: number;

  ngOnInit(): void {
    // console.log(this.data, this.isChildren);
  }

  replyFunc(item: any): void {
    this.activeId = item.ID;
    this.ActiveIDOut.emit(item);
    if (item.ID !== 0 && isPlatformBrowser(this.platformId)) {
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

}
