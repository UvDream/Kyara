import { Injectable, Inject, PLATFORM_ID } from '@angular/core';
import { baseUrl } from '../../environments/env';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Params } from './params';
import { isPlatformBrowser } from '@angular/common';
import { Router } from '@angular/router';

interface Response {
  code: number;
  data: any;
  msg: string;
}
@Injectable({
  providedIn: 'root'
})

export class HttpService {
  public httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
      Authorization: ''
    })
  };
  constructor(
    private http: HttpClient,
    @Inject(PLATFORM_ID) private platformId: object,
    private router: Router
  ) { }

  public request(params: any): any {
    if (params.method === 'post' || params.method === 'POST') {
      return this.post(params.url, params.data);
    } else {
      return this.get(params.url, params.params);
    }
  }
  public get = (url: string, params: any) => {
    const headers = this.getAuthorization();
    return this.http.get(baseUrl + url, { params, headers }).toPromise().then(this.handleSuccess).catch(this.handleError);
  }
  public post = (url: string, data: Params) => {
    const headers = this.getAuthorization();
    return this.http.post(baseUrl + url, data, { headers }).toPromise().then(this.handleSuccess).catch(this.handleError);
  }
  public getAuthorization(): HttpHeaders {
    if (isPlatformBrowser(this.platformId)) {
      const authorization = localStorage.getItem('Authorization');
      let headers = new HttpHeaders();
      if (authorization) {
        return headers = headers.set('Authorization', authorization);
      }
      return headers;
    }
  }
  private handleSuccess = (res: Response) => {
    switch (res.code) {
      case 200:
        return res;
      case 400:
        const path = this.router.url.match(/\/(\S*)\//)[1];
        if (path === 'admin') {
          this.router.navigate(['/account/login']);
        }
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
