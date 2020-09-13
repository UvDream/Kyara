/*
 * @Author: wangzhongjie
 * @Date: 2020-09-02 15:19:45
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-09-13 21:19:27
 * @Description:工具函数
 * @Email: UvDream@163.com
 */
export const loadStyle = (url: any) => {
  const link = document.createElement('link');
  link.type = 'text/css';
  link.rel = 'stylesheet';
  link.href = url;
  const head = document.getElementsByTagName('head')[0];
  head.appendChild(link);
};
export function getBase64(file: File): Promise<string | ArrayBuffer | null> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}
