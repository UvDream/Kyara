import { Component, Inject, OnInit, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BlogService } from '@service/blog.service';
import { ArticleService } from '@service/article.service';
import { CopyText } from '@util/util';
import { isPlatformBrowser } from '@angular/common';


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
  classification: number;
  public nodes = [
  ];
  visible = false;
  ngOnInit(): void {
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
  close(): void {
    this.visible = false;
  }
  async getInterviewDetail(id: number): Promise<void> {
    const res = await this.blogHttp.getInterviewDetail({ id });
    if (res.code === 200) {
      this.detail = res.data;
    }
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
  // 获取分类
  getClassify = async () => {
    const res = await this.httpService.getArticleClassification();
    if (res.code !== 200) { return; }
    this.nodes = this.changeData(res.data);
  }

}
