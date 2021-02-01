package service

import (
	"server/global"
	"server/model"
)

func SaveConfig(r model.SysConfig) (err error,msg string) {
	db := global.GVA_DB
		err=db.Model(model.SysConfig{}).Updates(model.SysConfig{
			//更新博客相关配置
			BlogName: r.BlogName,
			BlogStartTime: r.BlogStartTime,
			BlogLogo: r.BlogLogo,
			DomainName: r.DomainName,
			//	更新用户相关配置
			AuthorName: r.AuthorName,
			AuthorMotto: r.AuthorMotto,
			AuthorLink: r.AuthorLink,
			AuthorAvatar:r.AuthorAvatar,
		}).Error
		if err!=nil {
			return err,"更新失败"
		}
	return err,"更新成功"
}
