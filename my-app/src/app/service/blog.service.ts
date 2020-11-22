import { Injectable } from '@angular/core';
import { HttpService } from './http.service';
@Injectable({
  providedIn: 'root'
})
export class BlogService {

  constructor(private http: HttpService) { }
  // 博客留言
  addComment = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/article/comment',
      data
    });
  }
  // 随机获取头像
  getAvatar = (params?: any) => {
    return this.http.request({
      method: 'get',
      url: '/openApi/avatar',
      params
    });
  }
}
