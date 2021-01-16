import {Injectable} from '@angular/core';
import {HttpService} from './http.service';

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

  constructor(private http: HttpService) {
  }

  setUserInfo(info: UserInfoItem): void {
    localStorage.setItem('user', JSON.stringify(info));
    this.userInfo = info;
  }

  getUserInfo(): void {
    if (!this.userInfo) {
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
