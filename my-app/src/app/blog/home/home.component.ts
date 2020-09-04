import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.less'],
})
export class HomeComponent implements OnInit {
  public list: Array<any>;
  public form = {
    search: '',
    page: 1,
    pageSize: 10,
  };
  public totalCount = 0;
  public isVisible = false;
  public passwordVal: string;
  private articleId: string;
  passwordVisible = false;
  constructor(private http: HttpClient, private titleService: Title, private router: Router, private message: NzMessageService) { }
  ngOnInit(): void {
    this.getData();
    this.setTitle('首页');
  }
  getData = () => {
    this.http
      .post('http://127.0.0.1:3000/article/articleList', this.form)
      .subscribe((res: any) => {
        this.list = res.data.msg;
        this.totalCount = res.data.totalCount;
      });
  }
  // tslint:disable-next-line:typedef
  public setTitle(newTitle: string) {
    this.titleService.setTitle(newTitle);
  }
  pageChange = () => {
    this.getData();
  }
  havePassword = (id: string) => {
    this.isVisible = true;
    this.articleId = id;
  }
  modalFunc = (data: boolean) => {
    this.isVisible = data;
  }
  sureModal = () => {
    const url = 'http://127.0.0.1:3000/article/articlePassword?id=' + this.articleId + '&password=' + this.passwordVal;
    this.http.get(url).subscribe((res: any) => {
      if (res.code === 200) {
        this.message.success(res.msg, {
          nzDuration: 1000
        });
        this.isVisible = false;
        this.router.navigate(['/detail'], { queryParams: { id: this.articleId, password: this.passwordVal } });
      } else {
        this.message.error(res.msg, {
          nzDuration: 1000
        });
      }
    });
  }
}

