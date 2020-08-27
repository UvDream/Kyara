import axios, { AxiosResponse, AxiosRequestConfig } from 'axios';
import config from './config';

declare type Methods =
  | 'GET'
  | 'OPTIONS'
  | 'HEAD'
  | 'POST'
  | 'PUT'
  | 'DELETE'
  | 'TRACE'
  | 'CONNECT';

class HttpRequest {
  public queue: any;
  public constructor() {
    this.queue = {};
  }
  destroy(url: string) {
    delete this.queue[url];
    if (!Object.keys(this.queue).length) {
      // hide loading
    }
  }
  interceptors(instance: any, url?: string) {
    // 请求拦截
    instance.interceptors.request.use(
      (config: AxiosRequestConfig) => {
        // 添加全局的loading...
        if (!Object.keys(this.queue).length) {
        }
        if (url) {
          this.queue[url] = true;
        }
        return config;
      },
      (error: any) => {
        return Promise.reject(error);
      },
    );
    // 响应拦截
    instance.interceptors.response.use(
      (res: AxiosResponse) => {
        let data;
        if (res.data == undefined) {
          data = res.request.responseText;
        } else {
          data = res.data;
        }
        switch (data.code) {
          case '110':
            break;
          default:
        }
        return data;
      },
      (error: any) => {
        if (url) {
          this.destroy(url);
        }
        if (error && error.request) {
          let status = error.request.status;
          switch (status) {
            case 401:
              break;
            case 404:
              break;
            case 415:
              break;
            case 500:
              break;
            default:
              break;
          }
        }
        return Promise.reject(error);
      },
    );
  }
  async request(options: AxiosRequestConfig) {
    const instance = axios.create();
    options = Object.assign(config, options);
    await this.interceptors(instance, options.url);
    return instance(options);
  }
}

const $axios = new HttpRequest();

export default $axios as any;
