import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';
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

  constructor(private route: ActivatedRoute, private http: HttpClient) { }
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
    const id = +this.route.snapshot.paramMap.get('id');
    this.getDetail(id);
  }
  getDetail = (id) => {
    const url = 'http://127.0.0.1:3000/article/articleDetail?id=' + id;
    this.http.get(url).subscribe((res: any) => {
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
    });
  }
  showModal = (data: boolean) => {
    this.isVisible = data;
  }
  tabClick = (data: TabItem) => {
    this.activeId = data.id;
  }
}
