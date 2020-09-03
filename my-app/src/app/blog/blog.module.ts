import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BlogRoutingModule } from './blog-routing.module';
import { BlogComponent } from './blog.component';
import { HomeComponent } from './home/home.component';
import { DetailComponent } from './detail/detail.component';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { ListComponent } from './home/component/list/list.component';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { RightContentComponent } from './home/component/right-content/right-content.component';
@NgModule({
  imports: [CommonModule, BlogRoutingModule, NzButtonModule, NzGridModule],
  exports: [BlogComponent],
  declarations: [
    BlogComponent,
    HomeComponent,
    DetailComponent,
    ListComponent,
    RightContentComponent,
  ],
  providers: [],
})
export class BlogModule {}
