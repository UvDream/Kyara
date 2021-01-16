import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {BlogConfigService} from '@service/blog-config.service';
import {UserService} from '@service/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.less'],
})
export class HeaderComponent implements OnInit {
  constructor(private router: Router, public blogHttp: BlogConfigService, public userService: UserService) {
  }

  visible = false;
  dropShow = false;
  public user = {};

  ngOnInit(): void {
    this.userService.getUserInfo();
  }

  open(): void {
    this.visible = true;
  }

  close(): void {
    this.visible = false;
  }

  toHome = () => {
    this.router.navigate(['/home']);
  };

  nzVisibleChange(data: boolean): void {
    this.dropShow = data;
  }
}
