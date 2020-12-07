import { Component, OnInit } from '@angular/core';
import { AdminService } from '@service/admin.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
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
    private message: NzMessageService,
    private modal: NzModalService
  ) { }

  ngOnInit(): void {
    this.getData();
  }
  async getData(): Promise<void> {
    const res = await this.adminService.imgList(this.form);
    if (res.code === 200) {
      this.list = res.data.list;
      this.totalCount = res.data.total;
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
  delete(data: object): void {
    this.modal.confirm({
      nzTitle: '提示',
      nzContent: '确认删除此图片?',
      nzOnOk: async () => {
        const res = await this.adminService.deleteImg(data);
        if (res.code === 200) {
          this.message.success('删除成功!');
          this.getData();
        }
      },
      nzOnCancel: () => {
        this.message.info('取消删除');
      }
    });

  }

}
