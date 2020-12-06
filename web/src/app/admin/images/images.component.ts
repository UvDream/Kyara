import { Component, OnInit } from '@angular/core';
import { AdminService } from '@service/admin.service';

@Component({
  selector: 'app-images',
  templateUrl: './images.component.html',
  styleUrls: ['./images.component.less']
})
export class ImagesComponent implements OnInit {
  list = [];
  totalCount = 0;
  form = {
    page: 1,
    pageSize: 10
  };
  constructor(
    private adminService: AdminService,
  ) { }

  ngOnInit(): void {
    this.getData();
  }
  async getData(): Promise<void> {
    const res = await this.adminService.imgList(this.form);
    if (res.code === 200) {
      this.list = res.data.list;
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

}
