import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { Router } from '@angular/router';
@Component({
  selector: 'app-right-content',
  templateUrl: './right-content.component.html',
  styleUrls: ['./right-content.component.less'],
})
export class RightContentComponent implements OnInit {
  tabs = [
    {
      active: true,
      name: '1',
      icon: 'uv-thumbs-up-outline',
    },
    {
      active: false,
      name: '2',
      icon: 'uv-chatbubble-ellipses-outline',
    },
    {
      active: false,
      name: '3',
      icon: 'uv-star-outline',
    },
  ];
  public pathName = '';
  location: Location;
  constructor(private router: Router) { }

  ngOnInit(): void {
    this.router.events.subscribe(res => {
      this.pathName = location.pathname;
    });
  }
}
