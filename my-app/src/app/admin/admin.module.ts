import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AdminComponent } from './admin.component';
import { AdminRoutingModule } from './admin-routing.module';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { EditArticleComponent } from './edit-article/edit-article.component';
@NgModule({
  imports: [CommonModule, AdminRoutingModule, NzLayoutModule, NzMenuModule],
  declarations: [AdminComponent, EditArticleComponent],
  exports: [AdminComponent],
})
export class AdminModule {}
