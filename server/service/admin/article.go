package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//新增文章
func AddArticle(r model.SysArticle) (err error, msg string) {
	db := global.GVA_DB
	if r.ArticleID == "" {
		//增加文章
		err := db.Create(&r).Error
		if err != nil {
			return err, "文章创建失败"
		}
	} else {
		//查询文章是否存在
		err := db.Where("ID=?", r.ArticleID).Find(&model.SysArticle{}).Error
		if err != nil {
			return err, "文章不存在"
		}
		err = db.Model(&model.SysArticle{}).Where("ID=?", r.ArticleID).Update(&r).Error
		if err != nil {
			return err, "更新失败"
		} else {
			return err, "更新文章成功"
		}
	}
	return nil, ""
}
