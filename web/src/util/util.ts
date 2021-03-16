/*
 * @Author: wangzhongjie
 * @Date: 2020-09-02 15:19:45
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2021-03-16 10:59:24
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
export function toTimeDH(timeDate: string, type?: string): string {
  const time = new Date(timeDate).getTime();
  const nowTime = new Date().getTime();
  const diffTime = nowTime - time;
  const day = diffTime / 1000 / 60 / 60 / 24;
  const dayRound = Math.floor(day);
  const hours = diffTime / 1000 / 60 / 60 - (24 * dayRound);
  const hoursRound = Math.floor(hours);
  if (type === 'DD') {
    if (day <= 1) {
      return '一天内';
    } else {
      return dayRound + '天前';
    }
  } else if (type === 'DD-HH') {
    return dayRound + '天' + hoursRound + '小时';
  } else {
    return diffTime.toString();
  }

}
/**
 * @description: 复制代码
 */
export function CopyText(): void {
  console.log('拷贝代码');
  const copyText = document.getElementsByClassName('copyBtn');
  // tslint:disable-next-line:prefer-for-of
  for (let i = 0; i < copyText.length; i++) {
    const element = copyText[i] as HTMLElement;
    element.onclick = (e: any) => {
      const oInput = document.createElement('textarea');
      oInput.value = e.target.children[0].innerText;
      document.body.appendChild(oInput);
      oInput.select();
      document.execCommand('Copy');
      oInput.className = 'oInput';
      oInput.style.display = 'none';
    };
  }

}
export function downloadMD(doc: any, name: string): void {
  const downloadLink = document.createElement('a');
  downloadLink.download = name + '.md';
  downloadLink.style.display = 'none';
  const blob = new Blob([doc]);
  downloadLink.href = URL.createObjectURL(blob);
  document.body.appendChild(downloadLink);
  downloadLink.click();
  document.body.removeChild(downloadLink);
}
