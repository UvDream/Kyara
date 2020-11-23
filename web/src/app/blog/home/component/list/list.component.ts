import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.less'],
})
export class ListComponent implements OnInit {
  constructor(private router: Router) { }
  @Input() data: any;
  @Output() HavePassword = new EventEmitter();
  ngOnInit(): void { }
  toDetail = (data) => {
    if (data.is_password === '1') {
      this.HavePassword.emit(data.ID);
    } else {
      this.router.navigate(['/detail'], { queryParams: { id: data.ID } });
    }
  }

}
