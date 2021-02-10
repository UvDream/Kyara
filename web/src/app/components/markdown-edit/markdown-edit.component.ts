import { Component, OnInit, Output, EventEmitter, Input, Inject, PLATFORM_ID } from '@angular/core';
import Vditor from 'vditor';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-markdown-edit',
  templateUrl: './markdown-edit.component.html',
  styleUrls: ['./markdown-edit.component.less']
})
export class MarkdownEditComponent implements OnInit {
  vditor: Vditor;
  // 编辑器内容回显
  @Input() EditValue: string;
  // 编辑器内容回调
  @Output() EditValueOutput = new EventEmitter();
  // 字数回调
  @Output() EditCounterOutput = new EventEmitter();
  constructor(
    @Inject(PLATFORM_ID) private platformId: object
  ) { }

  ngOnInit(): void {
    if (isPlatformBrowser(this.platformId)) {
      this.vditor = new Vditor('vditor', {
        height: 800,
        toolbarConfig: {
          pin: true,
        },
        counter: {
          enable: true,
          after: (length: number, counter: any) => {
            // console.log('统计字数');
            // console.log(length, counter);
            this.EditCounterOutput.emit(length);
          },
        },
        cache: {
          enable: false,
        },
        blur: (value: string) => {
          // console.log('编辑器的值');
          // console.log(value);
          this.EditValueOutput.emit(value);
        },
        after: () => {
          this.vditor.setValue(this.EditValue);
        }
      });
    }
  }

}
