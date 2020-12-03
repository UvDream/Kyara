package model

import "github.com/jinzhu/gorm"

type ImageList struct {
	gorm.Model
	SSID string `json:"ssid" gorm:"comment:'唯一辨识符'"`
	Type string `json:"type" gorm:"comment:'上传使用场景,例如:文章编辑上传article'"`
}
