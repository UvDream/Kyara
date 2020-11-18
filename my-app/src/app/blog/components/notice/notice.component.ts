import { Component, OnInit, OnDestroy } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
// import { Platform } from '@angular/cdk/platform';
@Component({
  selector: 'app-notice',
  templateUrl: './notice.component.html',
  styleUrls: ['./notice.component.less']
})
export class NoticeComponent implements OnInit, OnDestroy {
  noticeShow = false;
  time: any;
  constructor(
    public config: BlogConfigService,
    // private platform: Platform,
  ) {

  }

  ngOnInit(): void {
    this.noticeShow = true;
    setTimeout(() => {
      if (this.config.BlogNotice === '') {
        this.noticeShow = false;
      }
      this.scrollLeft();
      this.time = setInterval(() => {
        this.scrollFunc();
      }, 10);
    }, 1000);

  }
  ngOnDestroy(): void {
    clearInterval(this.time);
  }
  closeNotice(): void {
    this.noticeShow = false;
  }
  scrollLeft(): void {
    const textMain = document.getElementById('text-main') as HTMLDivElement;
    const text = document.getElementById('text') as HTMLDivElement;
    text.style.width = textMain.offsetWidth + 'px';
    const newNode = document.createElement('div');
    newNode.id = 'textCopy';
    newNode.style.width = textMain.offsetWidth + 'px';
    newNode.style.display = 'inline-block';
    newNode.innerHTML = text.innerHTML;
    textMain.appendChild(newNode);
  }
  scrollFunc(): void {
    const text = document.getElementById('text') as HTMLDivElement;
    const textMain = document.getElementById('text-main') as HTMLDivElement;
    if (text.offsetWidth - textMain.scrollLeft <= 4) {
      textMain.scrollLeft = 0;
    } else {
      textMain.scrollLeft++;
    }
  }
}
