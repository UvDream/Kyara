package theme

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"server/model/system"
)

type Theme struct {
	ResponseTheme
	Theme string `json:"theme" gorm:"type:longblob" binding:"required"` // 主题
}
type ResponseTheme struct {
	gorm.Model
	//名字
	Name string `json:"name" gorm:"type:varchar(100);"`
	//描述
	Description string `json:"description" gorm:"type:varchar(100);"`
	//简略图
	Thumbnail string `json:"thumbnail" gorm:"type:varchar(100);"`
	//作者ID
	AuthID uuid.UUID `json:"auth_id" gorm:"comment:作者的UUID"`
	//是否公开
	IsPublic bool `json:"is_public" gorm:"type:tinyint(1);"`
	//作者
	Auth system.SysUser `gorm:"references:UUID" json:"author"`
}
