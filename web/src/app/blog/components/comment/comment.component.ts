import { Component, Input, OnInit } from '@angular/core';
@Component({
  selector: 'app-comment-component',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.less']
})
export class CommentComponent implements OnInit {
  @Input() data: Array<any>;
  @Input() isChildren = false;
  constructor(
  ) { }
  private activeId: number;

  ngOnInit(): void {
    console.log(this.data, this.isChildren);
  }
  replyFunc(item: number): void {
    this.activeId = item;

  }

}
