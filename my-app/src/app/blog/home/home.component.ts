import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.less'],
})
export class HomeComponent implements OnInit {
  public list: Array<any>;
  constructor(private http: HttpClient, private titleService: Title) {}
  ngOnInit(): void {
    this.getData();
    this.setTitle('首页');
  }
  getData = () => {
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
  };
  // tslint:disable-next-line:typedef
  public setTitle(newTitle: string) {
    this.titleService.setTitle(newTitle);
  }
}
