import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ArticleService } from '@service/article.service';
import { ArticleCatalogService } from '@service/article-catalog.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.less'],
})
export class HomeComponent implements OnInit {
  Loading = false;
  public list: Array<any>;
  public form = {
    search: '',
    page: 2,
    pageSize: 10,
  };
  public totalCount = 0;
  public isVisible = false;
  public passwordVal: string;
  private articleId: string;
  passwordVisible = false;
  constructor(
    private catalog: ArticleCatalogService,
    private titleService: Title,
    private router: Router,
    private message: NzMessageService,
    private articleService: ArticleService
  ) { }
  ngOnInit(): void {
    this.getData();
    this.setTitle('汪中杰的个人博客-首页');
    this.catalog.SetCatLog([]);
    this.articleService.toBaidu('/home').then(res => {
      if (res.code === 200) {
        console.log(res.msg);
      }
    });
  }
  getData = async () => {
    this.Loading = true;
    const res = await this.articleService.articleList(this.form);
    if (res.code === 200) {
      this.Loading = false;
      this.list = res.data.msg;
      this.totalCount = res.data.totalCount;
    }

  }
  public setTitle = (newTitle: string) => {
    this.titleService.setTitle(newTitle);
  }
  pageChange = (num: number) => {
    console.log('页码变化', num, this.form.page);
    sessionStorage.setItem('page', this.form.page.toString());
    this.getData();
  }
  getPage(): void {
    const page = sessionStorage.getItem('page');
    page ?
      this.form.page = Number(page) :
      this.form.page = 1;
    // console.log('初始化获取页码', page);
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

