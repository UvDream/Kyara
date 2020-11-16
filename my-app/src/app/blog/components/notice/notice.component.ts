import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-notice',
  templateUrl: './notice.component.html',
  styleUrls: ['./notice.component.less']
})
export class NoticeComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    this.scrollLeft();
    setInterval(() => {
      this.scrollFunc();
    }, 10);

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
    if (text.offsetWidth - textMain.scrollLeft <= 0) {
      textMain.scrollLeft = 0;
    } else {
      textMain.scrollLeft++;
    }
  }
}
