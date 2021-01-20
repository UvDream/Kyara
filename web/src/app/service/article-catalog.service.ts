import { Injectable } from '@angular/core';
import { Location, PlatformLocation } from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class ArticleCatalogService {
  public ArticleCatList = [];
  public pathName = '';
  location: Location;
  constructor(
    private platformLocation: PlatformLocation
  ) { }
  SetCatLog = (arr?: any) => {
    const pathName = this.platformLocation.hash;
    if (pathName.match(/\#(\S*)\?/) != null) {
      this.pathName = pathName.match(/\#(\S*)\?/)[1];
    } else {
      this.pathName = '';
    }
    this.ArticleCatList = arr;
  }

}
