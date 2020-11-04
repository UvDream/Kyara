import { Component, OnInit } from '@angular/core';
import {BlogConfigService} from '../service/blog-config.service';
@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.less'],
})
export class BlogComponent implements OnInit {
  constructor(private http: BlogConfigService) { }
  title = '首页';
  ngOnInit(): void {
    this.http.getConfig();
  }
}
