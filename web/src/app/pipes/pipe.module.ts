import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MarkdownPreviewPipe } from './markdown-preview.pipe';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [
    MarkdownPreviewPipe
  ],
  declarations: [MarkdownPreviewPipe]
})


export class PipeModule {
  // tslint:disable-next-line:typedef
  static forRoot() {
    return {
      ngModule: PipeModule,
      providers: [],
    };
  }
}
