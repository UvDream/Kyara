import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'TimeDiffPipe'
})
export class TimeDiffPipe implements PipeTransform {

  transform(value: any, args?: any): any {
    let timeVal = '';
    const time = new Date(value).getTime();
    const nowTime = new Date().getTime();
    const diffTime = (nowTime - time) / 1000; // 转换为秒
    // 先计算出年份
    const years = parseInt((diffTime / 31536000).toString(), 0);
    years > 0 ? timeVal = timeVal + years + '年' : '';
    // 计算月
    const monthsDiffTime = diffTime - years * 31536000;
    const months = parseInt((monthsDiffTime / 2626560).toString(), 0);
    months > 0 ? timeVal = timeVal + months + '个月' : '';
    // 计算日
    const dayDiffTime = monthsDiffTime - months * 2626560;
    const days = parseInt((dayDiffTime / 86400).toString(), 0);
    days > 0 ? timeVal = timeVal + days + '天' : '';
    if (years === 0) {
      const hDiffTime = dayDiffTime - days * 86400;
      const hours = parseInt((hDiffTime / 3600).toString(), 0);
      timeVal = timeVal + hours + '小时';
    }
    return timeVal;
  }

}
