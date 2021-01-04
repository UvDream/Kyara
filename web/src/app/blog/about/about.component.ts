import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { BlogConfigService } from '@service/blog-config.service';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.less']
})
export class AboutComponent implements OnInit {

  constructor(private titleService: Title, public config: BlogConfigService) { }

  ngOnInit(): void {
    this.titleService.setTitle('关于本站-汪中杰的个人博客');
    const style = `
    .uvdream {}
    .uvdream h1 {
        text-align: center;
        background: linear-gradient(#fff 80%, #ffb11b 20%);
    }
`;
    document.getElementById('blogger').innerHTML = style;
  }

}
