import { Injectable } from '@angular/core';
import { HttpService } from './http.service';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {

  constructor(private http: HttpService) {
  }

  // 文章列表
  articleList = (data?: any) => {
    return this.http.request({
      method: 'post',
      url: '/article/articleList',
      data
    });
  }
  // 文章详情
  getDetail = (params: any) => {
    return this.http.request({
      method: 'get',
      url: '/article/articleDetail',
      params
    });
  }
  // 验证加密文章密码
  surePassword = (params: any) => {
    return this.http.request({
      method: 'get',
      url: '/article/articleDetail',
      params
    });
  }
  // 获取文章分类
  getArticleClassification = (params?: any) => {
    return this.http.request({
      method: 'get',
      url: '/article/articleClassify',
      params
    });
  }
  // 保存文章
  addArticle = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/admin/add',
      data
    });
  }
  //  热门文章
  hotArticle = () => {
    return this.http.request({
      method: 'get',
      url: '/article/hotArticle',
    });
  }
  //  获取所有tag
  getTag = () => {
    return this.http.request({
      method: 'get',
      url: '/article/tag',
    });
  }
  // 增加tag
  addTag = (params?: object) => {
    return this.http.request({
      method: 'get',
      url: '/admin/addTag',
      params
    });
  }
  // 获取配置
  getConfig = () => {
    return this.http.request({
      method: 'get',
      url: '/article/config',
    });
  }
  // 获取github仓库
  getGithub = () => {
    return this.http.request({
      method: 'get',
      url: '/article/github',
    });
  }
  // admin 获取详情
  getArticleDetail = (id: string) => {
    return this.http.request({
      method: 'get',
      url: '/admin/articleDetail?id=' + id,
    });
  }
  // 博客访问量
  viewBlog = (id?: string) => {
    return this.http.request({
      method: 'get',
      url: '/article/views?id=' + id,
    });
  }
  // 获取公告
  getNotice = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/article/noticeList',
      data
    });
  }
  // 百度收录
  toBaidu = (url?: string) => {
    const data = {
      argument: url
    };
    if (!environment) {
      return '';
    }
    return this.http.request({
      method: 'post',
      url: '/other/baidu',
      data
    });
  }
  // 获取分类动态
  getClassify = () => {
    return this.http.request({
      method: 'get',
      url: '/other/classifyStatistics'
    });
  }
  // 获取博客动态
  getBlogDynamic = () => {
    return this.http.request({
      method: 'get',
      url: '/other/blogDynamic'
    });
  }
  // 获取文章归档数据
  getBlogArchive = () => {
    return this.http.request({
      method: 'get',
      url: '/article/blogArchives'
    });
  }


}
