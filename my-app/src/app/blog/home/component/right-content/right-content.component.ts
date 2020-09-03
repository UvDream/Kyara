import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-right-content',
  templateUrl: './right-content.component.html',
  styleUrls: ['./right-content.component.less'],
})
export class RightContentComponent implements OnInit {
  tabs = [
    {
      active: true,
      name: 'Tab 1',
      icon: 'uv-thumbs-up-outline',
    },
    {
      active: false,
      name: 'Tab 2',
      icon: 'uv-chatbubble-ellipses-outline',
    },
    {
      active: false,
      name: 'Tab 2',
      icon: 'uv-star-outline',
    },
  ];
  constructor() {}

  ngOnInit(): void {}
}
