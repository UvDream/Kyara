import { RouteReuseStrategy, ActivatedRouteSnapshot, DetachedRouteHandle } from '@angular/router';

export class AppRouteReuseStrategy implements RouteReuseStrategy {

  public static handlers: { [key: string]: DetachedRouteHandle } = {};
  private static waitDelete: string;
  // 删除
  public static deleteRouteSnapshot(name: string): void {
    if (AppRouteReuseStrategy.handlers[name]) {
      delete AppRouteReuseStrategy.handlers[name];
    } else {
      AppRouteReuseStrategy.waitDelete = name;
    }
  }
  /** 表示对所有路由允许复用 如果你有路由不想利用可以在这加一些业务逻辑判断 */
  public shouldDetach(route: ActivatedRouteSnapshot): boolean {
    // 添加控制器过滤 false则不对其缓存，反之...
    if (!route.data.keep) {
      return false;
    }
    if (!route.routeConfig || route.routeConfig.loadChildren) {
      return false;
    }
    return true;
  }

  /** 当路由离开时会触发。按path作为key存储路由快照&组件当前实例对象 */
  public store(route: ActivatedRouteSnapshot, handle: DetachedRouteHandle): void {
    if (AppRouteReuseStrategy.waitDelete && AppRouteReuseStrategy.waitDelete === route.routeConfig.path) {
      // 如果待删除是当前路由则不存储快照
      AppRouteReuseStrategy.waitDelete = null;
      return;
    }
    AppRouteReuseStrategy.handlers[route.routeConfig.path] = handle;
  }

  /** 若 path 在缓存中有的都认为允许还原路由 */
  public shouldAttach(route: ActivatedRouteSnapshot): boolean {
    return !!route.routeConfig && !!AppRouteReuseStrategy.handlers[route.routeConfig.path];
  }

  /** 从缓存中获取快照，若无则返回null */
  public retrieve(route: ActivatedRouteSnapshot): DetachedRouteHandle {
    if (!route.routeConfig) { return null; }
    return AppRouteReuseStrategy.handlers[route.routeConfig.path];
  }

  /** 进入路由触发，判断是否同一路由 */
  public shouldReuseRoute(future: ActivatedRouteSnapshot, curr: ActivatedRouteSnapshot): boolean {
    return future.routeConfig === curr.routeConfig;
  }
}
