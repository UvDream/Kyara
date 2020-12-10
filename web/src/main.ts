import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';
import { environment } from './environments/environment';
import { loadStyle, loadScript } from './util/util';
import { iconfontUrl, iconfontVersion } from './environments/env';
if (environment.production) {
  enableProdMode();
}
iconfontVersion.forEach((ele: any, index: number) => {
  loadStyle(iconfontUrl.replace('$key', ele.icon));
  loadStyle(iconfontUrl.replace('$key', ele.svg));
});
iconfontVersion.forEach((ele: any) => {
  loadScript(iconfontUrl.replace('$key', ele.svg));
});
document.addEventListener('DOMContentLoaded', () => {
  platformBrowserDynamic()
    .bootstrapModule(AppModule)
    .catch((err) => console.error(err));
});
