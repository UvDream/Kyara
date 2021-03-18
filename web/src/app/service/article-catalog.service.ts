import { Inject, Injectable, PLATFORM_ID } from '@angular/core';
import { Location } from '@angular/common';
import { isPlatformBrowser } from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class ArticleCatalogService {
  public ArticleCatList = [];
  public pathName = '';
  location: Location;
  constructor(
    @Inject(PLATFORM_ID) private platformId: object
  ) { }
  SetCatLog = (arr?: any) => {
    if (isPlatformBrowser(this.platformId)) {
      this.pathName = location.pathname;
    }
    this.ArticleCatList = arr;
  }

}
