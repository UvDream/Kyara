import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-edit-article',
  templateUrl: './edit-article.component.html',
  styleUrls: ['./edit-article.component.less'],
})
export class EditArticleComponent implements OnInit {
  markdownContent?: string;
  value?: string;
  // 是否置顶
  public form = {
    // 置顶
    top: false,
    // 访问密码
    is_password: false,
    view_password: '',
    // 是否开启评论
    is_comment: true,
    // 转载规则
    transfer_rules: 0
  };
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
