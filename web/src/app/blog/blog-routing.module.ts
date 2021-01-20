import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BlogComponent } from './blog.component';
import { HomeComponent } from './home/home.component';
import { DetailComponent } from './detail/detail.component';
import { GithubComponent } from './github/github.component';
import { PhotosComponent } from './photos/photos.component';
import { AboutComponent } from './about/about.component';
import { ArchivesComponent } from './archives/archives.component';
import { CommentsComponent } from './comments/comments.component';

export const homeRoutes: Routes = [
  {
    path: '',
    component: BlogComponent,
    children: [
      {
        path: '',
        component: HomeComponent,
      },
      {
        path: 'detail',
        data: {
          title: 'detail'
        },
        component: DetailComponent,
      },
      {
        path: 'github',
        component: GithubComponent,
      },
      {
        path: 'photos',
        component: PhotosComponent,
      },
      {
        path: 'comment',
        component: CommentsComponent,
      },
      {
        path: 'about',
        component: AboutComponent,
      },
      {
        path: 'archives',
        component: ArchivesComponent,
      },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(homeRoutes)],
  exports: [RouterModule],
})
export class BlogRoutingModule { }
