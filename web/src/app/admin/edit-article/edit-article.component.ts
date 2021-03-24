import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute, Router } from '@angular/router';
import { AdminService } from '@service/admin.service';
import { NzUploadChangeParam } from 'ng-zorro-antd/upload';
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
    private router: Router,
    private httpService: ArticleService,
    private message: NzMessageService,
    private httpAdmin: AdminService,
  ) { }
  isPassword = false;
  // 配置抽屉
  drawerVisible = false;
  isTop = false;
  coverTypeDisabled = false;
  value?: string;
  classification: number;
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
    // 分类id
    classification_id: '',
    // 文章字数
    word_count: 0
  };
  tagVal = '';
  expandKeys = ['100', '1001'];
  tagList = [];
  routerID = '';
  public nodes = [
  ];

  ngOnInit(): void {
    this.getClassify();
    this.getTag();
    this.route.queryParams.subscribe((params) => {
      if (params.id) {
        this.getArticleDetail(params.id);
        this.routerID = params.id;
      }
    });
  }
  dataURItoBlob = (dataURI: any) => {
    const mimeString = dataURI.split(',')[0].split(':')[1].split(';')[0]; // mime类型
    const byteString = atob(dataURI.split(',')[1]);
    const arrayBuffer = new ArrayBuffer(byteString.length);
    const intArray = new Uint8Array(arrayBuffer);
    for (let i = 0; i < byteString.length; i++) {
      intArray[i] = byteString.charCodeAt(i);
    }
    return new Blob([intArray], { type: mimeString });
  }
  async uploadImg(data: any): Promise<void> {
    const formData = new FormData();
    formData.append('type', 'article');
    formData.append('file', data);
    const res = await this.httpAdmin.uploadImage(formData);
    if (res.code === 200) {
      console.log(res, '上传成功');
      const msg = '![img](' + res.data.url + ')';
      this.form.article_content = this.form.article_content + msg;
    }
  }
  // 获取分类
  getClassify = async () => {
    const res = await this.httpService.getArticleClassification();
    if (res.code !== 200) { return; }
    this.nodes = this.changeData(res.data);
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
    this.form.classification_id = this.classification.toString();
    const res = await this.httpService.addArticle(this.form);
    if (res.code === 200) {
      this.form.article_id = res.data.ID;
      this.message.info(res.msg);
      if (this.routerID === '') {
        this.router.navigate(['/admin/editArticle'], { queryParams: { id: res.data.ID } });
      }
    } else {
      this.message.error(res.msg);
    }
  }
  // 解析markdown文件

  onFileChange(event: any): void {
    console.log(event);
    const input = event.target;
    const reader = new FileReader();
    reader.onload = () => {
      this.form.article_content = reader.result.toString();
    };
    reader.readAsText(input.files[0]);
  }
  checkChange = () => {
    if (this.isTop) {
      this.form.cover_type = '2';
      this.coverTypeDisabled = true;
    } else {
      this.coverTypeDisabled = false;
    }
  }
  // 获取详情
  async getArticleDetail(id: string): Promise<void> {
    const res = await this.httpService.getArticleDetail(id);
    console.log(res);
    if (res.code === 200) {
      this.form = res.data;
      this.form.article_id = res.data.ID;
      this.classification = Number(this.form.classification_id);
      res.data.top === '1' ? this.isTop = true : this.isTop = false;
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
      this.tagVal = '';
      this.message.success(res.msg);
    } else {
      this.message.error(res.msg);
    }
  }
  urlChange(data: string): void {
    this.form.img_url = data;
  }
  // 配置抽屉
  drawerOpen(): void {
    this.drawerVisible = true;
  }

  drawerClose(): void {
    this.drawerVisible = false;
  }
  // 新版编辑器获取内容
  editOutput(value: string): void {
    // console.log('获取编辑器的值', value);
    this.form.article_content = value;
  }
  // 新版编辑器你获取字数
  editCounter(value: number): void {
    // console.log('编辑器字数', value);
    this.form.word_count = value;
  }
}
