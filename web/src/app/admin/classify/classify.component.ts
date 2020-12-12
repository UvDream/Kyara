import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
import { NzTreeNode } from 'ng-zorro-antd/tree';
import { AdminService } from '@service/admin.service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';

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
    private modal: NzModalService,
    private message: NzMessageService
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
    };
  }
  treeClick(data: NzTreeNode): void {
    console.log(data);
    const obj = data.origin;
    this.form.type_name = obj.title;
    this.form.pid = Number(obj.parentId);
    this.form.icon = obj.icon;
    this.form.id = obj.key;
  }
  // 删除分类
  deleteFunc(node: TreeItem): void {
    this.modal.confirm({
      nzTitle: '提示',
      nzContent: '删除此目录?',
      nzOnOk: async () => {
        const res = await this.httpAdmin.deleteClassify({ id: node.key });
        if (res.code === 200) {
          this.message.success('删除成功');
          this.getTree();
          this.cancelFunc();
        }
      },
    });
  }
  // 添加子类
  addFunc(data: NzTreeNode): void {
    console.log(data);
    const obj = data.origin;
    this.form.type_name = '';
    this.form.pid = Number(obj.parentId);
    this.form.icon = '';
    this.form.id = '';
  }

}
