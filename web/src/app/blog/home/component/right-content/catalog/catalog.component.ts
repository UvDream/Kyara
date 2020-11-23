import { Component, OnInit, DoCheck } from '@angular/core';
import { Location } from '@angular/common';
import { ArticleCatalogService } from '@service/article-catalog.service';
import { Platform } from '@angular/cdk/platform';

@Component({
  selector: 'app-catalog',
  templateUrl: './catalog.component.html',
  styleUrls: ['./catalog.component.less']
})
export class CatalogComponent implements OnInit, DoCheck {

  constructor(
    private catalog: ArticleCatalogService,
    private platform: Platform,
  ) {

  }

  public list = [];
  location: Location;
  public pathName = '';
  ngOnInit(): void {
    if (this.platform.isBrowser) {
    this.pathName = location.pathname;
    }
  }


  ngDoCheck(): void {
    if (this.list !== this.catalog.ArticleCatList) {
      this.list = this.catalog.ArticleCatList;
    }
  }



}
