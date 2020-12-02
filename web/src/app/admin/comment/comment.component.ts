import { Component, OnInit } from '@angular/core';
import { AdminService } from '@service/admin.service';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.less']
})
export class CommentComponent implements OnInit {
  list = [];
  totalCount = 0;
  constructor(
    private adminService: AdminService,
  ) { }
  form = {
    search: '',
    page: 1,
    page_size: 10
  };
  ngOnInit(): void {
    this.getData();
  }
  onCurrentPageSizeChange(pageSize: number): void {
    this.form.page_size = pageSize;
    this.getData();
  }
  onCurrentPageChange(page: number): void {
    this.form.page = page;
    this.getData();
  }
  async getData(): Promise<void> {
    const res = await this.adminService.getBlogComment(this.form);
    if (res.code === 200) {
      this.list = res.data.list;
      this.totalCount = res.data.total;
    }
  }

}
