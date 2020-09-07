import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdminComponent } from './admin.component';
export const adminRoutes: Routes = [
    {
        path: '',
        component: AdminComponent,

    },
];

@NgModule({
    imports: [RouterModule.forChild(adminRoutes)],
    exports: [RouterModule],
})
export class AdminRoutingModule { }
