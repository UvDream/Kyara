import { Component, OnInit, Inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { Router } from '@angular/router';


@Component({
  selector: 'app-left-footer',
  templateUrl: './left-footer.component.html',
  styleUrls: ['./left-footer.component.less']
})
export class LeftFooterComponent implements OnInit {

  constructor(
    private router: Router,
    @Inject(PLATFORM_ID) private platformId: object
  ) { }

  ngOnInit(): void {
  }
  toAdmin(): void {
    if (isPlatformBrowser(this.platformId)) {
      localStorage.getItem('Authorization') ? this.router.navigate(['/admin']) : this.router.navigate(['/account/login']);
    }

  }

}
