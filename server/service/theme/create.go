package theme

import (
	code2 "server/code"
	"server/global"
	"server/model/theme"
)

func (*ThemesServiceGroup) CreateThemeService(theme theme.Theme) (t theme.Theme, code int, err error) {
	db := global.DB
	if err = db.Create(&theme).Error; err != nil {
		return t, code2.ErrorCreateTheme, err
	}
	return t, code2.SUCCESS, err
}
