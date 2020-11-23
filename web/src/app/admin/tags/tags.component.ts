import { Component, OnInit } from '@angular/core';
import { ArticleService } from '@service/article.service';
interface TagInterface {
  ID: number;
  tag_name: string;
  tag_count: number;
  color: string;
}
@Component({
  selector: 'app-tags',
  templateUrl: './tags.component.html',
  styleUrls: ['./tags.component.less']
})
export class TagsComponent implements OnInit {
  public tags: Array<TagInterface>;
  constructor(private service: ArticleService) { }

  ngOnInit(): void {
    this.getAllTag();
  }
  async getAllTag(): Promise<void> {
    const res = await this.service.getTag();
    if (res.code === 200) {
      this.tags = res.data;
    }
  }
}
