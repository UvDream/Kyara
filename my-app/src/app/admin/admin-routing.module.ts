import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdminComponent } from './admin.component';
import { EditArticleComponent } from './edit-article/edit-article.component';
import { ArticleListComponent } from './article-list/article-list.component';

export const adminRoutes: Routes = [
  {
    path: '',
    component: AdminComponent,
    children: [
      { path: '', redirectTo: '/editArticle', pathMatch: 'full' },
      { path: 'editArticle', component: EditArticleComponent },
      { path: 'articleList', component: ArticleListComponent },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(adminRoutes)],
  exports: [RouterModule],
})
export class AdminRoutingModule {}
