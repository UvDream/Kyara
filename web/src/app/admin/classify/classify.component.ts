import { Component, OnInit } from '@angular/core';
import { NzFormatEmitEvent, NzTreeComponent, NzTreeNodeOptions } from 'ng-zorro-antd/tree';
import { ArticleService } from '@service/article.service';

interface TreeItem {
  key: string;
  title: string;
  isLeaf?: boolean;
  icon?: string;
  expanded?: boolean;
  children: Array<TreeItem>;
}
@Component({
  selector: 'app-classify',
  templateUrl: './classify.component.html',
  styleUrls: ['./classify.component.less']
})
export class ClassifyComponent implements OnInit {

  constructor(
    private httpService: ArticleService,
  ) { }
  nodes = [];
  ngOnInit(): void {
    this.getTree();
  }
  async getTree(): Promise<void> {
    const res = await this.httpService.getArticleClassification();
    if (res.code === 200) {
      this.nodes = this.changeData(res.data);
    }
  }
  changeData = (arr: any): Array<TreeItem> => {
    const data: Array<TreeItem> = [];
    arr.forEach(element => {
      const obj: TreeItem = {
        title: '',
        key: '',
        icon: '',
        expanded: true,
        children: []
      };
      obj.icon = element.icon;
      obj.title = element.type_name;
      obj.key = element.ID;
      element.children.length > 0 ? obj.children = this.changeData(element.children) : obj.isLeaf = true;
      data.push(obj);
    });
    return data;
  }

}
