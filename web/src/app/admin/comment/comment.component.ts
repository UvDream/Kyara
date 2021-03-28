import { Component, OnInit } from '@angular/core';
import { AdminService } from '@service/admin.service';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.less']
})
export class CommentComponent implements OnInit {
  list = [];
  isVisible = false;
  totalCount = 0;
  constructor(
    private adminService: AdminService,
    private message: NzMessageService
  ) { }
  form = {
    search: '',
    page: 1,
    page_size: 10
  };
  replyForm = {
    id: '',
    comment: '',
    status: '',
    article_id: ''
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
  // 获取列表数据
  async getData(): Promise<void> {
    const res = await this.adminService.getBlogComment(this.form);
    if (res.code === 200) {
      this.list = res.data.list;
      this.totalCount = res.data.total;
    }
  }
  // 审核留言
  async check(id: number, status: boolean): Promise<void> {
    const obj = {
      ID: id,
      check: status
    };
    const res = await this.adminService.checkBlogComment(obj);
    if (res.code === 200) {
      this.getData();
    }
  }
  // 表格上回复方法
  replyFunc(item): void {
    this.replyForm.id = item.ID.toString();
    this.replyForm.status = item.status;
    this.replyForm.article_id = item.article_id;
    this.isVisible = true;
  }
  // 取消弹窗
  handleCancel(): void {
    this.isVisible = false;
    this.replyForm.comment = '';
  }
  // 确认弹窗
  async handleOk(): Promise<void> {
    if (this.replyForm.status === '0') {
      this.check(Number(this.replyForm.id), true);
    }
    if (this.replyForm.comment === '') {
      this.message.info('请输入回复内容');
      return;
    }
    const res = await this.adminService.replyComment(this.replyForm);
    if (res.code === 200) {
      this.isVisible = false;
      this.message.success('回复成功');
      this.getData();
    }
  }
  // 删除方法
  async deleteFunc(item: any, sure: boolean): Promise<void> {
    if (!sure) {
      this.message.info('取消删除');
    } else {
      const res = await this.adminService.deleteComment({ id: item.ID });
      if (res.code === 200) {
        this.message.success('删除成功!');
        this.getData();
      }
    }
  }

}
