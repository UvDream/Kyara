import { Component, Inject, OnInit, PLATFORM_ID } from '@angular/core';
import { ArticleService } from '@service/article.service';
// import '../../../../../../assets/echarts-wordcloud.js';
import { isPlatformBrowser } from '@angular/common';

interface TagInterface {
  ID: number;
  tag_name: string;
  tag_count: number;
}
@Component({
  selector: 'app-tags',
  templateUrl: './tags.component.html',
  styleUrls: ['./tags.component.less'],
})
export class TagsComponent implements OnInit {
  constructor(
    private service: ArticleService,
    @Inject(PLATFORM_ID) private platformId: object
  ) { }
  public tags: Array<TagInterface>;
  options: any;
  loading = true;
  tagArr = [];
  ngOnInit(): void {
    this.getAllTag();
  }
  getAllTag = async () => {
    this.loading = true;
    const res = await this.service.getTag();
    if (res.code === 200) {
      this.tags = res.data;
      this.tags.forEach(element => {
        const obj = { name: '', value: 0 };
        obj.name = element.tag_name;
        obj.value = element.tag_count;
        this.tagArr.push(obj);
      });
      this.loading = false;
      if (isPlatformBrowser(this.platformId)) {
        import('../../../../../../assets/echarts-wordcloud.js').then(() => {
          this.options = {
            tooltip: {
              show: true,
              position: 'top',
              textStyle: {
                fontSize: 12
              }
            },
            series: [{
              type: 'wordCloud',
              // 网格大小，各项之间间距
              gridSize: 30,
              // 形状 circle 圆，cardioid  心， diamond 菱形，
              // triangle-forward 、triangle 三角，star五角星
              shape: 'circle',
              // 字体大小范围
              sizeRange: [12, 20],
              // 文字旋转角度范围
              rotationRange: [0, 90],
              // 旋转步值
              rotationStep: 90,
              // 自定义图形
              // maskImage: maskImage,
              left: 'center',
              top: 'center',
              right: null,
              bottom: null,
              // 画布宽
              width: '100%',
              // 画布高
              height: '100%',
              // 是否渲染超出画布的文字
              drawOutOfBound: false,
              textStyle: {
                normal: {
                  color: () => {
                    return 'rgb(' + [
                      Math.round(Math.random() * 200 + 55),
                      Math.round(Math.random() * 200 + 55),
                      Math.round(Math.random() * 200 + 55)
                    ].join(',') + ')';
                  }
                },
                emphasis: {
                  shadowBlur: 10,
                  shadowColor: '#2ac'
                }
              },
              data: this.tagArr
            }]
          };
        });


      }
    }
  }
}
