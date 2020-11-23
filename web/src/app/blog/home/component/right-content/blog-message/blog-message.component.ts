import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
@Component({
  selector: 'app-blog-message',
  templateUrl: './blog-message.component.html',
  styleUrls: ['./blog-message.component.less']
})
export class BlogMessageComponent implements OnInit {

  constructor(public config: BlogConfigService) { }

  ngOnInit(): void {
  }

}
