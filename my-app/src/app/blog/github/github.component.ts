import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { Title } from '@angular/platform-browser';

interface GithubItem {
  name: string;
  full_name: string;
  node_id: string;
  description: string;
  forks_count: string;
  stargazers_count: string;
  html_url: string;
  language: string;
  color?: string;
  fontColor?: string;
}
@Component({
  selector: 'app-github',
  templateUrl: './github.component.html',
  styleUrls: ['./github.component.less']
})
export class GithubComponent implements OnInit {

  constructor(private http: ArticleService,  private titleService: Title) { }
  public GithubList: Array<GithubItem>;
  public Loading = false;
  public ColorObject = {
    TypeScript: {
      color: '#2b7489',
      fontColor: '#fff'
    },
    Go: {
      color: '#00ADD8',
      fontColor: '#fff'
    },
    JavaScript: {
      color: '#f1e05a',
      fontColor: '#fff'
    },
    CSS: {
      color: '#563d7c',
      fontColor: '#fff'
    },
    HTML: {
      color: '#e34c26',
      fontColor: '#fff'
    },
    Vue: {
      color: '#2c3e50',
      fontColor: '#fff'
    },
    Less: {
      color: '#1d365d',
      fontColor: '#fff'
    },
    '': {
      color: '#ededed',
      fontColor: '#fff'
    },
    Java: {
      color: '#b07219',
      fontColor: '#fff'
    },
    SCSS: {
      color: '#c6538c',
      fontColor: '#fff'
    },
    Dart: {
      color: '#00B4AB',
      fontColor: '#fff'
    }
  };
  ngOnInit(): void {
    // this.GithubList = [{
    //   name: '#b07219',
    //   full_name: '',
    //   node_id: '',
    //   description: '',
    //   forks_count: '',
    //   stargazers_count: '',
    //   html_url: '',
    //   language: 'TypeScript',
    //   color: '#b07219'
    // }];
    this.getGithub();
    this.titleService.setTitle('github仓库列表-汪中杰的个人博客');

  }
  getGithub = async () => {
    this.Loading = true;
    const res = await this.http.getGithub();
    if (res.code === 200) {
      this.Loading = false;
      this.GithubList = res.data;
      this.GithubList.forEach((element: GithubItem) => {
        if (element.language !== '') {
          element.color = this.ColorObject[element.language].color;
          element.fontColor = this.ColorObject[element.language].fontColor;
        }

      });
    }
  }

}
