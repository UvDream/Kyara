import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
import { AdminService } from '@service/admin.service';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-user-config',
  templateUrl: './user-config.component.html',
  styleUrls: ['./user-config.component.less']
})
export class UserConfigComponent implements OnInit {
  public form = {
    author_name: '',
    author_motto: '',
    author_link: '',
    author_avatar: ''
  };
  constructor(
    public configHttp: BlogConfigService,
    public adminHttp: AdminService,
    private message: NzMessageService
  ) { }

  ngOnInit(): void {
    this.form.author_name = this.configHttp.AuthorName;
    this.form.author_motto = this.configHttp.AuthorMotto;
    this.form.author_link = this.configHttp.AuthorLink;
    this.form.author_avatar = this.configHttp.AuthorAvatar;
  }
  async saveFunc(): Promise<void> {
    const res = await this.adminHttp.blogConfig(this.form);
    if (res.code !== 200) {
      this.message.error('保存错误,请联系作者!');
    } else {
      this.message.success('保存成功!');
      this.configHttp.getConfig();
    }
  }

}
