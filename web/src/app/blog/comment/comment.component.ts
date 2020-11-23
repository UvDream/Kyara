import { Component, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BlogService } from '@service/blog.service';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.less']
})
export class CommentComponent implements OnInit {
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
  constructor(
    private message: NzMessageService,
    private blogHttp: BlogService,
    private titleService: Title
  ) { }

  ngOnInit(): void {
    this.titleService.setTitle('留言-汪中杰的个人博客');
    this.getAvatar();
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
    }
  }


}
