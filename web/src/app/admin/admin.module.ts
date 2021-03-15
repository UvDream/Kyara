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
import { NzTreeModule } from 'ng-zorro-antd/tree';
import { NzTabsModule } from 'ng-zorro-antd/tabs';
import { UserConfigComponent } from './blog-config/user-config/user-config.component';
import { OtherConfigComponent } from './blog-config/other-config/other-config.component';
import { BlogAboutComponent } from './blog-config/blog-about/blog-about.component';
import { NzPopoverModule } from 'ng-zorro-antd/popover';
import { NzDrawerModule } from 'ng-zorro-antd/drawer';
import { ArticleConfigComponent } from './blog-config/article-config/article-config.component';
import { NzImageModule } from 'ng-zorro-antd/image';
import { MarkdownEditModule } from '../components/markdown-edit/markdown-edit.module';
import { AddQuestionComponent } from './add-question/add-question.component';
import { QuestionClassifyComponent } from './question-classify/question-classify.component';
import { QuestionListComponent } from './question-list/question-list.component';
import { NzCollapseModule } from 'ng-zorro-antd/collapse';
import { MarkdownPreviewPipe } from '@pipes/markdown-preview.pipe';
import { PipeModule } from '@pipes/pipe.module';



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
    NzPopconfirmModule,
    NzTreeModule,
    NzTabsModule,
    NzPopoverModule,
    NzDrawerModule,
    NzImageModule,
    MarkdownEditModule,
    NzCollapseModule,
    PipeModule
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
    UserConfigComponent,
    OtherConfigComponent,
    BlogAboutComponent,
    ArticleConfigComponent,
    AddQuestionComponent,
    QuestionClassifyComponent,
    QuestionListComponent,
    // MarkdownEditComponent,
    // MarkdownPreviewPipe
  ],
  exports: [AdminComponent],
})
export class AdminModule { }
