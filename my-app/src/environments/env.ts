/*
 * @Author: wangzhongjie
 * @Date: 2020-04-14 13:56:51
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-09-07 15:44:50
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
    baseUrl = `http://118.25.110.93`;
}
export { iconfontUrl, iconfontVersion, baseUrl };
