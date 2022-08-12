package theme

import (
	"gorm.io/gorm/clause"
	code2 "server/code"
	"server/global"
	"server/models/theme"
)

type ThemesServiceGroup struct{}

// GetThemeListService 获取主题列表
func (*ThemesServiceGroup) GetThemeListService(keyword string, uuid string) (t []theme.Theme, code int, err error) {
	db := global.DB
	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%").Or("description like ?", "%"+keyword+"%").Or("theme like ?", "%"+keyword+"%")
	}
	if err = db.Model(&theme.Theme{}).Where("user_id = ?", uuid).Preload(clause.Associations).Find(&t).Error; err != nil {
		return t, code2.ErrorGetTheme, err
	}
	return t, code2.SUCCESS, err
}

//GetPublicThemeListService 获取公共主题列表
func (*ThemesServiceGroup) GetPublicThemeListService(keyword string, uuid string) (t []theme.Theme, code int, err error) {
	db := global.DB
	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%").Or("description like ?", "%"+keyword+"%").Or("theme like ?", "%"+keyword+"%")
	}
	if err = db.Model(&theme.Theme{}).Where("is_public=?", 1).Where("user_id<>?", uuid).Preload(clause.Associations).Find(&t).Error; err != nil {
		return t, code2.ErrorGetTheme, err
	}
	return t, code2.SUCCESS, err
}
