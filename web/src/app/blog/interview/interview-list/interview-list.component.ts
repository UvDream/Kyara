import { Component, Inject, OnInit, PLATFORM_ID } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BlogService } from '@service/blog.service';
import { CopyText } from '../../../../util/util';
import { isPlatformBrowser } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-interview-list',
  templateUrl: './interview-list.component.html',
  styleUrls: ['./interview-list.component.less']
})
export class InterviewListComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private blogHttp: BlogService,
    @Inject(PLATFORM_ID) private platformId: object,
    private router: Router
  ) { }
  form = {
    page: 1,
    page_size: 10,
    classify: ''
  };
  list = [];
  classifyName = '';

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      console.log(params);
      this.classifyName = params.name;
      this.form.classify = params.id.toString();
      this.getInterviewList();
    });
    if (isPlatformBrowser(this.platformId)) {
      CopyText();
    }
  }
  async getInterviewList(): Promise<void> {
    const res = await this.blogHttp.getInterviewList(this.form);
    if (res.code === 200) {
      this.list = res.data.data;
    }
  }
  toDetail(id: number): void {
    this.router.navigate(['/interviewDetail'], { queryParams: { id } });
  }
}
