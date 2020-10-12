import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'numberFormat'
})
export class NumberFormatPipe implements PipeTransform {

  transform(value: any, args?: any): string {
    if (value === '') {
      return '暂无';
    }
    if (value > 999) {
      return '999+';
    }
    return value;
  }

}
