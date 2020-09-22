import { Injectable } from '@angular/core';
import { baseUrl } from '../../environments/env';
import { HttpClient } from '@angular/common/http';
import { Params } from './params';
interface Response {
  code: number;
  data: any;
  msg: string;
}
@Injectable({
  providedIn: 'root'
})

export class HttpService {

  constructor(private http: HttpClient) { }
  public request(params: any): any {
    if (params.method === 'post' || params.method === 'POST') {
      return this.post(params.url, params.data);
    } else {
      return this.get(params.url, params.params);
    }
  }
  public get = (url: string, params: any) => {

    return this.http.get(baseUrl + url, { params }).toPromise().then(this.handleSuccess).catch(this.handleError);
  }
  public post = (url: string, data: Params) => {
    return this.http.post(baseUrl + url, data).toPromise().then(this.handleSuccess).catch(this.handleError);
  }
  private handleSuccess = (res: Response) => {
    switch (res.code) {
      case 200:
        return res;
      case 400:
        return res;
      default:
        return;
    }
  }
  private handleError = (error: any) => {
    if (error.status === 400) {
      console.error('请求参数错误');
    } else if (error.status === 500) {
      console.error('服务器内部错误');
    } else if (error.status === 404) {
      console.error('url找不到');
    }
    return {
      code: 400,
      msg: '请求错误'
    };
  }

}
