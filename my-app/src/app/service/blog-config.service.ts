import { Injectable } from '@angular/core';
import { ArticleService } from './article.service'
@Injectable({
  providedIn: 'root'
})
export class BlogConfigService {

  constructor(private httpService: ArticleService) { }
  public BlogName = '';
  public AuthorName = '';
  public AuthorMotto = '';
  public BlogLogo = '';
  public AuthorAvatar = '';
  public ArticleCount = '';
  public BlogStartTime = '';
  getConfig = async () => {
    const res = await this.httpService.getConfig();
    console.log(res, '获取配置');
    if (res.code !== 200) { return; }
    this.BlogName = res.data.blog_name;
    this.AuthorName = res.data.author_name;
    this.AuthorMotto = res.data.author_motto;
    this.BlogLogo = res.data.blog_logo;
    this.AuthorAvatar = res.data.author_avatar;
    this.ArticleCount = res.data.article_count;
    this.BlogStartTime = res.data.blog_start_time;
  }
}
