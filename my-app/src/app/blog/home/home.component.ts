import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ArticleService } from '../../service/article.service'
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
  // tslint:disable-next-line:max-line-length
  constructor(
    private http: HttpClient,
    private titleService: Title,
    private router: Router,
    private message: NzMessageService,
    private articleService: ArticleService
  ) { }
  ngOnInit(): void {
    this.getData();
    this.setTitle('首页');
  }
  getData = async () => {
    const res = await this.articleService.articleList(this.form);
    if (res.code === 200) {
      this.list = res.data.msg;
      this.totalCount = res.data.totalCount;
    }

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
  sureModal = async () => {
    const res = await this.articleService.surePassword({ id: this.articleId, password: this.passwordVal });
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
  }
}

