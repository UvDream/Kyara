import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BlogConfigService } from '../../../../service/blog-config.service';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.less'],
})
export class HeaderComponent implements OnInit {
  constructor(private router: Router, public http: BlogConfigService) { }

  ngOnInit(): void { }
  toHome = () => {
    this.router.navigate(['/home']);
  }
}
