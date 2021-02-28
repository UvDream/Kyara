import { Component, OnInit, Output, EventEmitter, Input, Inject, PLATFORM_ID, OnChanges } from '@angular/core';
import Vditor from 'vditor';
import { isPlatformBrowser } from '@angular/common';
import { Emoji } from './emoji';
@Component({
  selector: 'app-markdown-edit',
  templateUrl: './markdown-edit.component.html',
  styleUrls: ['./markdown-edit.component.less']
})
export class MarkdownEditComponent implements OnInit, OnChanges {
  vditor: Vditor;
  @Input() height = 600;
  // 编辑器内容回显
  @Input() EditValue: string;
  @Input() EditToolbar = [
    'emoji',
    'headings',
    'bold',
    'italic',
    'strike',
    'link',
    '|',
    'list',
    'ordered-list',
    'check',
    'outdent',
    'indent',
    '|',
    'quote',
    'line',
    'code',
    'inline-code',
    'insert-before',
    'insert-after',
    '|',
    'upload',
    // 'record',
    'table',
    '|',
    'undo',
    'redo',
    '|',
    'fullscreen',
    'edit-mode'];
  // 编辑器内容回调
  @Output() EditValueOutput = new EventEmitter();
  // 字数回调
  @Output() EditCounterOutput = new EventEmitter();
  public editContent: string;
  constructor(
    @Inject(PLATFORM_ID) private platformId: object
  ) { }
  ngOnChanges(): void {
    this.vditor.setValue(this.EditValue);
  }
  ngOnInit(): void {
    console.log('编辑器');
    if (isPlatformBrowser(this.platformId)) {
      this.vditor = new Vditor('vditor', {
        // value: this.EditValue,
        height: this.height,
        toolbar: this.EditToolbar,
        toolbarConfig: {
          pin: true,
        },
        hint: {
          emoji: Emoji
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
          enable: true,
        },
        blur: (value: string) => {
          // console.log('编辑器的值');
          // console.log(value);
          this.EditValueOutput.emit(value);
        },
        after: () => {
          if (this.EditValue !== '') {
            this.vditor.setValue(this.EditValue);
          }
        }
      });
    }
  }

}
