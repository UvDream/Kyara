import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BlogConfigService } from '@service/blog-config.service';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.less'],
})
export class HeaderComponent implements OnInit {
  constructor(private router: Router, public http: BlogConfigService) { }
  visible = false;
  ngOnInit(): void { }
  open(): void {
    this.visible = true;
  }

  close(): void {
    this.visible = false;
  }
  toHome = () => {
    this.router.navigate(['/home']);
  }
}
