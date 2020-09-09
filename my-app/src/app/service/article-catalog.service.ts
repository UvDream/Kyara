import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ArticleCatalogService {
  public ArticleCatList = [];
  constructor() { }
  SetArticleCat = (arr: any) => {
    console.log('传入的文章', arr);
  }
}
