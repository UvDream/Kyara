import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
import { AdminService } from '@service/admin.service';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-blog-about',
  templateUrl: './blog-about.component.html',
  styleUrls: ['./blog-about.component.less']
})
export class BlogAboutComponent implements OnInit {
  public form = {
    blog_name: '',
    blog_logo: '',
    blog_start_time: '',
    domain_name: ''
  };
  constructor(
    public configHttp: BlogConfigService,
    public adminHttp: AdminService,
    private message: NzMessageService
  ) { }

  ngOnInit(): void {
    this.form = {
      blog_name: this.configHttp.BlogName,
      blog_logo: this.configHttp.BlogLogo,
      blog_start_time: this.configHttp.BlogTime,
      domain_name: this.configHttp.DomainName
    };
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
