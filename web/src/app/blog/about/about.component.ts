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
    /* 一级标题 */
    .uvdream h1 {
        font-size     : 2.1em;
        line-height   : 1.1em;
        padding-top   : 16px;
        padding-bottom: 10px;
        margin-bottom : 4px;
        text-align:center;
        border-bottom : 1px solid #c99833;
    }
`;
    document.getElementById('blogger').innerHTML = style;
  }

}
