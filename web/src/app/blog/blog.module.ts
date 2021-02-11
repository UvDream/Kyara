import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
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
import { NzTabsModule } from 'ng-zorro-antd/tabs';
import { BlogMessageComponent } from './home/component/right-content/blog-message/blog-message.component';
import { TagsComponent } from './home/component/right-content/tags/tags.component';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzPaginationModule } from 'ng-zorro-antd/pagination';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { CatalogComponent } from './home/component/right-content/catalog/catalog.component';
import { NzAnchorModule } from 'ng-zorro-antd/anchor';
import { HotArticleComponent } from './home/component/right-content/hot-article/hot-article.component';
import { NumberFormatPipe } from '../pipes/number-format.pipe';
import { GithubComponent } from './github/github.component';
import { ArchivesComponent } from './archives/archives.component';
import { PhotosComponent } from './photos/photos.component';
import { AboutComponent } from './about/about.component';
import { NzSpinModule } from 'ng-zorro-antd/spin';
import { NzTimelineModule } from 'ng-zorro-antd/timeline';
import { NzDrawerModule } from 'ng-zorro-antd/drawer';
import { NzSkeletonModule } from 'ng-zorro-antd/skeleton';
import { BackTopComponent } from './components/back-top/back-top.component';
import { NzBackTopModule } from 'ng-zorro-antd/back-top';
import { NoticeComponent } from './components/notice/notice.component';
import { UpdateLogComponent } from './about/update-log/update-log.component';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzSwitchModule } from 'ng-zorro-antd/switch';
import { CommentsComponent } from './comments/comments.component';
import { CommentComponent } from './components/comment/comment.component';
import { CommentFormComponent } from './components/comment-form/comment-form.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { ChartsComponent } from './home/component/charts/charts.component';
import { DynamicChartComponent } from './home/component/charts/dynamic-chart/dynamic-chart.component';
import { RadarChartComponent } from './home/component/charts/radar-chart/radar-chart.component';
import { NzImageModule } from 'ng-zorro-antd/image';
import { TimeDiffPipe } from '@pipes/time-diff.pipe';
import { UserInfoComponent } from './home/component/user-info/user-info.component';
import { LoadingComponent } from './components/loading/loading.component';
import { MusicComponent } from './home/component/music/music.component';
import { MarkdownPreviewComponent } from '../components/markdown-preview/markdown-preview.component';

@NgModule({
  imports: [
    CommonModule,
    BlogRoutingModule,
    NzButtonModule,
    NzGridModule,
    NzTabsModule,
    NzDropDownModule,
    NzModalModule,
    NzPaginationModule,
    NzInputModule,
    FormsModule,
    NzMessageModule,
    NzIconModule,
    NzAnchorModule,
    NzSpinModule,
    NzTimelineModule,
    NzDrawerModule,
    NzSkeletonModule,
    NzBackTopModule,
    ReactiveFormsModule,
    NzFormModule,
    NzSwitchModule,
    NzImageModule,
    NgxEchartsModule.forRoot({
      echarts: () => import('echarts')
    })
  ],
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
    BlogMessageComponent,
    TagsComponent,
    CatalogComponent,
    HotArticleComponent,
    NumberFormatPipe,
    GithubComponent,
    ArchivesComponent,
    PhotosComponent,
    AboutComponent,
    BackTopComponent,
    NoticeComponent,
    UpdateLogComponent,
    CommentsComponent,
    CommentComponent,
    CommentFormComponent,
    ChartsComponent,
    DynamicChartComponent,
    RadarChartComponent,
    TimeDiffPipe,
    UserInfoComponent,
    LoadingComponent,
    MusicComponent,
    MarkdownPreviewComponent
  ],
  providers: [],
})
export class BlogModule { }
