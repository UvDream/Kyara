import { Component, EventEmitter, Input, OnInit, Output, DoCheck } from '@angular/core';

@Component({
  selector: 'app-image-preview',
  templateUrl: './image-preview.component.html',
  styleUrls: ['./image-preview.component.less']
})
export class ImagePreviewComponent implements OnInit, DoCheck {
  @Input() visible = false;
  @Input() src: string;
  @Output() readonly visibleChange = new EventEmitter<boolean>();
  constructor() { }
  imgWidth = '100%';
  imgHeight = '100%';
  ngOnInit(): void {

  }
  ngDoCheck(): void {
    if (this.src) {
      const img = new Image();
      img.src = this.src;
      if (img.complete) {
        console.log(img.width, '图片高度');
        console.log(document.body.clientWidth / img.width, document.body.clientHeight / img.height);
        const ratio = document.body.clientWidth / img.width;
        if (ratio > 5) {
          this.imgWidth = img.width * 2 + 'px';
        }
        if (ratio < 1) {
          this.imgWidth = img.width / 2 + 'px';
        }
        if (ratio < 2 && ratio >= 1) {
          this.imgWidth = img.width / 1.5 + 'px';
        }
      }
    }
  }
  imgClick(): void {
    this.visible = false;
    this.visibleChange.emit(false);
  }

}
