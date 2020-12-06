import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AdminComponent } from './admin.component';
import { AdminRoutingModule } from './admin-routing.module';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { EditArticleComponent } from './edit-article/edit-article.component';
import { NzInputModule } from 'ng-zorro-antd/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzTreeSelectModule } from 'ng-zorro-antd/tree-select';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { ArticleListComponent } from './article-list/article-list.component';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NoticeComponent } from './notice/notice.component';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzFormModule } from 'ng-zorro-antd/form';
import { ClassifyComponent } from './classify/classify.component';
import { TagsComponent } from './tags/tags.component';
import { NzTagModule } from 'ng-zorro-antd/tag';
import { CommentComponent } from './comment/comment.component';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { BlogConfigComponent } from './blog-config/blog-config.component';
import { UploadImageComponent } from './components/upload-image/upload-image.component';
import { ImagesComponent } from './images/images.component';

@NgModule({
  imports: [
    CommonModule,
    AdminRoutingModule,
    NzLayoutModule,
    NzMenuModule,
    NzInputModule,
    FormsModule,
    NzSelectModule,
    NzUploadModule,
    NzCheckboxModule,
    NzButtonModule,
    NzTreeSelectModule,
    NzIconModule,
    NzMessageModule,
    NzTableModule,
    NzModalModule,
    NzFormModule,
    ReactiveFormsModule,
    NzTagModule,
    NzPopconfirmModule
  ],
  declarations: [
    AdminComponent,
    EditArticleComponent,
    ArticleListComponent,
    NoticeComponent,
    ClassifyComponent,
    TagsComponent,
    CommentComponent,
    BlogConfigComponent,
    UploadImageComponent,
    ImagesComponent,
  ],
  exports: [AdminComponent],
})
export class AdminModule { }
