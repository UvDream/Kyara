import {Component, OnInit} from '@angular/core';
import {ArticleService} from '../../../../../service/article.service';
import {Router} from '@angular/router';

interface HotList {
  ID: number;
  title: string;
  like_count: string;
  avatar: string;
}

@Component({
  selector: 'app-hot-article',
  templateUrl: './hot-article.component.html',
  styleUrls: ['./hot-article.component.less']
})
export class HotArticleComponent implements OnInit {
  public list: Array<HotList>;

  constructor(private service: ArticleService, private router: Router) {
  }

  ngOnInit(): void {
    this.getHotArticle();
  }

  getHotArticle = async () => {
    const res = await this.service.hotArticle();
    if (res.code === 200) {
      this.list = res.data;
      const arr = this.randomNumber(5);
      this.list.forEach((element: HotList, index: number) => {
        element.avatar = './assets/images/hot-article/' + arr[index] + '.jpg';
      });
    }
  }
  randomNumber = (num: number) => {
    const arr = [];
    for (let i = 0; i < num; i++) {
      arr[i] = i + 1;
    }
    arr.sort(() => {
      return 0.5 - Math.random();
    });
    return arr;
  }
  toDetail = (id: number) => {
    this.router.navigate(['/detail'], {queryParams: {id}});
  }

}
