import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-edit-article',
  templateUrl: './edit-article.component.html',
  styleUrls: ['./edit-article.component.less'],
})
export class EditArticleComponent implements OnInit {
  value?: string;
  // 是否置顶
  public form = {
    // 标题
    title: '',
    // 摘要
    introduction: '',
    // 封面类型
    cover_type: '0',
    // 文章类容
    article_content: '',
    // 置顶
    top: false,
    // 访问密码
    is_password: false,
    view_password: '',
    // 是否开启评论
    is_comment: true,
    // 转载规则
    transfer_rules: '0',
    // icon
    icon: ''
  };
  constructor() { }
  ngOnInit(): void { }
  textChange = () => {
    this.markdownToHtml(this.form.article_content);
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
