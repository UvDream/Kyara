import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
})
export class HomeComponent implements OnInit {
  public list: Array<any>;
  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.getData();
  }
  // tslint:disable-next-line:typedef
  getData() {
    this.http
      .post('http://127.0.0.1:3000/article/articleList', {
        search: '',
        page: 1,
        pageSize: 5,
      })
      .subscribe((res: any) => {
        console.log(res);
        this.list = res.data.msg;
      });
  }
}
