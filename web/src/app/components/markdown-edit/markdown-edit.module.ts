import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MarkdownEditComponent } from './markdown-edit.component';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [MarkdownEditComponent],
  declarations: [MarkdownEditComponent]
})
export class MarkdownEditModule { }
