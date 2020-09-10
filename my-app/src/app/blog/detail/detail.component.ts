import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ArticleService } from '../../service/article.service';
import { ArticleCatalogService } from '../../service/article-catalog.service';
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
  ) { }
  public htmlContent: any;
  public markDown: any;
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
    this.route.queryParams.subscribe(params => {
      id = params.id;
      password = params.password;
      if (password === undefined) { password = ''; }
    });
    this.getDetail(id, password);
  }
  getDetail = async (id, password) => {
    const res = await this.request.getDetail({ id, password });
    if (res.code !== 200) { return; }
    this.htmlContent = res.data.article_html;
    this.markDown = res.data.article_content;
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
          this.parseDom(html).forEach(element => {
            if (element.nodeName === 'H1' || element.nodeName === 'H2' || element.nodeName === 'H3') {
              const obj = {
                // tslint:disable-next-line:no-string-literal
                id: element['id'],
                // tslint:disable-next-line:no-string-literal
                title: element['innerText'],
                children: []
              };
              if (element.nodeName === 'H1') {
                arr.push(obj);
              }
              if (element.nodeName === 'H2') {
                arr[arr.length - 1].children.push(obj);
              }
              if (element.nodeName === 'H3') {
                arr[arr.length - 1].children[arr[arr.length - 1].children.length - 1].children.push(obj);
              }
            }
          });
          this.articleCat.SetCatLog(arr);

          return html;
        }
      })
    );
  }
  parseDom = (arg: any) => {
    const objE = document.createElement('div');
    objE.innerHTML = arg;
    return objE.childNodes;
  }
  showModal = (data: boolean) => {
    this.isVisible = data;
  }
  tabClick = (data: TabItem) => {
    this.activeId = data.id;
  }
}
