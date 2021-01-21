/*
 * @Author: wangzhongjie
 * @Date: 2020-04-14 13:56:51
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2021-01-21 16:50:05
 * @Description:配置文件
 * @Email: UvDream@163.com
 */
import { environment } from './environment';
const iconfontVersion = [{ icon: '1762601_ytyd1qlptvq' }, { svg: '1753589_p1dg92w9of' }, { svg: '2341925_u0kxv65a6rj' }];
const iconfontUrl = `//at.alicdn.com/t/font_$key`;

let baseUrl: string;
if (!environment.production) {
    baseUrl = `http://localhost:3000`;
} else {
    baseUrl = `http://www.uvdream.cn/api`;
}
export { iconfontUrl, iconfontVersion, baseUrl };
