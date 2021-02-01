import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';

@Component({
  selector: 'app-blog-config',
  templateUrl: './blog-config.component.html',
  styleUrls: ['./blog-config.component.less']
})
export class BlogConfigComponent implements OnInit {

  constructor(
    private configHttp: BlogConfigService
  ) { }

  ngOnInit(): void {
    this.configHttp.getConfig();
  }

}
