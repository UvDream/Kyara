import { Component, OnInit } from '@angular/core';
import * as echarts from 'echarts';
import { ArticleService } from '@service/article.service';

@Component({
  selector: 'app-dynamic-chart',
  templateUrl: './dynamic-chart.component.html',
  styleUrls: ['./dynamic-chart.component.less']
})
export class DynamicChartComponent implements OnInit {
  options: any;
  loading = true;
  constructor(
    private httpArticle: ArticleService
  ) { }

  ngOnInit(): void {
    this.getData();
  }

  async getData(): Promise<void> {
    this.loading = true;
    const res = await this.httpArticle.getBlogDynamic();
    if (res.code === 200) {
      this.options = {
        title: {
          text: '动态日历'
        },
        visualMap: {
          min: 0,
          max: res.data.maxCount,
          inRange: {
            color: ['#ebedf0', '#c6e48b', '#7bc96f', '#239a3b', '#196027']
          },
          show: false
        },
        tooltip: {
          position: 'top',
          formatter: (params) => {
            return params.value[0] + ':' + params.value[1];
          }
        },
        grid: {
          top: '0%',
          left: '0%'
        },
        backgroundColor: '#fff',
        calendar: {
          cellSize: [14, 14],
          range: [res.data.thatDay, res.data.today],
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
          data: res.data.list
        }
      };
      this.loading = false;
    }
  }

}
