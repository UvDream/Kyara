package service

import (
	"server/global"
	"server/model"
	"server/model/request"
	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
)

//文章列表服务
func ArticleList(u request.ArticleListStruct) (err error, list []model.SysArticle, totalCount int64,msg string) {
	db := global.GVA_DB
	var articleList []model.SysArticle
	var topArticle []model.SysArticle
	var total int64
	//查询总数
	err = db.Find(&articleList).Count(&total).Error
	// 先查询置顶文章
	err = db.Where(&model.SysArticle{Top: "1"}).Find(&topArticle).Error
	topLen := len(topArticle)
	//再查其他
	u.PageSize=u.PageSize-topLen
	limit := u.PageSize
	offset := u.PageSize * (u.Page - 1)
	err = db.Where("top=? ", "0").Limit(limit).Offset(offset).Find(&articleList).Error

	articleList = append(topArticle, articleList...)
	//查询作者名
	for i, k := range articleList {
		a := model.SysUser{}
		//	根据userId查询作者姓名
		err = db.Where("UUID=?", k.UserID).Find(&a).Error
		if err!=nil{
			return err,articleList,0,"获取用户错误"
		}
		articleList[i].UserName = a.NickName
	}

	return err, articleList, total,msg
}
//文章详情服务
func GetArticleDetail(c *gin.Context) (err error,article model.SysArticle) {
	db := global.GVA_DB
	id:=c.Query("id")
	article =model.SysArticle{}
	err = db.Where("id=? ", id).Find(&article).Error
	luteEngine := lute.New()
	article.ArticleHtml=luteEngine.MarkdownStr("uvdrem",article.ArticleContent)
	a := model.SysUser{}
	err = db.Where("UUID=?", article.UserID).Find(&a).Error
	article.UserName=a.NickName
	return err,article
}
//验证密码
func CheckPassword(id string,password string) (error bool ,msg string) {
	article :=model.SysArticle{}
	db := global.GVA_DB
	err := db.Where("id=? ", id).Find(&article).Error
	if article.ViewPassword!=password && err==nil{
		return  false,"密码错误"
	}
	return true,"密码正确"
}

//查询分类
func ArticleClassify() (err error,a []model.SysClassify) {
	db := global.GVA_DB
	//查询一级
	err = db.Where("parent_id=? ", "").Find(&a).Error
	if err!=nil{
		return
	}
	for k,i:=range a {
		a[k].Children=findChildren(i)
	}
	return err,a
}
//递归寻找分类
func findChildren(parent model.SysClassify)(children []model.SysClassify)  {
	db := global.GVA_DB
	var c []model.SysClassify
	err := db.Where("parent_id=? ", parent.ID).Find(&c).Error
	if err!=nil{
		return
	}
	for k,i :=range c{
		c[k].Children=findChildren(i)
	}
	return c
}
//热门文章
func HotArticle()(err error,list []model.SysArticle){
	db := global.GVA_DB
	var arr []model.SysArticle
	err =db.Order("like_count desc").Find(&arr).Error
	for k,i:=range arr{
		if k < 5 {
			list=append(list,i)
		}
	}
	return err,list
}