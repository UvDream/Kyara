package theme

import (
	"github.com/google/uuid"
	code2 "server/code"
	"server/global"
	"server/model/theme"
)

type ThemesServiceGroup struct{}

// GetThemeListService 获取主题列表
func (*ThemesServiceGroup) GetThemeListService(keyword string, uuid uuid.UUID) (t []theme.ResponseTheme, code int, err error) {
	db := global.DB
	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%").Or("description like ?", "%"+keyword+"%").Or("theme like ?", "%"+keyword+"%")
	}
	if err = db.Model(&theme.Theme{}).Where("auth_id = ?", uuid).Find(&t).Error; err != nil {
		return t, code2.ErrorGetTheme, err
	}
	return t, code2.SUCCESS, err
}
