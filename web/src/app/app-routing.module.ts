import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('./blog/blog.module').then((m) => m.BlogModule)
  },
  { path: 'admin', loadChildren: () => import('./admin/admin.module').then((m) => m.AdminModule) },
  { path: 'account', loadChildren: () => import('./account/login.module').then((m) => m.LoginModule) },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      initialNavigation: 'enabled',
      relativeLinkResolution: 'legacy',
      useHash: true
    }),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule { }
