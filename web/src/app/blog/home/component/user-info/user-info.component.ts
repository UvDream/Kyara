import { Component, OnInit, Inject, PLATFORM_ID } from '@angular/core';
import { UserService } from '@service/user.service';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.less']
})
export class UserInfoComponent implements OnInit {
  form = {
    username: '',
    password: '',
    Captcha: '',
    CaptchaId: ''
  };
  public imgUrl = '';

  constructor(
    public userService: UserService,
    private router: Router,
    private message: NzMessageService,
    @Inject(PLATFORM_ID) private platformId: object
  ) {
  }

  ngOnInit(): void {
    this.userService.getUserInfo();
    this.getVerificationCode();
  }
  async getVerificationCode(): Promise<void> {
    const res = await this.userService.getVerificationCode();
    if (res.code === 200) {
      this.imgUrl = res.data.picPath;
      this.form.CaptchaId = res.data.captchaId;
    }
  }
  async submitForm(): Promise<void> {
    const res = await this.userService.login(this.form);
    if (res.code === 200) {
      this.message.success('登陆成功!');
      this.userService.setUserInfo(res.data.user);
      if (isPlatformBrowser(this.platformId)) {
        localStorage.setItem('Authorization', res.data.token);
      }

    } else {
      this.message.error(res.msg);
      this.getVerificationCode();
    }
  }
  login(): void {
    this.router.navigate(['/account/login']);
  }

}
