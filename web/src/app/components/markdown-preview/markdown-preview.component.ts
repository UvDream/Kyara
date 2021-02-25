import { Component, OnInit, Input, Inject, PLATFORM_ID, OnChanges, Output, EventEmitter } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import Vditor from 'vditor';
import { NzImageService } from 'ng-zorro-antd/image';


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
    @Inject(PLATFORM_ID) private platformId: object,
    private nzImageService: NzImageService,
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
        // 自定义渲染
        renderers: {
          // 标题自定义渲染
          renderHeading: (node: ILuteNode, entering: boolean) => {
            if (entering) {
              return [`<h${node.__internal_object__.HeadingLevel}><span class="prefix"></span><span class="content"> `, Lute.WalkContinue];
            } else {
              return [`</span><span class="suffix"></span> </h${node.__internal_object__.HeadingLevel} > `, Lute.WalkContinue];
            }
          },
          // renderLinkDest: (node: ILuteNode, entering: boolean) => {
          //   const src = node.TokensStr();
          //   return [`wzj`, Lute.WalkContinue];
          // },
          // renderImage: (node: ILuteNode, entering: boolean) => {
          // console.log('image');
          // const src = node.TokensStr();

          // if (entering) {
          // return [`<p>${src}</p>`, Lute.WalkContinue];
          // } else {
          //   return [` </p>`, Lute.WalkContinue];
          // }
          // }
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
  contentClick(event: any): void {
    if (event.target.nodeName.toString() === 'IMG') {
      console.log(event);
      const images = [
        {
          src: event.target.currentSrc,
          width: event.target.naturalWidth,
          height: event.target.naturalHeight,
          alt: event.target.alt
        }
      ];
      this.nzImageService.preview(images, { nzZoom: 1, nzRotate: 0 });
    }
  }


}
