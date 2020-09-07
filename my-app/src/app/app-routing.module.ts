import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  { path: '', redirectTo: '/blog', pathMatch: 'full' },
  {
    path: 'blog',
    loadChildren: () => import('./blog/blog.module').then((m) => m.BlogModule),
    // component: BlogComponent,
    // children: [
    //   { path: '', redirectTo: '/blog/home', pathMatch: 'full' },
    //   { path: 'home', component: HomeComponent },
    //   { path: 'detail/:id', component: DetailComponent },
    // ],
  },
  { path: 'admin', loadChildren: () => import('./admin/admin.module').then((m) => m.AdminModule) },
  { path: 'login', loadChildren: () => import('./login/login.module').then((m) => m.LoginModule) },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      initialNavigation: 'enabled',
    }),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule { }
