import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { NotFoundComponent } from './other/not-found/not-found.component';


const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('./blog/blog.module').then((m) => m.BlogModule)
  },
  { path: 'admin', loadChildren: () => import('./admin/admin.module').then((m) => m.AdminModule) },
  { path: 'account', loadChildren: () => import('./account/account.module').then((m) => m.LoginModule) },
  { path: '**', component: NotFoundComponent },
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
