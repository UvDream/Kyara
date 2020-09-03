import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-tags',
  templateUrl: './tags.component.html',
  styleUrls: ['./tags.component.less'],
})
export class TagsComponent implements OnInit {
  constructor() {}
  public tags = [
    { id: 1, name: '前端' },
    { id: 2, name: 'web' },
    { id: 3, name: '后端' },
    { id: 4, name: '前端' },
    { id: 5, name: 'web' },
    { id: 6, name: '后端' },
  ];
  ngOnInit(): void {}
}
