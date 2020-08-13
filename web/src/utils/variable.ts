/*
 * @Author: wangzhongjie
 * @Date: 2020-05-14 09:27:02
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-08-12 14:34:00
 * @Description:主题色变量
 * @Email: UvDream@163.com
 */

// 字体变量
const baseSize = {
  '--font-size-large-x': '22px',
  '--font-size-large': '18px',
  '--font-size-medium': '14px',
  '--font-size-medium-x': '16px',
  '--font-size-small-s': '10px',
  '--font-size-small': '12px',
};

//浅色
export const lightTheme = {
  // 主题色
  '--themeColor': '#007CFF',
  // 背景色
  '--bgColor': ' rgb(239,239,239)',
  '--bgColor1': 'rgb(249,249,249)',
  '--bgColor2': '#fff',
  // 文字颜色
  '--textColor': '#777',
  '--textColor1': '#777',
  ...baseSize,
};

// 深色
export const darkTheme = {
  '--themeColor': '#007CFF',
  '--bgColor': '#212121',
  '--bgColor1': '#212121',
  '--textColor': ' #777',
  '--textColor1': ' #777',
  ...baseSize,
};
