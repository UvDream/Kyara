import { Injectable } from '@angular/core';
import { HttpService } from './http.service';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private http: HttpService) { }
  // 添加公告
  addNotice = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/admin/addNotice',
      data
    });
  }
  // 获取博客留言
  getBlogComment = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/admin/blogComment',
      data
    });
  }
  // 审核博客留言
  checkBlogComment = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/admin/checkComment',
      data
    });
  }
}
