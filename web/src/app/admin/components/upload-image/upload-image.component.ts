import { Component, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzUploadChangeParam } from 'ng-zorro-antd/upload';
import { NzUploadFile } from 'ng-zorro-antd/upload';
import { Observable, Observer } from 'rxjs';
import { AdminService } from '@service/admin.service'
@Component({
  selector: 'app-upload-image',
  templateUrl: './upload-image.component.html',
  styleUrls: ['./upload-image.component.less']
})
export class UploadImageComponent implements OnInit {

  constructor(private msg: NzMessageService, private adminHttp: AdminService) { }
  loading = false;
  avatarUrl?: string;
  ngOnInit(): void {


  }
  beforeUpload = async (file: any, _fileList: NzUploadFile[]) => {
    console.log(file);
    const formData = new FormData();
    formData.append('type', 'article');
    formData.append('file', file);
    this.loading = true;
    const res = await this.adminHttp.uploadImage(formData);
    if (res.code === 200) {
      this.loading = false;
      this.avatarUrl = res.data.url;
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
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList);
    }
    if (info.file.status === 'done') {
      this.msg.success(`${info.file.name} file uploaded successfully`);
    } else if (info.file.status === 'error') {
      this.msg.error(`${info.file.name} file upload failed.`);
    }
  }


}
