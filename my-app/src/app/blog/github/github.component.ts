import { Component, OnInit } from '@angular/core';
import { ArticleService } from '../../service/article.service';
@Component({
  selector: 'app-github',
  templateUrl: './github.component.html',
  styleUrls: ['./github.component.less']
})
export class GithubComponent implements OnInit {

  constructor(private http: ArticleService) { }
  public GithubList = [];
  ngOnInit(): void {
    this.getGithub();
  }
  getGithub = async () => {
    const res = await this.http.getGithub();
    console.log(res);
    if (res.code === 200) {
      this.GithubList = res.data;
    }
  }
 
}
