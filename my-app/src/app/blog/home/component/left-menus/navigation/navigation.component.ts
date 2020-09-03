import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.less'],
})
export class NavigationComponent implements OnInit {
  constructor() {}
  public menus = [
    { title: '首页', id: 1, icon: 'uv-home-outline' },
    { title: '仓库', id: 2, icon: 'uv-git-merge-outline' },
    { title: '相册', id: 6, icon: 'uv-image-outline' },
    { title: '归档', id: 3, icon: 'uv-file-tray-full-outline' },
    { title: '留言', id: 4, icon: 'uv-at-sharp' },
    { title: '关于', id: 5, icon: 'uv-cafe-outline' },
  ];
  ngOnInit(): void {}
}
