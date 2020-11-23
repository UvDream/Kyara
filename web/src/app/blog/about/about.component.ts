import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.less']
})
export class AboutComponent implements OnInit {

  constructor( private titleService: Title) { }

  ngOnInit(): void {
    this.titleService.setTitle('关于本站-汪中杰的个人博客');
  }

}
