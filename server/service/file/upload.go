package file

import (
	"mime/multipart"
	"server/global"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, int, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	switch global.Config.System.OssType {
	case "local":
		return &LocalService{}
	case "qiniu":
		return &QiniuService{}
	default:
		return &LocalService{}
	}
}

func DeleteOss(position string) OSS {
	switch position {
	case "local":
		return &LocalService{}
	case "qiniu":
		return &QiniuService{}
	default:
		return &LocalService{}
	}
}
