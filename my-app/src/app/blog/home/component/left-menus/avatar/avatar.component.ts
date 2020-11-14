import { Component, OnInit } from '@angular/core';
import { BlogConfigService } from '@service/blog-config.service';
@Component({
  selector: 'app-avatar',
  templateUrl: './avatar.component.html',
  styleUrls: ['./avatar.component.less']
})
export class AvatarComponent implements OnInit {

  constructor(public config: BlogConfigService) { }

  ngOnInit(): void {
  }

}
