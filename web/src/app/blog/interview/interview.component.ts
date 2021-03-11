import { Component, OnInit } from '@angular/core';
import { BlogService } from '@service/blog.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-interview',
  templateUrl: './interview.component.html',
  styleUrls: ['./interview.component.less']
})
export class InterviewComponent implements OnInit {

  constructor(
    private blogHttp: BlogService,
    private router: Router
  ) { }
  public interviewList = [];
  ngOnInit(): void {
    this.getData();
  }
  async getData(): Promise<void> {
    const res = await this.blogHttp.getInterview();
    console.log(res);
    if (res.code === 200) {
      this.interviewList = res.data;
    }
  }
  tabClick(id: number, name: string): void {
    this.router.navigate(['/interviewList'], { queryParams: { id, name } });
  }

}
