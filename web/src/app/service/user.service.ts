import { Injectable } from '@angular/core';
import { HttpService } from './http.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpService) { }
  // 获取验证码
  getVerificationCode = () => {
    return this.http.request({
      method: 'post',
      url: '/base/captcha',
    });
  }
  // 用户登陆
  login = (data) => {
    return this.http.request({
      method: 'post',
      url: '/base/login',
      data
    });
  }

}
