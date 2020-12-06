import { Component, OnInit, Input, Output, EventEmitter, forwardRef } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzUploadChangeParam } from 'ng-zorro-antd/upload';
import { NzUploadFile } from 'ng-zorro-antd/upload';
import { AdminService } from '@service/admin.service';
@Component({
  selector: 'app-upload-image',
  templateUrl: './upload-image.component.html',
  styleUrls: ['./upload-image.component.less'],
})
export class UploadImageComponent implements OnInit {

  constructor(private msg: NzMessageService, private adminHttp: AdminService) { }
  loading = false;
  avatarUrl?: string;
  @Input() type = 'article';
  @Input() url: string;
  @Output() urlChange = new EventEmitter();
  ngOnInit(): void {
    this.avatarUrl = this.url;
  }

  beforeUpload = async (file: any, _fileList: NzUploadFile[]) => {
    console.log(file);
    const formData = new FormData();
    formData.append('type', this.type);
    formData.append('file', file);
    this.loading = true;
    const res = await this.adminHttp.uploadImage(formData);
    if (res.code === 200) {
      this.loading = false;
      this.avatarUrl = res.data.url;
      this.urlChange.emit(this.avatarUrl);
    }
    // return new Observable((observer: Observer<boolean>) => {
    //   const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    //   if (!isJpgOrPng) {
    //     this.msg.error('You can only upload JPG file!');
    //     observer.complete();
    //     return;
    //   }
    //   const isLt2M = file.size! / 1024 / 1024 < 2;
    //   if (!isLt2M) {
    //     this.msg.error('Image must smaller than 2MB!');
    //     observer.complete();
    //     return;
    //   }
    //   observer.next(isJpgOrPng && isLt2M);
    //   observer.complete();
    // });
  }

  handleChange(info: NzUploadChangeParam): void {

  }


}
