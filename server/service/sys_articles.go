package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

func ArticleList(u request.ArticleListStruct)(err error,list []model.SysArticle,totalCount int64)  {
		//limit:=u.PageSize
		//offset:=u.PageSize*(u.Page-1)
		db := global.GVA_DB
		var articleList []model.SysArticle
		var total int64
		//err=db.Limit(limit).Offset(offset).Find(&articleList).Count(&total).Error
		// 先查询置顶文章
		var topArticle []model.SysArticle
		err=db.Where(&model.SysArticle{Top:"1"}).Find(&topArticle).Error
		fmt.Println(topArticle)
		return err,articleList,total
}