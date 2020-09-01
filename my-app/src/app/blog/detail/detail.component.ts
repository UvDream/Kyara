import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.less'],
})
export class DetailComponent implements OnInit {
  public htmlContent: any;
  public markDown: any;
  // public vditor: Vditor;
  @ViewChild('vditor') myDom: HTMLDivElement;
  constructor(private route: ActivatedRoute, private http: HttpClient) {}

  ngOnInit(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    this.getDetail(id);
  }
  // tslint:disable-next-line:typedef
  getDetail(id) {
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
}
