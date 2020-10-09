import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { ArticleCatalogService } from '../../../../../service/article-catalog.service';
@Component({
  selector: 'app-catalog',
  templateUrl: './catalog.component.html',
  styleUrls: ['./catalog.component.less']
})
export class CatalogComponent implements OnInit {

  constructor(
    private catalog: ArticleCatalogService,
    private route: ActivatedRoute,
  ) {

  }
  public list = [];
  location: Location;
  public pathName = '';
  ngOnInit(): void {
    this.pathName = location.pathname;
    setTimeout(() => {
      this.list = this.catalog.ArticleCatList;
    }, 500);
  }


}
