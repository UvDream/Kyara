import { Component, OnInit } from '@angular/core';
import { ArticleService } from '../../service/article.service';
interface TreeItem {
  key: string;
  title: string;
  isLeaf?: boolean;
  children: Array<TreeItem>;
}
@Component({
  selector: 'app-edit-article',
  templateUrl: './edit-article.component.html',
  styleUrls: ['./edit-article.component.less'],
})
export class EditArticleComponent implements OnInit {
  constructor(private httpService: ArticleService) { }

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
  expandKeys = ['100', '1001'];
  public nodes = [
  ];

  ngOnInit(): void {
    this.getClassify();
  }
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
  // 获取分类
  getClassify = async () => {
    const res = await this.httpService.getArticleClassification();
    if (res.code !== 200) { return; }
    console.log('老数据', res.data);
    this.nodes = this.changeData(res.data);
    console.log('新数据', this.nodes);
  }
  changeData = (arr: any): Array<TreeItem> => {
    const data: Array<TreeItem> = [];
    arr.forEach(element => {
      const obj: TreeItem = {
        title: '',
        key: '',
        children: []
      };
      obj.title = element.type_name;
      obj.key = element.ID;
      element.children.length > 0 ? obj.children = this.changeData(element.children) : obj.isLeaf = true;
      data.push(obj);
    });
    return data;
  }
}
