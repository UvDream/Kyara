import { Component, OnInit } from '@angular/core';
import * as echarts from 'echarts';


@Component({
  selector: 'app-dynamic-chart',
  templateUrl: './dynamic-chart.component.html',
  styleUrls: ['./dynamic-chart.component.less']
})
export class DynamicChartComponent implements OnInit {
  options: any;
  constructor() { }

  ngOnInit(): void {
    this.options = {
      title: {
        text: '动态日历'
      },
      visualMap: {
        min: 0,
        max: 10000,
        inRange: {
          color: ['#ebedf0', '#c6e48b', '#7bc96f', '#239a3b', '#196027']
        },
        show: false
      },
      tooltip: {
        position: 'top',
        formatter: (params) => {
          return params.value[0] + ',' + params.value[1];
        }
      },
      grid: {
        top: '0%',
        left: '0%'
      },
      backgroundColor: '#fff',
      calendar: {
        cellSize: [14, 14],
        range: [this.getVirtulData().thatday, this.getVirtulData().today],
        itemStyle: {
          borderColor: '#fff',
          borderWidth: 4
        },
        splitLine: {
          show: false
        },
        monthLabel: {
          formatter: '{M}月'
        },
        dayLabel: {
          nameMap: 'cn'
        },
        yearLabel: { show: false }
      },
      series: {
        type: 'heatmap',
        coordinateSystem: 'calendar',
        data: this.getVirtulData().data
      }
    };
  }
  getVirtulData(): any {
    const today = echarts.number.parseDate(new Date());
    const dayTime = 3600 * 24 * 1000;
    const thatday = today - dayTime * 365;

    const data = [];
    for (let time = thatday; time < today; time += dayTime) {
      data.push([
        echarts.format.formatTime('yyyy-MM-dd', time),
        Math.floor(Math.random() * 10000)
      ]);
    }
    console.log(data)
    return {
      data,
      today: echarts.format.formatTime('yyyy-MM-dd', today),
      thatday: echarts.format.formatTime('yyyy-MM-dd', thatday)
    };
  }

}
