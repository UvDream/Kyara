import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
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
    private articleHttp: ArticleService,
  ) { }
  title = '首页';
  ngOnInit(): void {
    this.http.getConfig();
    this.viewBlog();
  }
  async viewBlog(): Promise<void> {
    const res = await this.articleHttp.viewBlog('blog');
    console.log(res);
  }
}
