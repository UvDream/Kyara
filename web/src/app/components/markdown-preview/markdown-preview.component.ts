import { Component, OnInit, Input, Inject, PLATFORM_ID, OnChanges, Output, EventEmitter } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import Vditor from 'vditor';

@Component({
  selector: 'app-markdown-preview',
  templateUrl: './markdown-preview.component.html',
  styleUrls: ['./markdown-preview.component.less']
})
export class MarkdownPreviewComponent implements OnInit, OnChanges {
  vditor: Vditor;

  @Input() Markdown: string;
  @Output() CatalogOut = new EventEmitter();
  public mainElement;
  constructor(
    @Inject(PLATFORM_ID) private platformId: object
  ) { }

  ngOnInit(): void {
    this.mainElement = document.getElementById('markdown') as HTMLDivElement;
  }
  ngOnChanges(): void {
    this.markdownToHtml(this.Markdown);
  }
  markdownToHtml(markdown: string): void {
    if (isPlatformBrowser(this.platformId) && markdown !== undefined) {
      Vditor.preview(this.mainElement, markdown, {
        mode: 'dark',
        speech: {
          enable: true,
        },
        anchor: 0,
        hljs: {
          enable: true,
          lineNumber: true,
          style: 'native',
        },
        markdown: {
          toc: true,
        },
        transform: (html: string) => {
          const arr = [];
          this.parseDom(html).forEach((element) => {
            if (
              element.nodeName === 'H1' ||
              element.nodeName === 'H2' ||
              element.nodeName === 'H3'
            ) {
              const obj = {
                // tslint:disable-next-line:no-string-literal
                id: element['id'],
                // tslint:disable-next-line:no-string-literal
                title: element['innerText'],
                name: element.nodeName,
                children: [],
              };
              arr.push(obj);
            }
          });
          this.CatalogOut.emit(arr);
          return html;
        }
      });
    }
  }
  parseDom = (arg: any) => {
    if (isPlatformBrowser(this.platformId)) {
      const objE = document.createElement('div');
      objE.innerHTML = arg;
      return objE.childNodes;
    }
  }


}
