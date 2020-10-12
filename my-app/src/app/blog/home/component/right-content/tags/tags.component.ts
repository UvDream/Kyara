import { Component, OnInit } from '@angular/core';
import { ArticleService } from '../../../../../service/article.service';
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
  constructor(private service: ArticleService) { }
  public tags: Array<TagInterface>;
  ngOnInit(): void {
    this.getAllTag();
  }
  getAllTag = async () => {
    const res = await this.service.getTag();
    if (res.code === 200) {
      this.tags = res.data;
    }
  }
}
