import { Component, OnInit, ViewChild, Inject, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ArticleService } from '@service/article.service';
import { ArticleCatalogService } from '@service/article-catalog.service';
import { Title } from '@angular/platform-browser';
import { Platform } from '@angular/cdk/platform';
import { isPlatformBrowser } from '@angular/common';

interface TabItem {
  ID: number;
  name: string;
  img_url: string;
}
@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.less'],
})
export class DetailComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private request: ArticleService,
    private articleCat: ArticleCatalogService,
    private titleService: Title,
    private platform: Platform,
    @Inject(PLATFORM_ID) private platformId: object
  ) { }
  public loading = true;
  // 文章标题
  public title: string;
  // html内容
  public htmlContent: any;
  // markdown内容
  public markDown: any;
  // 作者
  public author: string;
  // 时间
  public time: string;
  // 文章字数
  public wordCount: string;
  // 浏览次数
  public viewCount: string;
  // 评论数
  public commentCount: string;
  // 图片
  public topImg: string;
  // 更新时间
  public UpdatedAt: string;
  // 赞赏提示文字
  public collectText: string;
  @ViewChild('vditor') myDom: HTMLDivElement;
  // 赞赏弹窗
  isVisible = false;
  public tabList = [];
  public imgUrl = '';
  public activeId = 1;
  ngOnInit(): void {
    let id = '';
    let password = '';
    this.route.queryParams.subscribe((params) => {
      id = params.id;
      password = params.password;
      if (password === undefined) {
        password = '';
      }
    });
    this.getDetail(id, password);
    this.viewArticle(id);
    const url = '/detail?id=' + id;
    this.request.toBaidu(url).then(res => {
      if (res.code === 200) {
        console.log(res.msg);
      }
    });
  }
  catalogOut(arr: any): void {
    this.articleCat.SetCatLog(arr);
  }
  // 获取文章详情
  getDetail = async (id, password) => {
    this.loading = true;
    const res = await this.request.getDetail({ id, password });
    if (res.code !== 200) {
      return;
    }
    this.htmlContent = res.data.article_html;
    this.markDown = res.data.article_content;
    // console.log(res.data);
    this.title = res.data.title;
    this.author = res.data.user_name;
    this.time = res.data.UpdatedAt;
    this.wordCount = res.data.word_count;
    this.viewCount = res.data.view_count;
    this.commentCount = res.data.comment_count;
    this.topImg = res.data.img_url;
    this.UpdatedAt = res.data.UpdatedAt;
    this.tabList = res.data.collect_list;
    if (res.data.collect_list.length) {
      this.imgUrl = this.tabList[0].img_url;
      this.activeId = this.tabList[0].ID;
    }

    this.collectText = res.data.collect_text;
    this.setTitle(this.title + '(汪中杰的个人博客)');
    this.loading = false;
  }
  parseDom = (arg: any) => {
    if (isPlatformBrowser(this.platformId)) {
      const objE = document.createElement('div');
      objE.innerHTML = arg;
      return objE.childNodes;
    }
  }
  showModal = (data: boolean) => {
    this.isVisible = data;
  }
  tabClick = (data: TabItem) => {
    this.activeId = data.ID;
    this.imgUrl = data.img_url;
  }
  public setTitle = (newTitle: string) => {
    this.titleService.setTitle(newTitle);
  }
  async viewArticle(id: string): Promise<void> {
    console.log('文章访问');
    const res = await this.request.viewBlog(id);
  }

  print(): void {
    if (this.platform.isBrowser) {
      setTimeout(() => {
        window.print();
      }, 500);
    }

  }
}
