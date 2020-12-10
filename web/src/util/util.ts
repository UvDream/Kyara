/*
 * @Author: wangzhongjie
 * @Date: 2020-09-02 15:19:45
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-12-10 16:58:05
 * @Description:工具函数
 * @Email: UvDream@163.com
 */

/**
 * @description: 添加iconfont
 * @param 连接url url
 * @param 是否是svg svg
 * @return 无返回值
 */
export const loadStyle = (url: string, svg?: boolean) => {
  const head = document.getElementsByTagName('head')[0];
  if (!svg) {
    const link = document.createElement('link');
    link.type = 'text/css';
    link.rel = 'stylesheet';
    link.href = url + '.css';
    head.appendChild(link);
  } else {
    const script = document.createElement('script');
    script.type = 'text/javascript';
    script.src = url + '.js';
    head.appendChild(script);
  }

};

export function getBase64(file: File): Promise<string | ArrayBuffer | null> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}
export function toTimeDH(timeDate: string, type: string): string {
  const time = new Date(timeDate).getTime();
  const nowTime = new Date().getTime();
  const diffTime = nowTime - time;
  const day = diffTime / 1000 / 60 / 60 / 24;
  const dayRound = Math.floor(day);
  const hours = diffTime / 1000 / 60 / 60 - (24 * dayRound);
  const hoursRound = Math.floor(hours);
  if (type === 'DD') {
    return dayRound + '天前';
  } else if (type === 'DD-HH') {
    return dayRound + '天' + hoursRound + '小时';
  }

}
