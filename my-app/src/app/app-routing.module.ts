import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BlogComponent } from './blog/blog.component';
import { AdminComponent } from './admin/admin.component';
import { HomeComponent } from './blog/home/home.component';
import { DetailComponent } from './blog/detail/detail.component';

const routes: Routes = [
  { path: '', redirectTo: '/blog', pathMatch: 'full' },
  {
    path: 'blog',
    component: BlogComponent,
    children: [
      { path: '', redirectTo: '/blog/home', pathMatch: 'full' },
      { path: 'home', component: HomeComponent },
      { path: 'detail/:id', component: DetailComponent },
    ],
  },
  { path: 'admin', component: AdminComponent },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      initialNavigation: 'enabled',
    }),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule {}
