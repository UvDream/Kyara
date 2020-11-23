/*
 * @Author: wangzhongjie
 * @Date: 2020-04-14 13:56:51
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-11-13 11:23:52
 * @Description:配置文件
 * @Email: UvDream@163.com
 */
import { environment } from './environment';
const iconfontVersion = ['1762601_ytyd1qlptvq'];
const iconfontUrl = `//at.alicdn.com/t/font_$key.css`;
let baseUrl: any;
if (!environment.production) {
    baseUrl = `http://localhost:3000`;
} else {
    baseUrl = `http://www.uvdream.cn/api`;
}
export { iconfontUrl, iconfontVersion, baseUrl };
