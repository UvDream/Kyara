import { Component, OnInit } from '@angular/core';
import { BlogService } from '@service/blog.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-question-list',
  templateUrl: './question-list.component.html',
  styleUrls: ['./question-list.component.less']
})
export class QuestionListComponent implements OnInit {

  constructor(
    private blogHttp: BlogService,
    private router: Router
  ) { }
  list = [];
  form = {
    page: 1,
    page_size: 10,
    classify: ''
  };
  ngOnInit(): void {
    this.getInterviewList();
  }

  async getInterviewList(): Promise<void> {
    const res = await this.blogHttp.getInterviewList(this.form);
    if (res.code === 200) {
      this.list = res.data.data;
    }
  }
  editFunc(id: number): void {
    this.router.navigate(['/admin/addQuestion'], { queryParams: { id } });
  }

}
