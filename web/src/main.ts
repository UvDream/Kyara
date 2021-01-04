import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';
import { environment } from './environments/environment';
import { loadStyle } from './util/util';
import { iconfontUrl, iconfontVersion } from './environments/env';
if (environment.production) {
  enableProdMode();
}
iconfontVersion.forEach((ele: any) => {
  ele.icon ? loadStyle(iconfontUrl.replace('$key', ele.icon), false) :
    loadStyle(iconfontUrl.replace('$key', ele.svg), false);
});
iconfontVersion.forEach((ele: any) => {
  if (ele.svg) {
    loadStyle(iconfontUrl.replace('$key', ele.svg), true);
  }
});
document.addEventListener('DOMContentLoaded', () => {
  platformBrowserDynamic()
    .bootstrapModule(AppModule)
    .catch((err) => console.error(err));
});
