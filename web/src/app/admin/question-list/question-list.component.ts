import { Component, OnInit } from '@angular/core';
import { BlogService } from '@service/blog.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-question-list',
  templateUrl: './question-list.component.html',
  styleUrls: ['./question-list.component.less']
})
export class QuestionListComponent implements OnInit {

  constructor(
    private blogHttp: BlogService,
    private router: Router,
    private route: ActivatedRoute
  ) { }
  list = [];
  form = {
    page: 1,
    page_size: 10,
    classify: ''
  };
  totalCount = 0;
  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      if (params.id) {
        this.form.classify = params.id.toString();
      }
    });
    this.getInterviewList();

  }

  async getInterviewList(): Promise<void> {
    const res = await this.blogHttp.getInterviewList(this.form);
    if (res.code === 200) {
      this.list = res.data.data;
      this.totalCount = res.data.total_count;
    }
  }
  pageChange(num: number): void {
    this.form.page = num;
    this.getInterviewList();
  }
  pageSizeChange(num: number): void {
    this.form.page_size = num;
    this.getInterviewList();
  }
  editFunc(id: number): void {
    this.router.navigate(['/admin/addQuestion'], { queryParams: { id } });
  }

}
