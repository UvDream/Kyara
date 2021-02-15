import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
@Component({
  selector: 'app-comment-component',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.less']
})
export class CommentComponent implements OnInit {
  @Input() data: Array<any>;
  @Input() isChildren = false;
  @Output() ActiveIDOut = new EventEmitter();
  constructor(
  ) { }
  private activeId: number;

  ngOnInit(): void {
    console.log(this.data, this.isChildren);
  }
  replyFunc(item: any): void {
    this.activeId = item.ID;
    this.ActiveIDOut.emit(item);
  }

}
