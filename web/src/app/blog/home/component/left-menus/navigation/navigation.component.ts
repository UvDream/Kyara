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
  constructor(private router: Router, private catalogService: ArticleCatalogService,) { }
  public menus = [
    { title: '首页', id: 1, icon: 'uv-home-outline', url: '/home' },
    { title: '仓库', id: 2, icon: 'uv-git-merge-outline', url: '/github' },
    { title: '相册', id: 6, icon: 'uv-image-outline', url: '/photos' },
    { title: '归档', id: 3, icon: 'uv-file-tray-full-outline', url: '/archives' },
    { title: '留言', id: 4, icon: 'uv-at-sharp', url: '/comment' },
    { title: '关于', id: 5, icon: 'uv-cafe-outline', url: '/about' },
  ];
  public ActiveMenus = 1;
  ngOnInit(): void { }
  // tslint:disable-next-line:typedef
  navFunc(item: MenuItem) {
    this.ActiveMenus = item.id;
    this.router.navigate([item.url]);
    this.catalogService.SetCatLog()
  }
}
