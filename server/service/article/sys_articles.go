package article

import (
	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model"
	"server/model/request"
)

//文章列表服务
func ArticleList(u request.ArticleListStruct) (err error, list []model.SysArticle, totalCount int64, msg string) {
	db := global.GVA_DB
	var articleList []model.SysArticle
	var topArticle []model.SysArticle
	var total int64
	//查询总数
	if u.Search=="" {
		err = db.Table("sys_articles").Count(&total).Error
	}else{
		err = db.Where("title LIKE ? ","%"+u.Search+"%").Or("article_content LIKE ?","%"+u.Search+"%").Table("sys_articles").Count(&total).Error
	}

	// 先查询置顶文章
	if u.ClassificationID == "" && u.Tag == "" {
		if(u.Search==""){
			err = db.Where(&model.SysArticle{Top: "1"}).Find(&topArticle).Error
		}
	}
	if u.ClassificationID != "" {
		err = db.Where(&model.SysArticle{Top: "1", ClassificationID: u.ClassificationID}).Where("title LIKE ?","%"+u.Search+"%").Or("article_content LIKE ?","%"+u.Search+"%").Order("update_time desc").Find(&topArticle).Error
	}
	if u.Tag != "" {
		topArticle, err = tagToArticle(u.Tag, "1")
	}
	topLen := len(topArticle)
	//再查其他
	u.PageSize = u.PageSize - topLen
	limit := u.PageSize
	offset := u.PageSize * (u.Page - 1)
	if u.ClassificationID == "" && u.Tag == "" {
		if u.Search==""{
			err = db.Where("top = ?","0").Limit(limit).Offset(offset).Order("update_time desc").Find(&articleList).Error
		}else{
			err = db.Where("title LIKE ? ","%"+u.Search+"%").Or("article_content LIKE ?","%"+u.Search+"%").Limit(limit).Offset(offset).Order("update_time desc").Find(&articleList).Error
		}



	}
	if u.ClassificationID != "" {
		err = db.Where("top=? AND classification_id=?", "0", u.ClassificationID).Where("title LIKE ?","%"+u.Search+"%").Limit(limit).Offset(offset).Order("update_time desc").Find(&articleList).Error
	}
	if u.Tag != "" {
		// 根据tag查询文章
		articleList, err = tagToArticle(u.Tag, "0")
	}
	articleList = append(topArticle, articleList...)
	//查询作者名
	for i, k := range articleList {
		a := model.SysUser{}
		//	根据userId查询作者姓名
		err = db.Where("UUID=?", k.UserID).Find(&a).Error
		if err != nil {
			return err, articleList, 0, "获取用户错误"
		}
		articleList[i].UserName = a.NickName
	}

	return err, articleList, total, msg
}

//根据tag_id 查询文章
func tagToArticle(tag string, top string) (articleList []model.SysArticle, err error) {
	db := global.GVA_DB
	//	先去查询中间表
	var tagLink []model.SysArticleTag
	err = db.Where("tag_id=?", tag).Find(&tagLink).Error
	// 再根据结果查询文章
	for _, i := range tagLink {
		article := model.SysArticle{}
		err = db.Where("id=? AND top=?", i.ArticleID, top).Find(&article).Error
		if err == nil {
			articleList = append(articleList, article)
		}
	}
	return articleList, err
}

//文章详情服务
func GetArticleDetail(c *gin.Context) (err error, article model.SysArticle, msg string) {
	db := global.GVA_DB
	id := c.Query("id")
	article = model.SysArticle{}
	err = db.Where("id=? ", id).Find(&article).Error
	luteEngine := lute.New()
	article.ArticleHtml = luteEngine.MarkdownStr("UvDream", article.ArticleContent)
	//查询文章作者
	a := model.SysUser{}
	err = db.Where("UUID=?", article.UserID).Find(&a).Error
	if err != nil {
		return err, article, "获取文章作者失败"
	}
	article.UserName = a.NickName
	//查询文章赞赏码
	var collectList []model.CollectionCode
	err = db.Where("article_id=?", id).Find(&collectList).Error
	if err != nil {
		return err, article, "获取文章赞赏码失败"
	}
	article.CollectList = collectList
	return err, article, "获取文章详情成功"
}

//验证密码
func CheckPassword(id string, password string) (error bool, msg string) {
	article := model.SysArticle{}
	db := global.GVA_DB
	err := db.Where("id=? ", id).Find(&article).Error
	if article.ViewPassword != password && err == nil {
		return false, "密码错误"
	}
	return true, "密码正确"
}

//查询分类
func ArticleClassify() (err error, a []model.SysClassify) {
	db := global.GVA_DB
	//查询一级
	err = db.Where("parent_id=? ", "").Find(&a).Error
	if err != nil {
		return
	}
	for k, i := range a {
		a[k].Children = findChildren(i)
	}
	return err, a
}

//递归寻找分类
func findChildren(parent model.SysClassify) (children []model.SysClassify) {
	db := global.GVA_DB
	var c []model.SysClassify
	err := db.Where("parent_id=? ", parent.ID).Find(&c).Error
	if err != nil {
		return
	}
	for k, i := range c {
		c[k].Children = findChildren(i)
	}
	return c
}

//热门文章
func HotArticle() (err error, list []model.SysArticle) {
	db := global.GVA_DB
	var arr []model.SysArticle
	err = db.Order("view_count desc").Find(&arr).Error
	for k, i := range arr {
		if k < 5 {
			i.ArticleContent = ""
			list = append(list, i)
		}
	}
	return err, list
}

//所有tag
func AllTag() (err error, tagArr []model.SysTag, msg string) {
	db := global.GVA_DB
	err = db.Find(&tagArr).Error
	if err != nil {
		return err, tagArr, "获取tag失败"
	}
	for i, k := range tagArr {
		//	用tagID查询文章tag中间表
		var total int64
		err = db.Where("tag_id=?", k.ID).Table("sys_article_tags").Count(&total).Error
		if err != nil {
			return err, tagArr, "查询tag数量错误"
		}
		tagArr[i].TagCount = total
	}
	return err, tagArr, "获取tag成功"
}
