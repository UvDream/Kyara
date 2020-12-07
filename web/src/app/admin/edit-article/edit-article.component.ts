import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute } from '@angular/router';

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
  constructor(
    private route: ActivatedRoute,
    private httpService: ArticleService,
    private message: NzMessageService,
  ) { }
  isPassword = false;
  isTop = false;
  coverTypeDisabled = false;
  value?: string;
  // 是否置顶
  public form = {
    article_id: 0,
    // 标题
    title: '',
    // 摘要
    introduction: '',
    // 封面类型
    cover_type: '0',
    // 文章类容
    article_content: '',
    // 置顶
    top: '0',
    // 访问密码
    is_password: '0',
    view_password: '',
    // 是否开启评论
    is_comment: true,
    // 转载规则
    transfer_rules: '0',
    // icon
    icon: '',
    // 图片地址
    img_url: '',
    // tag
    tag_array: [],
  };
  tagVal = '';
  expandKeys = ['100', '1001'];
  tagList = [];
  public nodes = [
  ];

  ngOnInit(): void {
    this.getClassify();
    this.getTag();
    this.route.queryParams.subscribe((params) => {
      console.log(params);
      this.getArticleDetail(params.id);
    });
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
  // 提交
  submitFunc = async () => {
    console.log('提交', this.form);
    this.isPassword ? this.form.is_password = '1' : this.form.is_password = '0';
    this.isTop ? this.form.top = '1' : this.form.top = '0';
    const res = await this.httpService.addArticle(this.form);
    console.log(res);
    if (res.code === 200) {
      this.form.article_id = res.data.ID;
      this.message.info(res.msg);
    } else {
      this.message.error(res.msg);
    }
  }
  checkChange = () => {
    if (this.isTop) {
      this.form.cover_type = '2';
      this.coverTypeDisabled = true;
    } else {
      this.coverTypeDisabled = false;
    }
  }
  async getArticleDetail(id: string): Promise<void> {
    const res = await this.httpService.getArticleDetail(id);
    console.log(res);
    if (res.code === 200) {
      this.form = res.data;
      this.form.article_id = res.data.ID;
      this.markdownToHtml(this.form.article_content);
    }
  }
  // 获取tag
  async getTag(): Promise<void> {
    const res = await this.httpService.getTag();
    if (res.code === 200) {
      this.tagList = res.data;
    }
  }
  // 增加tag
  async addTag(): Promise<void> {
    const res = await this.httpService.addTag({ tag: this.tagVal });
    if (res.code === 200) {
      this.getTag();
      this.message.success(res.msg);
    } else {
      this.message.error(res.msg);
    }
  }
}
