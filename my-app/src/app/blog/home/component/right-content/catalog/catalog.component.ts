import { Component, OnInit, DoCheck } from '@angular/core';
import { Location } from '@angular/common';
import { ArticleCatalogService } from '@service/article-catalog.service';
@Component({
  selector: 'app-catalog',
  templateUrl: './catalog.component.html',
  styleUrls: ['./catalog.component.less']
})
export class CatalogComponent implements OnInit, DoCheck {

  constructor(
    private catalog: ArticleCatalogService,
  ) {

  }

  public list = [];
  location: Location;
  public pathName = '';
  ngOnInit(): void {
    this.pathName = location.pathname;
  }


  ngDoCheck(): void {
    if (this.list !== this.catalog.ArticleCatList) {
      this.list = this.catalog.ArticleCatList;
    }
  }



}
