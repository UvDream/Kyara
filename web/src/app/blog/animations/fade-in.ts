/*
 * @Author: wangzhongjie
 * @Date: 2020-11-16 11:30:43
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2021-02-04 13:46:53
 * @Description: 渐入动画
 * @Email: UvDream@163.com
 */
import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';
export const fadeIn = trigger('fadeIn', [
    state('in', style({ transform: 'translateX(0)' })),
    transition('void=>*', [
        animate(1000, keyframes([
            style({ opacity: 0, transform: 'translateY(-100%)' }),
            style({ opacity: 1, transform: 'translateY(-50%)' }),
            style({ opacity: 1, transform: 'translateY(0%)' })
        ]))
    ]),
    transition('*=>void', [
        animate(1000, keyframes([
            style({ transform: 'translateY(0%)' }),
            style({ transform: 'translateY(-50%)' }),
            style({ transform: 'translateY(-100%)' })
        ]))
    ])
]);
