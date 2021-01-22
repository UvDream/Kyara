import { Component, OnInit, DoCheck } from '@angular/core';
import { ArticleCatalogService } from '@service/article-catalog.service';

@Component({
  selector: 'app-right-content',
  templateUrl: './right-content.component.html',
  styleUrls: ['./right-content.component.less'],
})
export class RightContentComponent implements OnInit, DoCheck {
  tabs = [
    {
      active: true,
      name: '1',
      icon: 'icon-A7',
    },
    {
      active: false,
      name: '2',
      icon: 'icon-A12',
    },
    // {
    //   active: false,
    //   name: '3',
    //   icon: 'uv-star-outline',
    // },
  ];
  public pathName = '';
  constructor(private service: ArticleCatalogService) { }

  ngOnInit(): void {
    this.pathName = this.service.pathName;
  }
  ngDoCheck(): void {
    if (this.service.pathName !== this.pathName) {
      this.pathName = this.service.pathName;
    }
  }
}
