import { Component, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BlogService } from '@service/blog.service';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-comment',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.less']
})
export class CommentsComponent implements OnInit {
  totalCount = 0;
  userExist = false;
  commentForm = {
    page: 1,
    page_size: 10
  };
  commentData = [
  ];
  constructor(
    private blogHttp: BlogService,
    private titleService: Title
  ) { }

  ngOnInit(): void {
    this.titleService.setTitle('留言-汪中杰的个人博客');
    this.getComment();
  }
  PageIndexChange(page: number): void {
    this.commentForm.page = page;
    this.getComment();
  }
  // 获取留言
  async getComment(): Promise<void> {
    const res = await this.blogHttp.getComment(this.commentForm);
    console.log(res);
    if (res.code === 200) {
      this.commentData = res.data.data;
      this.totalCount = res.data.total;
    }
  }


}
