package file

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
	"mime/multipart"
	"server/code"
	"server/global"
	"time"
)

type QiniuService struct{}

// UploadFile 上传到七牛
func (*QiniuService) UploadFile(file *multipart.FileHeader) (path string, key string, c int, err error) {
	putPolicy := storage.PutPolicy{
		Scope: global.Config.Qiniu.Bucket,
	}
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := GetQiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return "", "", code.ErrorUploadQiNiu, e
	}
	dataLen := file.Size
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	err = formUploader.Put(context.Background(), &ret, upToken, fileKey, f, dataLen, &putExtra)
	if err != nil {
		return "", "", code.ErrorUploadQiNiu, err
	}
	return global.Config.Qiniu.DomainName + "/" + ret.Key, ret.Key, code.ErrorUploadQiNiuSuccess, err
}

// GetQiniuConfig 获取七牛配置
func GetQiniuConfig() *storage.Config {
	cfg := storage.Config{}
	// 空间对应的机房
	domain := global.Config.Qiniu.Domain
	//区域设置
	if domain == "HuaDong" {
		cfg.Zone = &storage.ZoneHuadong
	} else if domain == "HuaBei" {
		cfg.Zone = &storage.ZoneHuabei
	} else if domain == "HuaNan" {
		cfg.Zone = &storage.ZoneHuanan
	} else if domain == "BeiMei" {
		cfg.Zone = &storage.ZoneBeimei
	} else if domain == "XinJiaPo" {
		cfg.Zone = &storage.ZoneXinjiapo
	}
	// 是否使用https域名
	if global.Config.Qiniu.DomainProtocol == "https" {
		cfg.UseHTTPS = true
	} else {
		cfg.UseHTTPS = false
	}
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return &cfg
}

// DeleteFile 七牛云删除文件
func (*QiniuService) DeleteFile(key string) error {
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	cfg := GetQiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.Config.Qiniu.Bucket, key); err != nil {
		global.Log.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("七牛云删除错误, err:" + err.Error())
	}
	return nil
}
