import { Component, OnInit } from '@angular/core';
@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.less'],
})
export class BlogComponent implements OnInit {
  constructor() { }
  title = '首页';
  ngOnInit(): void {
  }
}
