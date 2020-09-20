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
    NzTreeSelectModule
  ],
  declarations: [AdminComponent, EditArticleComponent, UploadComponent],
  exports: [AdminComponent],
})
export class AdminModule { }
