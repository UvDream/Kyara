package file

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"server/code"
	"server/global"
	file2 "server/model/file"
	"server/utils"
)

type ImageService struct{}

// UploadImageService 图片上传
func (i *ImageService) UploadImageService(c *gin.Context) (data file2.File, ce int, err error) {
	file, fileHeader, _ := c.Request.FormFile("file")
	print(file, fileHeader)
	//	获取存储方式
	ossType := global.Config.System.OssType
	if ossType == "local" {
		filePosition := global.Config.System.FilePosition
		data, ce, err := SaveFileLocal(file, fileHeader, filePosition)
		return data, ce, err
	}
	return data, code.UploadImageSuccess, nil
}

// SaveFileLocal 保存文件到本地
func SaveFileLocal(file multipart.File, fileHeader *multipart.FileHeader, filePosition string) (data file2.File, c int, err error) {
	//	获取文件名
	fileName := fileHeader.Filename
	//	获取文件大小
	fileSize := fileHeader.Size
	//	获取文件类型
	fileType := fileHeader.Header.Get("Content-Type")
	//	获取文件内容
	fileContent, _ := ioutil.ReadAll(file)
	//	文件路径
	filePath := "." + filePosition + "/" + fileName
	url := filePosition + "/" + fileName
	//路径是否存在,不存在创建
	err = utils.CreateDir("." + filePosition)
	if err != nil {
		return data, code.ErrorCreateDir, err
	}
	//	保存文件
	err = ioutil.WriteFile(filePath, fileContent, 0666)
	if err != nil {
		return data, code.ErrorSaveFile, err
	}
	data, c, err = SaveFileData(fileName, filePath, url, fileSize, fileType)
	if err != nil {
		return data, c, err
	}
	return data, code.SaveFileSuccess, nil
}

// SaveFileData 保存数据到数据库
func SaveFileData(name string, path string, url string, size int64, fileType string) (data file2.File, ce int, err error) {
	db := global.DB
	var file file2.File
	file.Name = name
	file.Path = path
	file.URL = url
	file.Size = size
	file.Type = fileType
	if err := db.Create(&file).Error; err != nil {
		return file, code.ErrorSaveFileData, err
	}
	return file, code.SUCCESS, nil
}
