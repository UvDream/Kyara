import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
interface ClassifyItem {
  name: string;
  max: number;
  count: number;
}
@Component({
  selector: 'app-radar-chart',
  templateUrl: './radar-chart.component.html',
  styleUrls: ['./radar-chart.component.less']
})
export class RadarChartComponent implements OnInit {
  options: any;
  constructor(private articleHttp: ArticleService) { }
  data: [];
  arr: ClassifyItem[] = [];
  ngOnInit(): void {
    this.getData();
  }
  async getData(): Promise<void> {
    const res = await this.articleHttp.getClassify();
    console.log(res);
    if (res.code === 200) {
      this.data = res.data.data;
      this.arr = res.data.list;
      this.options = {
        title: {
          text: '分类雷达图'
        },
        tooltip: {
          show: true,
          padding: [5, 10],
          formatter: (params) => {
            const parm = this.arr;
            let obj = '';
            for (let i = 0; i < parm.length; i++) {
              obj = obj + '<div style="display: flex;align-items:center;justify-content:space-between;"><span>' + parm[i].name
                + '：</span><span style="margin-left:5px">' + params.data.value[i] + '</span></div>\n';
            }
            return params.seriesName + obj;
          }

        },
        radar: {
          radius: '60%',
          center: ['50%', '55%'],
          // shape: 'circle',
          splitNumber: 4,
          name: {
            textStyle: {
              color: '#999',
              borderRadius: 3,
              padding: [5, 5]
            }
          },
          indicator: this.arr
        },
        series: [{
          name: '数量',
          type: 'radar',
          areaStyle: {
            normal: {
              width: 1,
              opacity: 0.7,
            },
          },
          data: [{
            itemStyle: {
              normal: {
                color: '#67abff',
                lineStyle: {
                  color: '#67abff',
                },
              },
            },
            value: this.data
          }]
        }]
      };
    }
  }

}
