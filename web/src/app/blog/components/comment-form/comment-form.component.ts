import { Component, Input, OnInit, DoCheck, Inject, PLATFORM_ID } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BlogService } from '@service/blog.service';
import { UserService } from '@service/user.service';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-comment-form',
  templateUrl: './comment-form.component.html',
  styleUrls: ['./comment-form.component.less']
})
export class CommentFormComponent implements OnInit, DoCheck {
  @Input() replyID: number;
  userExist = false;
  form = {
    user_name: '',
    email: '',
    avatar: '',
    blog_url: '',
    comment_content: '',
    user_id: '',
    ID: '',
    parent_id: '',
    is_private: false
  };
  constructor(
    private message: NzMessageService,
    private blogHttp: BlogService,
    private userService: UserService,
    @Inject(PLATFORM_ID) private platformId: object
  ) { }
  ngDoCheck(): void {
    if (this.replyID) {
      this.form.parent_id = this.replyID.toString();
    }
  }
  ngOnInit(): void {
    if (isPlatformBrowser(this.platformId)) {
      const isComment = localStorage.getItem('comment');
      this.userService.getUserInfo();
      if (this.userService.userInfo) {
        this.userExist = true;
        this.form.user_name = this.userService.userInfo.nickName;
        this.form.email = this.userService.userInfo.email;
        this.form.avatar = this.userService.userInfo.headerImg;
        this.form.blog_url = this.userService.userInfo.blog_url;
      } else {
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
      }
    }

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
      if (isPlatformBrowser(this.platformId)) {
        localStorage.setItem('comment', JSON.stringify(this.form));
      }
      this.userExist = true;
    }
  }

}
