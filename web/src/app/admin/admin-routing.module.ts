import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdminComponent } from './admin.component';
import { EditArticleComponent } from './edit-article/edit-article.component';
import { ArticleListComponent } from './article-list/article-list.component';
import { NoticeComponent } from './notice/notice.component';
import { ClassifyComponent } from './classify/classify.component';
import { TagsComponent } from './tags/tags.component';
import { CommentComponent } from './comment/comment.component';
import { BlogConfigComponent } from './blog-config/blog-config.component';
import { ImagesComponent } from './images/images.component';
import { AddQuestionComponent } from './add-question/add-question.component';
import { QuestionClassifyComponent } from './question-classify/question-classify.component';
import { QuestionListComponent } from './question-list/question-list.component';


export const adminRoutes: Routes = [
  { path: '', redirectTo: '/admin/editArticle', pathMatch: 'full' },
  {
    path: '',
    component: AdminComponent,
    children: [
      { path: 'editArticle', component: EditArticleComponent },
      { path: 'articleList', component: ArticleListComponent },
      { path: 'notice', component: NoticeComponent },
      { path: 'tags', component: TagsComponent },
      { path: 'classify', component: ClassifyComponent },
      { path: 'commentList', component: CommentComponent },
      { path: 'blogConfig', component: BlogConfigComponent },
      { path: 'images', component: ImagesComponent },
      { path: 'addQuestion', component: AddQuestionComponent },
      { path: 'questionList', component: QuestionListComponent },
      { path: 'questionClassify', component: QuestionClassifyComponent },
    ],
  },

];

@NgModule({
  imports: [RouterModule.forChild(adminRoutes)],
  exports: [RouterModule],
})
export class AdminRoutingModule { }
