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
  // 上传图片
  uploadImage = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/fileUploadAndDownload/upload',
      data
    });
  }
  // 图片列表
  imgList = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/fileUploadAndDownload/getFileList',
      data
    });
  }
  // 删除图片
  deleteImg = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/fileUploadAndDownload/deleteFile',
      data
    });
  }
  // 修改或者修改分类
  editClassify = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/admin/editClassify',
      data
    });
  }
}
