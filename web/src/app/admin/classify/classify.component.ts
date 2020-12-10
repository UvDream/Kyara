import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { NzFormatEmitEvent } from 'ng-zorro-antd/tree';
import { AdminService } from '@service/admin.service';

interface TreeItem {
  key: string;
  title: string;
  isLeaf?: boolean;
  icon?: string;
  expanded?: boolean;
  children: Array<TreeItem>;
  parentId?: number;
}
@Component({
  selector: 'app-classify',
  templateUrl: './classify.component.html',
  styleUrls: ['./classify.component.less']
})
export class ClassifyComponent implements OnInit {

  constructor(
    private httpService: ArticleService,
    private httpAdmin: AdminService,
  ) { }
  nodes = [];
  form = {
    type_name: '',
    id: '',
    parent_id: null,
    icon: '',
    pid: null,
  };
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
        children: [],
        parentId: null
      };
      obj.icon = element.icon;
      obj.title = element.type_name;
      obj.parentId = element.parent_id;
      obj.key = element.ID;
      element.children.length > 0 ? obj.children = this.changeData(element.children) : obj.isLeaf = true;
      data.push(obj);
    });
    return data;
  }
  async sureFunc(): Promise<void> {
    console.log(this.form);
    if (this.form.pid) {
      this.form.parent_id = this.form.pid.toString();
    }
    const res = await this.httpAdmin.editClassify(this.form);
    if (res.code === 200) {
      console.log(res);
      this.getTree();
      this.form.id = res.data.ID;
    }
  }
  cancelFunc(): void {
    this.form = {
      type_name: '',
      id: '',
      parent_id: null,
      icon: '',
      pid: null,
    }
  }
  treeClick(data: NzFormatEmitEvent): void {
    console.log(data.node.origin);
    const obj = data.node.origin;
    this.form.type_name = obj.title;
    this.form.pid = Number(obj.parentId);
    this.form.icon = obj.icon;
    this.form.id = obj.key;
  }

}
