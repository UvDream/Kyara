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
  form = {
    user_name: '',
    email: '',
    avatar: '',
    blog_url: '',
    comment_content: '',
    user_id: '',
    ID: '',
    is_private: false
  };
  commentForm = {
    page: 1,
    page_size: 10
  };
  commentData = [
  ];
  constructor(
    private message: NzMessageService,
    private blogHttp: BlogService,
    private titleService: Title
  ) { }

  ngOnInit(): void {
    this.titleService.setTitle('留言-汪中杰的个人博客');
    const isComment = localStorage.getItem('comment');
    if (isComment) {
      this.userExist = true;
      this.form.user_name = JSON.parse(isComment).user_name;
      this.form.email = JSON.parse(isComment).email;
      this.form.avatar = JSON.parse(isComment).avatar;
      this.form.blog_url = JSON.parse(isComment).blog_url;
    } else {
      this.userExist = false;
      this.getAvatar();
    }
    // 
    this.getComment();
  }
  async getAvatar(): Promise<void> {
    const res = await this.blogHttp.getAvatar();
    if (res.code === 200) {
      this.form.avatar = res.data;
    }
  }
  async commentFunc(): Promise<void> {
    if (this.form.user_name === '') {
      this.message.create('error', '请输入昵称!');
      return;
    }
    if (this.form.email === '') {
      this.message.create('error', '请输入邮箱!(将保密)');
      return;
    }
    if (this.form.comment_content === '') {
      this.message.create('error', '请输入留言内容');
      return;
    }
    const res = await this.blogHttp.addComment(this.form);
    if (res.code === 200) {
      this.message.create('success', '留言成功,博主会及时审核回复展示!');
      this.form.comment_content = '';
      localStorage.setItem('comment', JSON.stringify(this.form));
    }
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
