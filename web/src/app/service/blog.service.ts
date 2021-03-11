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
  // 获取留言
  getComment = (data?: object) => {
    return this.http.request({
      method: 'post',
      url: '/article/getComment',
      data
    });
  }
  // 获取题库分类
  getInterview = (params?: object) => {
    return this.http.request({
      method: 'get',
      url: '/interview/classify',
      params
    });
  }
  // 获取题库列表
  getInterviewList = (data?: object) => {
    return this.http.request({
      method: 'post',
      url: '/interview/list',
      data
    });
  }
  // 获取题库详情
  getInterviewDetail = (params?: object) => {
    return this.http.request({
      method: 'get',
      url: '/interview/detail',
      params
    });
  }

}
