import { Component, Inject, OnInit, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BlogService } from '@service/blog.service';
import { ArticleService } from '@service/article.service';
import { CopyText } from '@util/util';
import { isPlatformBrowser } from '@angular/common';
import { NzMessageService } from 'ng-zorro-antd/message';

interface TreeItem {
  key: string;
  title: string;
  isLeaf?: boolean;
  children: Array<TreeItem>;
}
@Component({
  selector: 'app-add-question',
  templateUrl: './add-question.component.html',
  styleUrls: ['./add-question.component.less']
})
export class AddQuestionComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private blogHttp: BlogService,
    private httpService: ArticleService,
    @Inject(PLATFORM_ID) private platformId: object,
    private message: NzMessageService
  ) { }
  public detail = {
    title: '',
    CreateAt: '',
    tag: [],
    tag_arr: [],
    answer_md: '',
    level: '',
    classify_id: ''
  };
  tagVal = [];
  status = 1;
  classification: number;
  public markdownContent: string;
  public cacheMarkdownContent: string;
  selectedValue = '';
  public tagList = [];
  public classifyList = [];
  visible = false;
  ngOnInit(): void {
    this.getTag();
    this.getClassify();
    this.route.queryParams.subscribe((params) => {
      if (params.id) {
        this.getInterviewDetail(params.id);
      }
    });
    if (isPlatformBrowser(this.platformId)) {
      CopyText();
    }
  }
  editFunc(id: number): void {
    this.visible = true;
    this.status = id;
    id === 1 ? this.markdownContent = this.detail.title : this.markdownContent = this.detail.answer_md;
  }
  close(): void {
    this.visible = false;
  }
  async getClassify(): Promise<void> {
    const res = await this.blogHttp.getInterview();
    if (res.code === 200) {
      this.classifyList = res.data;
    }
  }
  // 获取tag
  async getTag(): Promise<void> {
    const res = await this.httpService.getTag();
    if (res.code === 200) {
      this.tagList = res.data;
    }
  }
  async getInterviewDetail(id: number): Promise<void> {
    const res = await this.blogHttp.getInterviewDetail({ id });
    if (res.code === 200) {
      this.detail = res.data;
    }
  }


  EditValueOutput(value: string): void {
    // console.log(value, '编辑器');
    this.cacheMarkdownContent = value;
  }
  async saveFunc(): Promise<void> {
    console.log(this.detail);
    const res = await this.blogHttp.addInterview(this.detail);
    if (res.code === 200) {
      this.message.success('保存成功!');
      this.detail = res.data.data;
    }
  }
  sureFunc(): void {
    this.visible = false;
    this.status === 1 ? this.detail.title = this.cacheMarkdownContent : this.detail.answer_md = this.cacheMarkdownContent;
  }

}
