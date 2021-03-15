import { Component, OnInit } from '@angular/core';
import { BlogService } from '@service/blog.service';

@Component({
  selector: 'app-question-classify',
  templateUrl: './question-classify.component.html',
  styleUrls: ['./question-classify.component.less']
})
export class QuestionClassifyComponent implements OnInit {

  constructor(
    private blogHttp: BlogService,
  ) { }
  public interviewList = [];
  isVisible = false;
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
  editFunc(item?: object): void {
    this.isVisible = true;
  }
  // 删除分类
  deleteFunc(id: number): void {
    this.isVisible = true;
  }
  handleCancel(): void {
    this.isVisible = false;
  }
  handleOk(): void {
    this.isVisible = false;
  }
}
