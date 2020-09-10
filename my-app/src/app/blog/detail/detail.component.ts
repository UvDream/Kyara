import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ArticleService } from '../../service/article.service';
import { ArticleCatalogService } from '../../service/article-catalog.service';
import { Title } from '@angular/platform-browser';

interface TabItem {
  id: number;
  title: string;
  imgUrl: string;
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
    private titleService: Title
  ) {}
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
  @ViewChild('vditor') myDom: HTMLDivElement;
  isVisible = false;
  public tabList = [
    {
      id: 1,
      title: '支付宝支付',
      imgUrl:
        'https://gitee.com/Uvdream/images/raw/master/images/20200812100455.png',
    },
    {
      id: 2,
      title: '微信支付',
      imgUrl:
        'https://gitee.com/Uvdream/images/raw/master/images/20200812100455.png',
    },
  ];
  public imgUrl = this.tabList[0].imgUrl;
  public activeId = this.tabList[0].id;
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
  }
  // 获取文章详情
  getDetail = async (id, password) => {
    const res = await this.request.getDetail({ id, password });
    if (res.code !== 200) {
      return;
    }
    this.htmlContent = res.data.article_html;
    this.markDown = res.data.article_content;
    console.log(res.data);
    this.title = res.data.title;
    this.author = res.data.user_name;
    this.time = res.data.UpdatedAt;
    this.wordCount = res.data.word_count;
    this.viewCount = res.data.view_count;
    this.commentCount = res.data.comment_count;
    this.topImg = res.data.img_url;
    this.setTitle(this.title + '(汪中杰的个人博客)');
    const mainElement = document.getElementById('vditor') as HTMLDivElement;
    import('vditor').then((Vditor: any) =>
      Vditor.preview(mainElement, this.markDown, {
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
                children: [],
              };
              if (element.nodeName === 'H1') {
                arr.push(obj);
              }
              if (element.nodeName === 'H2') {
                arr[arr.length - 1].children.push(obj);
              }
              if (element.nodeName === 'H3') {
                arr[arr.length - 1].children[
                  arr[arr.length - 1].children.length - 1
                ].children.push(obj);
              }
            }
          });
          this.articleCat.SetCatLog(arr);

          return html;
        },
      })
    );
  };
  parseDom = (arg: any) => {
    const objE = document.createElement('div');
    objE.innerHTML = arg;
    return objE.childNodes;
  };
  showModal = (data: boolean) => {
    this.isVisible = data;
  };
  tabClick = (data: TabItem) => {
    this.activeId = data.id;
  };
  public setTitle = (newTitle: string) => {
    this.titleService.setTitle(newTitle);
  };
}
