import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AdminService } from '@service/admin.service';
import { NzMessageService } from 'ng-zorro-antd/message';

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
  // 验证表单
  validateForm!: FormGroup;

  constructor(
    private articleService: ArticleService,
    private fb: FormBuilder,
    private adminService: AdminService,
    private message: NzMessageService
  ) { }
  indeterminate = false;
  loading = true;
  checked = false;
  isVisible = false;
  public form = {
    search: '',
    page: 1,
    page_size: 10,
  };
  public data = {
    ID: 0,
    title: '',
    content: '',
    show: false,
  };
  public totalCount = 0;
  list: ListItem[] = [];
  ngOnInit(): void {
    this.getData();
    this.validateForm = this.fb.group({
      title: [null, [Validators.required]],
      content: [null, [Validators.required]],
      show: [null]
    });
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
  onCurrentPageChange(page: number): void {
    this.form.page = page;
    this.getData();
  }
  edit(data: ListItem): void {
    console.log(data);
    this.isVisible = true;
    this.data.ID = data.ID;
    this.data.title = data.title;
    this.data.content = data.content;
    data.show === '1' ? this.data.show = true : this.data.show = false;
  }
  showModal(): void {
    this.isVisible = true;
  }

  handleOk(): void {
    console.log('Button ok clicked!');
    this.isVisible = false;
  }

  handleCancel(): void {
    console.log('Button cancel clicked!');
    this.isVisible = false;
  }

  add(): void {
    this.data = {
      ID: 0,
      title: '',
      content: '',
      show: false
    };
    this.isVisible = true;
  }
  async submitForm(): Promise<void> {
    // tslint:disable-next-line:forin
    for (const i in this.validateForm.controls) {
      this.validateForm.controls[i].markAsDirty();
      this.validateForm.controls[i].updateValueAndValidity();
    }
    if (this.validateForm.valid) {
      const obj = {
        ID: this.data.ID,
        title: this.data.title,
        content: this.data.content,
        show: '0'
      };
      if (this.validateForm.value.show) {
        obj.show = '1';
      }
      const res = await this.adminService.addNotice(obj);
      if (res.code === 200) {
        console.log('保存成功');
        this.getData();
        this.isVisible = false;
      }
    }
  }
  // 删除
  async deleteFunc(item: any, sure: boolean): Promise<void> {
    if (!sure) {
      this.message.info('取消删除');
    } else {
      const res = await this.adminService.deleteNotice({ id: item.ID });
      if (res.code === 200) {
        this.message.success('删除成功!');
        this.getData();
      }
    }
  }
}
