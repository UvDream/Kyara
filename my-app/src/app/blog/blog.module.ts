import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BlogRoutingModule } from './blog-routing.module';
import { BlogComponent } from './blog.component';
import { HomeComponent } from './home/home.component';
import { DetailComponent } from './detail/detail.component';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { ListComponent } from './home/list/list.component';
@NgModule({
  imports: [CommonModule, BlogRoutingModule, NzButtonModule],
  exports: [BlogComponent],
  declarations: [BlogComponent, HomeComponent, DetailComponent, ListComponent],
  providers: [],
})
export class BlogModule {}