import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BlogRoutingModule } from './blog-routing.module';
import { BlogComponent } from './blog.component';
import { HomeComponent } from './home/home.component';
import { DetailComponent } from './detail/detail.component';

@NgModule({
  imports: [CommonModule, BlogRoutingModule],
  exports: [BlogComponent],
  declarations: [BlogComponent, HomeComponent, DetailComponent],
  providers: [],
})
export class BlogModule {}
