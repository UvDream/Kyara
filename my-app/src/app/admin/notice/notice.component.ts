import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
interface ListItem {
  ID: number;
  title: string;
  content: string;
  CreatedAt: number;
  UpdatedAt: string;
  show: string;
}
@Component({
  selector: 'app-notice',
  templateUrl: './notice.component.html',
  styleUrls: ['./notice.component.less']
})
export class NoticeComponent implements OnInit {
  indeterminate = false;
  loading = true;
  checked = false;
  setOfCheckedId = new Set<number>();
  public form = {
    search: '',
    page: 1,
    page_size: 10,
  };
  public totalCount = 0;
  list: ListItem[] = [];
  constructor(private articleService: ArticleService) { }

  ngOnInit(): void {
    this.getData();
  }
  async getData(): Promise<void> {
    this.loading = true;
    const res = await this.articleService.getNotice(this.form);
    if (res.code === 200) {
      this.loading = false;
      this.totalCount = res.data.total_count;
      this.list = res.data.list;
    }
  }
  onCurrentPageSizeChange(pageSize: number): void {
    this.form.page_size = pageSize;
    this.getData();
  }
  onCurrentPageDataChange(page: number): void {
    this.form.page = page;
    this.getData();
  }
  edit(): void {

  }
}
