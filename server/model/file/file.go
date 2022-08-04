package file

import (
	"server/model/common"
)

type File struct {
	common.Model
	Name     string `json:"name" gorm:"type:varchar(255);"`
	URL      string `json:"url" gorm:"type:varchar(255);"`
	Size     int64  `json:"size" gorm:"type:int(10);"`
	Type     string `json:"type" gorm:"type:varchar(255);"`
	Position string `json:"position" gorm:"type:varchar(255);"`
	Key      string `json:"key" gorm:"type:varchar(255);"`
	AuthID   string `json:"auth_id" gorm:"type:varchar(255);"`
}
