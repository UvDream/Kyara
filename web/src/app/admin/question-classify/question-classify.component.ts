import { Component, OnInit } from '@angular/core';
import { BlogService } from '@service/blog.service';
import { NzModalService } from 'ng-zorro-antd/modal';
interface ClassifyItem {
  classify_icon: string;
  classify_name: string;
  ID: number;
}
@Component({
  selector: 'app-question-classify',
  templateUrl: './question-classify.component.html',
  styleUrls: ['./question-classify.component.less']
})
export class QuestionClassifyComponent implements OnInit {

  constructor(
    private blogHttp: BlogService,
    private modal: NzModalService
  ) { }
  public interviewList = [];
  isVisible = false;
  form = {
    classify_icon: '',
    classify_name: '',
    ID: 0
  };
  ngOnInit(): void {
    this.getData();
  }

  async getData(): Promise<void> {
    const res = await this.blogHttp.getInterview();
    if (res.code === 200) {
      this.interviewList = res.data;
    }
  }
  // 修改分类
  editFunc(item?: ClassifyItem): void {
    this.isVisible = true;
    item ? this.form = item : '';
  }
  // 删除分类
  deleteFunc(id: number): void {
    this.modal.confirm({
      nzTitle: '确认信息',
      nzContent: '确认删除这个分类',
      nzOnOk: async () => {
        const res = await this.blogHttp.deleteInterviewClassify({ id });
        if (res.code === 200) {
          this.getData();
        }
      }
    });

  }
  handleCancel(): void {
    this.isVisible = false;
  }
  // 上传图片回显
  urlChange(data: string): void {
    this.form.classify_icon = data;
  }
  async handleOk(): Promise<void> {
    const res = await this.blogHttp.editInterviewClassify(this.form);
    if (res.code === 200) {
      this.getData();
      this.isVisible = false;
    }
  }
}
