import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
@Component({
  selector: 'app-left-menus',
  templateUrl: './left-menus.component.html',
  styleUrls: ['./left-menus.component.less']
})
export class LeftMenusComponent implements OnInit {

  constructor(public http: BlogConfigService) { }

  ngOnInit(): void {
  }

}
