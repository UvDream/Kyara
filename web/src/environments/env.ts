/*
 * @Author: wangzhongjie
 * @Date: 2020-04-14 13:56:51
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2021-03-08 15:17:15
 * @Description:配置文件
 * @Email: UvDream@163.com
 */
import { environment } from './environment';
const iconfontVersion = [
    { icon: '1762601_ytyd1qlptvq' },
    // 彩色分类icon
    { svg: '1753589_n52485ovn5' },
    //  博客彩色图标
    { svg: '2341925_vc6wx8gzcja' }];
const iconfontUrl = `//at.alicdn.com/t/font_$key`;

let baseUrl: string;
if (!environment.production) {
    baseUrl = `http://localhost:3000`;
} else {
    baseUrl = `http://www.uvdream.cn/api`;
}
export { iconfontUrl, iconfontVersion, baseUrl };
