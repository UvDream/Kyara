import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BlogComponent } from './blog/blog.component';
import { AdminComponent } from './admin/admin.component';
import { HomeComponent } from './blog/home/home.component';
import { DetailComponent } from './blog/detail/detail.component';
@NgModule({
  declarations: [
    AppComponent,
    BlogComponent,
    AdminComponent,
    HomeComponent,
    DetailComponent,
  ],
  imports: [
    BrowserModule.withServerTransition({ appId: 'serverApp' }),
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
