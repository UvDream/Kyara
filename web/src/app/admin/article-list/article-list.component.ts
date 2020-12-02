import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { Router } from '@angular/router';
interface ListItem {
  ID: number;
  title: string;
  introduction: string;
  CreatedAt: number;
  UpdatedAt: string;
  user_name: string;
  view_count: string;
}
export interface Data {
  id: number;
  name: string;
  age: number;
  address: string;
  disabled: boolean;
}
@Component({
  selector: 'app-article-list',
  templateUrl: './article-list.component.html',
  styleUrls: ['./article-list.component.less']
})
export class ArticleListComponent implements OnInit {
  list: ListItem[] = [];
  loading = true;
  public form = {
    search: '',
    page: 1,
    pageSize: 10,
  };
  public totalCount = 0;
  indeterminate = false;
  checked = false;
  setOfCheckedId = new Set<number>();
  constructor(
    private articleService: ArticleService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.getData();
  }
  getData = async () => {
    this.loading = true;
    const res = await this.articleService.articleList(this.form);
    if (res.code === 200) {
      this.loading = false;
      this.list = res.data.msg;
      this.totalCount = res.data.totalCount;
    }

  }
  onCurrentPageSizeChange(pageSize: number): void {
    this.form.pageSize = pageSize;
    this.getData();
  }
  onCurrentPageChange(page: number): void {
    this.form.page = page;
    this.getData();
  }
  edit(data: ListItem, id: number): void {
    console.log(data, id);
    if (id === 1) {
      this.router.navigate(['/editArticle'], { queryParams: { id: data.ID } });

    }
  }

}
