import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BlogService } from '@service/blog.service';
@Component({
  selector: 'app-interview-detail',
  templateUrl: './interview-detail.component.html',
  styleUrls: ['./interview-detail.component.less']
})
export class InterviewDetailComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private blogHttp: BlogService
  ) { }
  public detail = {
    title: '',
    CreateAt: '',
    tag: [],
    tag_arr: [],
    answer_md: '',
    level: '',
    classify_id: ''
  };
  public showAnswer = false;
  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.showAnswer = false;
      this.getInterviewDetail(params.id);
    });
  }
  async getInterviewDetail(id: number): Promise<void> {
    const res = await this.blogHttp.getInterviewDetail({ id });
    if (res.code === 200) {
      this.detail = res.data;
    }
  }

}
