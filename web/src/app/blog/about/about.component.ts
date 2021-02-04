import { Component, OnInit, Inject, PLATFORM_ID } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { BlogConfigService } from '@service/blog-config.service';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.less']
})
export class AboutComponent implements OnInit {

  constructor(
    private titleService: Title,
    public config: BlogConfigService,
    @Inject(PLATFORM_ID) private platformId: object
  ) { }

  ngOnInit(): void {
    this.titleService.setTitle('关于本站-汪中杰的个人博客');
    const style = `
    .uvdream {}
    .uvdream h1 {
        text-align: center;
        background: linear-gradient(#fff 80%, #ffb11b 20%);
    }
`;
    if (isPlatformBrowser(this.platformId)) {
      document.getElementById('blogger').innerHTML = style;
    }
  }

}
