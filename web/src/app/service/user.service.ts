import { Injectable } from '@angular/core';
import { HttpService } from './http.service';
import { Inject, PLATFORM_ID } from '@angular/core';

import { isPlatformBrowser } from '@angular/common';

interface UserInfoItem {
  ID: number;
  userName: string;
  nickName: string;
  headerImg: string;
  email: string;
  blog_url: string;
  uuid: string;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  public userInfo: UserInfoItem;

  constructor(
    private http: HttpService,
    @Inject(PLATFORM_ID) private platformId: object
  ) {
  }

  setUserInfo(info: UserInfoItem): void {
    if (isPlatformBrowser(this.platformId)) {
      localStorage.setItem('user', JSON.stringify(info));
    }
    this.userInfo = info;
  }

  getUserInfo(): void {
    if (!this.userInfo && isPlatformBrowser(this.platformId)) {
      this.userInfo = JSON.parse(localStorage.getItem('user'));
    }

  }

  // 获取验证码
  getVerificationCode = () => {
    return this.http.request({
      method: 'post',
      url: '/base/captcha',
    });
  };
  // 用户登陆
  login = (data) => {
    return this.http.request({
      method: 'post',
      url: '/base/login',
      data
    });
  };

}
