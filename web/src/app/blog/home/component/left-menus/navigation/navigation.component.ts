import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ArticleCatalogService } from '@service/article-catalog.service';
interface MenuItem {
  title: string;
  id: number;
  icon: string;
  url: string;
}
@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.less'],
})
export class NavigationComponent implements OnInit {
  constructor(
    private router: Router,
    private catalogService: ArticleCatalogService
  ) { }
  public menus = [
    { title: '首页', id: 1, icon: 'icon-shouye', url: '/' },
    { title: '仓库', id: 2, icon: 'icon-fujian', url: '/github' },
    // { title: '相册', id: 6, icon: 'icon-tupian', url: '/photos' },
    { title: '面试', id: 6, icon: 'icon-guanjun', url: '/interview' },
    { title: '归档', id: 3, icon: 'icon-biaoqian', url: '/archives' },
    { title: '留言', id: 4, icon: 'icon-at', url: '/comment' },
    { title: '关于', id: 5, icon: 'icon-yiwen', url: '/about' },
  ];
  public ActiveMenus = 1;
  ngOnInit(): void { }
  navFunc(item: MenuItem): void {
    this.ActiveMenus = item.id;
    this.router.navigate([item.url]);
    this.catalogService.SetCatLog();
  }
}
