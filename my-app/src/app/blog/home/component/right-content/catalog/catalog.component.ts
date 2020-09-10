import { Component, OnInit } from '@angular/core';
import { ArticleCatalogService } from '../../../../../service/article-catalog.service';
@Component({
  selector: 'app-catalog',
  templateUrl: './catalog.component.html',
  styleUrls: ['./catalog.component.less']
})
export class CatalogComponent implements OnInit {

  constructor(private catalog: ArticleCatalogService) { }
  public list = [];
  ngOnInit(): void {
    setTimeout(() => {
      this.list = this.catalog.ArticleCatList;
      console.log(this.list);
    }, 1000);
  }


}
