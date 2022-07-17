package theme

import (
	"github.com/google/uuid"
	code2 "server/code"
	"server/global"
	"server/model/theme"
)

type ThemesServiceGroup struct{}

// GetThemeListService 获取主题列表
func (*ThemesServiceGroup) GetThemeListService(keyword string, uuid uuid.UUID) (theme theme.Theme, code int, err error) {
	db := global.DB
	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%").Or("description like ?", "%"+keyword+"%").Or("theme like ?", "%"+keyword+"%")
	}
	if err = db.Where("auth_id = ?", uuid).Preload("Auth").Find(&theme).Error; err != nil {
		return theme, code2.ErrorGetTheme, err
	}
	return theme, code2.SUCCESS, err
}
