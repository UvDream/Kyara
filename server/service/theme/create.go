package theme

import (
	code2 "server/code"
	"server/global"
	"server/models/theme"
)

//CreateThemeService 创建主题
func (*ThemesServiceGroup) CreateThemeService(theme theme.Theme) (theme.Theme, int, error) {
	db := global.DB
	if err := db.Create(&theme).Error; err != nil {
		return theme, code2.ErrorCreateTheme, err
	}
	return theme, code2.SUCCESS, nil
}

// DeleteThemeService 删除主题
func (*ThemesServiceGroup) DeleteThemeService(id string) (t theme.Theme, code int, err error) {
	db := global.DB
	if err := db.Where("id=?", id).First(&t).Error; err != nil {
		return t, code2.ErrorThemeNotFound, err
	}
	if err := db.Where("id=?", id).Preload("Auth").Delete(&t).Error; err != nil {
		return t, code2.ErrorDeleteTheme, err
	}
	return t, code2.SUCCESS, err
}

// UpdateThemeService 更新主题
func (*ThemesServiceGroup) UpdateThemeService(data theme.Theme) (t theme.Theme, code int, err error) {
	db := global.DB
	if err := db.Where("id=?", data.ID).First(&t).Error; err != nil {
		return t, code2.ErrorThemeNotFound, err
	}
	if err := db.Model(&t).Where("id=?", data.ID).Updates(&data).Error; err != nil {
		return t, code2.ErrorUpdateTheme, err
	}
	return t, code2.SUCCESS, err
}
