/*
 * @Author: wangzhongjie
 * @Date: 2020-04-14 13:56:51
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-12-10 17:01:23
 * @Description:配置文件
 * @Email: UvDream@163.com
 */
import { environment } from './environment';
const iconfontVersion = [{ icon: '1762601_ytyd1qlptvq' }, { svg: '1753589_tnfpi591qt' }];
const iconfontUrl = `//at.alicdn.com/t/font_$key`;

let baseUrl: string;
if (!environment.production) {
    baseUrl = `http://localhost:3000`;
} else {
    baseUrl = `http://www.uvdream.cn/api`;
}
export { iconfontUrl, iconfontVersion, baseUrl };
