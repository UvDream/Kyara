package file

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"server/code"
	"server/global"
	file2 "server/models/file"
	"server/utils"
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
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	data, ce, err = SaveFileData(fileHeader, url, key, claims.ID)
	return data, ce, err
}
func (*FilesService) DeleteFileService(id string) (file file2.File, codeNumber int, err error) {
	db := global.DB
	if err := db.Where("id = ?", id).First(&file).Error; err != nil {
		return file, code.ErrorFileNotFound, err
	}
	// 删除os 文件
	err = DeleteOss(file.Position).DeleteFile(file.Key)
	if err != nil {
		return file, code.ErrorDeleteFile, err
	}
	//删除数据库文件
	if err := db.Where("id = ?", id).Delete(&file).Error; err != nil {
		return file, code.ErrorDeleteFileData, err
	}
	return file, code.SUCCESS, err
}

// SaveFileData 保存数据到数据库
func SaveFileData(fileHeader *multipart.FileHeader, url string, key string, authID string) (data file2.File, ce int, err error) {
	db := global.DB
	var file file2.File
	ext := path.Ext(fileHeader.Filename)
	file.Name = strings.TrimSuffix(fileHeader.Filename, ext)
	file.URL = url
	file.Size = fileHeader.Size
	file.Type = fileHeader.Header.Get("Content-Type")
	file.Position = global.Config.System.OssType
	file.Key = key
	file.AuthID = authID
	if err := db.Create(&file).Error; err != nil {
		return file, code.ErrorSaveFileData, err
	}
	return file, code.SUCCESS, nil
}
