import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { ArticleService } from '../../service/article.service';
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

  constructor(private route: ActivatedRoute, private http: HttpClient, private request: ArticleService) { }
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
      })
    );
  }
  showModal = (data: boolean) => {
    this.isVisible = data;
  }
  tabClick = (data: TabItem) => {
    this.activeId = data.id;
  }
}
