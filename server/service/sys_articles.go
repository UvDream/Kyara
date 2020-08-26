package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
)

//文章列表服务
func ArticleList(u request.ArticleListStruct) (err error, list []model.SysArticle, totalCount int64) {
	db := global.GVA_DB
	var articleList []model.SysArticle
	var topArticle []model.SysArticle
	var total int64
	//查询总数
	err = db.Find(&articleList).Count(&total).Error
	// 先查询置顶文章
	err = db.Where(&model.SysArticle{Top: "1"}).Find(&topArticle).Error

	err = db.Where(&model.SysArticle{Top: "1"}).Find(&topArticle).Error
	topLen := len(topArticle)
	//再查其他
	limit := u.PageSize - topLen
	offset := u.PageSize * (u.Page - 1)

	err = db.Where("top=? ", "0").Limit(limit).Offset(offset).Find(&articleList).Error

	articleList = append(topArticle, articleList...)
	//查询作者名
	for i, k := range articleList {
		a := model.SysUser{}
		//	根据userId查询作者姓名
		err = db.Where("UUID=?", k.UserID).Find(&a).Error
		articleList[i].UserName = a.NickName
	}

	return err, articleList, total
}
//文章详情服务
func GetArticleDetail(c *gin.Context) (err error,article model.SysArticle) {
	db := global.GVA_DB
	id:=c.Query("id")
	 article =model.SysArticle{}
	err = db.Where("id=? ", id).Find(&article).Error
	luteEngine := lute.New()
	article.ArticleHtml=luteEngine.MarkdownStr("uvdrem",article.ArticleContent)
	return err,article
}
