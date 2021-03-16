import { Component, Input, OnInit, Output, EventEmitter, Inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { CopyText } from '@util/util';
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
    if (isPlatformBrowser(this.platformId)) {
      CopyText();
    }
  }

  replyFunc(item: any): void {
    this.activeId = item.ID;
    this.ActiveIDOut.emit(item);
    if (item.ID !== 0 && isPlatformBrowser(this.platformId)) {
      const timer = setInterval(() => {
        const osTop = document.getElementById('replayForm').offsetTop;
        document.documentElement.scrollTop -= 40;
        if (document.documentElement.scrollTop < osTop) {
          clearInterval(timer);
        }
      }, 100);
    }
  }

}
