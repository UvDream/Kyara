import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.less'],
})
export class ListComponent implements OnInit {
  constructor(private router: Router) { }
  @Input() data: any;
  ngOnInit(): void { }
  // tslint:disable-next-line:typedef
  toDetail(data) {
    if (data.is_password === '1') {
      console.log('需要密码');
    } else {
      this.router.navigate(['/detail'], { queryParams: { id: data.ID } });
    }
  }
}
