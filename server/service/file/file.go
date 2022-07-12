package file

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"server/code"
	"server/global"
	file2 "server/model/file"
	"strings"
)

type FilesService struct{}

// UploadFileService 文件上传
func (*FilesService) UploadFileService(c *gin.Context) (data file2.File, ce int, err error) {
	_, fileHeader, _ := c.Request.FormFile("file")
	url, key, codes, err := NewOss().UploadFile(fileHeader)
	if err != nil {
		return data, codes, err
	}
	data, ce, err = SaveFileData(fileHeader, url, key)
	return data, ce, err
}

// SaveFileData 保存数据到数据库
func SaveFileData(fileHeader *multipart.FileHeader, url string, key string) (data file2.File, ce int, err error) {
	db := global.DB
	var file file2.File
	ext := path.Ext(fileHeader.Filename)
	file.Name = strings.TrimSuffix(fileHeader.Filename, ext)
	file.URL = url
	file.Size = fileHeader.Size
	file.Type = fileHeader.Header.Get("Content-Type")
	file.Position = global.Config.System.OssType
	file.Key = key
	if err := db.Create(&file).Error; err != nil {
		return file, code.ErrorSaveFileData, err
	}
	return file, code.SUCCESS, nil
}
