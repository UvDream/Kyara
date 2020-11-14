import { Injectable } from '@angular/core';
import { Location } from '@angular/common';
@Injectable({
  providedIn: 'root'
})
export class ArticleCatalogService {
  public ArticleCatList = [];
  public pathName = '';
  location: Location;
  constructor() { }
  SetCatLog = (arr: any) => {
    this.pathName = location.pathname;
    console.log('检测', this.pathName, arr)
    this.ArticleCatList = arr;
  }

}
