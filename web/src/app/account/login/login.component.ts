import { Component, OnInit, Inject, PLATFORM_ID } from '@angular/core';
import { UserService } from '@service/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { Router } from '@angular/router';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less']
})
export class LoginComponent implements OnInit {
  form = {
    username: '',
    password: '',
    Captcha: '',
    CaptchaId: ''
  };

  public imgUrl = '';
  constructor(
    private blogHttp: UserService,
    private message: NzMessageService,
    private router: Router,
    @Inject(PLATFORM_ID) private platformId: object
  ) { }

  ngOnInit(): void {
    this.getVerificationCode();
  }
  async getVerificationCode(): Promise<void> {
    const res = await this.blogHttp.getVerificationCode();
    if (res.code === 200) {
      this.imgUrl = res.data.picPath;
      this.form.CaptchaId = res.data.captchaId;
    }
  }
  async submitForm(): Promise<void> {
    const res = await this.blogHttp.login(this.form);
    if (res.code === 200) {
      this.message.success('登陆成功!');
      this.router.navigate(['/admin']);
      this.blogHttp.setUserInfo(res.data.user);
      if (isPlatformBrowser(this.platformId)) {
        localStorage.setItem('Authorization', res.data.token);
      }
    } else {
      this.message.error(res.msg);
      this.getVerificationCode();
    }
  }

}
