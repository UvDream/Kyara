import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ArticleCatalogService {
  public ArticleCatList = [];
  constructor() { }
  SetCatLog = (arr: any) => {
    this.ArticleCatList = arr;
  }
}
