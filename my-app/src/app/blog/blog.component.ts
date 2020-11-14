import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
import { CookieService } from 'ngx-cookie-service';
import { ArticleService } from '@service/article.service';
@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.less'],
})
export class BlogComponent implements OnInit {
  time: number = 12 * 60 * 60 * 1000;
  constructor(
    private http: BlogConfigService,
    private cookieService: CookieService,
    private articleHttp: ArticleService,
  ) { }
  title = '首页';
  ngOnInit(): void {
    this.http.getConfig();
    console.log(this.cookieService.get('blogView'), '获取时间');
    if (this.cookieService.get('blogView') === '') {
      this.viewBlog();
      this.cookieService.set('blogView', 'true', new Date(new Date().getTime() + this.time));
      console.log(new Date(new Date().getTime() + this.time), '过期时间');
    }
  }
  async viewBlog(): Promise<void> {
    const res = await this.articleHttp.viewBlog('blog');
    console.log(res);
  }
}
