/*
 * @Author: wangzhongjie
 * @Date: 2020-11-16 11:30:43
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2020-11-16 14:37:12
 * @Description: 渐入动画
 * @Email: UvDream@163.com
 */
import { trigger, state, style, transition, animate, keyframes } from '@angular/animations';
export const fadeIn = trigger('fadeIn', [
    state('in', style({ transform: 'translateX(0)' })),
    transition('void => *', [
        animate(300, keyframes([
            style({ opacity: 0, transform: 'translateX(-100%)', offset: 0 }),
            style({ opacity: 1, transform: 'translateX(25px)', offset: 0.3 }),
            style({ opacity: 1, transform: 'translateX(0)', offset: 1.0 })
        ]))
    ]),
    transition('* => void', [
        animate(300, keyframes([
            style({ opacity: 1, transform: 'translateX(0)', offset: 0 }),
            style({ opacity: 1, transform: 'translateX(-25px)', offset: 0.7 }),
            style({ opacity: 0, transform: 'translateX(100%)', offset: 1.0 })
        ]))
    ])
]);