package model

import (
	"github.com/jinzhu/gorm"
)

type ExaFileUploadAndDownload struct {
	gorm.Model
	Name     string `json:"name" gorm:"comment:'文件名'"`
	Url      string `json:"url" gorm:"comment:'文件地址'"`
	Tag      string `json:"tag" gorm:"comment:'文件标签'"`
	Key      string `json:"key" gorm:"comment:'编号'"`
	Type     string `json:"type" gorm:"comment:'上传位置'"`
	Position string `json:"position" gorm:"comment:'图片所在位置0:如优,1:白熊,2:七牛'"`
}
