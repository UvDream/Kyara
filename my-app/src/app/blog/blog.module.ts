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
import { LeftMenusComponent } from './home/component/left-menus/left-menus.component';
import { HeaderComponent } from './home/component/header/header.component';
import { FooterComponent } from './home/component/footer/footer.component';
import { AvatarComponent } from './home/component/left-menus/avatar/avatar.component';
import { NavigationComponent } from './home/component/left-menus/navigation/navigation.component';
import { LeftFooterComponent } from './home/component/left-menus/left-footer/left-footer.component';
@NgModule({
  imports: [CommonModule, BlogRoutingModule, NzButtonModule, NzGridModule],
  exports: [BlogComponent],
  declarations: [
    BlogComponent,
    HomeComponent,
    DetailComponent,
    ListComponent,
    RightContentComponent,
    LeftMenusComponent,
    HeaderComponent,
    FooterComponent,
    AvatarComponent,
    NavigationComponent,
    LeftFooterComponent,
  ],
  providers: [],
})
export class BlogModule {}
