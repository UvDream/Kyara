import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.scss'],
})
export class DetailComponent implements OnInit {
  public htmlContent: any;
  constructor(private route: ActivatedRoute, private http: HttpClient) {}

  ngOnInit(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    console.log(id);
    this.getDetail(id);
  }
  // tslint:disable-next-line:typedef
  getDetail(id) {
    const url = 'http://127.0.0.1:3000/article/articleDetail?id=' + id;
    this.http.get(url).subscribe((res: any) => {
      this.htmlContent = res.data.article_html;
    });
  }
}
