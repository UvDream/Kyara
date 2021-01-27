import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ArticleService } from '@service/article.service';
@Component({
  selector: 'app-archives',
  templateUrl: './archives.component.html',
  styleUrls: ['./archives.component.less']
})
export class ArchivesComponent implements OnInit {
  public List = [

    {
      time: '2020-11-01T15:04:05+08:00',
      list: [
        {
          id: 2,
          created_at: '2020-11-06T17:22:46+08:00',
          updated_at: '2020-11-25T15:49:37+08:00',
          title: '文章标题文章标题文章标题文章标题文章标题文章标题文章标题文章标题'
        },

      ]
    },
    {
      time: '2020-11-01T15:04:05+08:00',
      list: [

      ]
    }
  ];
  constructor(private articleHttp: ArticleService, private router: Router) { }

  ngOnInit(): void {
    this.getData();
  }
  async getData(): Promise<void> {
    const res = await this.articleHttp.getBlogArchive();
    if (res.code === 200) {
      this.List = res.data;
    }
  }
  toArticle(id: number): void {
    this.router.navigate(['/detail'], { queryParams: { id } });
  }

}
