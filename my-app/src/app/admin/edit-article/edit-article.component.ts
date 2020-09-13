import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-edit-article',
  templateUrl: './edit-article.component.html',
  styleUrls: ['./edit-article.component.less'],
})
export class EditArticleComponent implements OnInit {
  markdownContent?: string;
  value?: string;
  constructor() { }
  ngOnInit(): void { }
  textChange = () => {
    console.log('改变');
    console.log(this.markdownContent);
    this.markdownToHtml(this.markdownContent);
  }
  // 渲染markdown
  markdownToHtml = (data) => {
    const mainElement = document.getElementById('vditor') as HTMLDivElement;

    import('vditor').then((Vditor: any) =>
      Vditor.preview(mainElement, data, {
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
      })
    );
  }
}
