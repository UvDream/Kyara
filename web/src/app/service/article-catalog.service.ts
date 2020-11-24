import { Injectable } from '@angular/core';
import { Location } from '@angular/common';
import { Platform } from '@angular/cdk/platform';

@Injectable({
  providedIn: 'root'
})
export class ArticleCatalogService {
  public ArticleCatList = [];
  public pathName = '';
  location: Location;
  constructor(
    private platform: Platform,
  ) { }
  SetCatLog = (arr?: any) => {
    if (this.platform.isBrowser) {
      this.pathName = location.pathname;
    }
    this.ArticleCatList = arr;
  }

}
