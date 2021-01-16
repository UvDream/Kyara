import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { AccountRoutes } from './account.routing';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { RegisterComponent } from './register/register.component';
@NgModule({
  imports: [
    CommonModule,
    AccountRoutes,
    NzFormModule,
    NzInputModule,
    NzGridModule,
    NzButtonModule,
    FormsModule,
    ReactiveFormsModule,
    NzMessageModule
  ],
  declarations: [LoginComponent, RegisterComponent],
  exports: [LoginComponent]
})
export class LoginModule { }
