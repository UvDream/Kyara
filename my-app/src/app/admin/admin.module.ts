import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AdminComponent } from './admin.component';
import { AdminRoutingModule } from './admin-routing.module';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { EditArticleComponent } from './edit-article/edit-article.component';
import { NzInputModule } from 'ng-zorro-antd/input';
import { FormsModule } from '@angular/forms';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { UploadComponent } from './edit-article/components/upload/upload.component';
import { NzTreeSelectModule } from 'ng-zorro-antd/tree-select';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { ArticleListComponent } from './article-list/article-list.component';
import { NzTableModule } from 'ng-zorro-antd/table';
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
    NzTableModule
  ],
  declarations: [AdminComponent, EditArticleComponent, UploadComponent, ArticleListComponent],
  exports: [AdminComponent],
})
export class AdminModule { }
