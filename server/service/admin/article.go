package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model"
	"server/model/request"
	"server/service/blog"
	"time"
)

//新增文章
func AddArticle(r model.SysArticle, c *gin.Context) (err error, msg string, data model.SysArticle) {
	db := global.GVA_DB
	blog.AddDynamic()
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	r.UserID=waitUse.UUID

	if r.ImgURL == "" {
		r.ImgURL = "https://gitee.com/UvDream/images/raw/master/images/20200826195538.png"
	}
	r.UpdateTime = time.Now()
	//luteEngine := lute.New()
	//r.ArticleHtml = luteEngine.MarkdownStr("UvDream", r.ArticleContent)
	//counter := &utils.Counter{}
	//counter.Stat(string(r.ArticleHtml))
	//r.WordCount=counter.Total
	if r.ArticleID == 0 {
		//增加文章
		err := db.Create(&r).Error
		if err != nil {
			return err, "文章创建失败", r
		} else {
			err, msg := updateTag(r.ID, r.TagArray)
			if err != nil {
				return err, msg, r
			}
			//增加赞赏码
			err, msg = addCollect(r, c)
			if err != nil {
				return err, msg, r
			}
			return err, "文章创建成功", r
		}
	} else {
		//查询文章是否存在
		err := db.Where("ID=?", r.ArticleID).Find(&model.SysArticle{}).Error
		if err != nil {
			return err, "文章不存在", r
		}
		err = db.Model(&model.SysArticle{}).Where("ID=?", r.ArticleID).Update(&r).Error
		if err != nil {
			return err, "更新失败", r
		}
		err, msg := updateTag(r.ArticleID, r.TagArray)
		if err != nil {
			return err, msg, r
		}
		//查询是否存在赞赏码
		var collect model.CollectionCode
		err=db.Where("article_id=?",r.ArticleID).Find(&collect).Error
		if err!=nil {
			err, msg = addCollect(r, c)
			if err != nil {
				return err, msg, r
			}
		}
		return err, "更新文章成功", r

	}

}
func addCollect(r model.SysArticle, c *gin.Context) (err error, msg string) {
	db := global.GVA_DB
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	if r.CollectList == nil {
		var collect []model.CollectionCode
		err = db.Where("type=?", "0").Find(&collect).Error
		if err != nil {
			return err, "查询默认赞赏码失败"
		}
		for _, k := range collect {
			var obj model.CollectionCode
			obj.ArticleID = r.ArticleID
			obj.Name = k.Name
			obj.ImgURL = k.ImgURL
			obj.UserID = waitUse.UUID
			err := db.Create(&obj).Error
			if err != nil {
				return err, "增加赞赏码失败"
			}
		}
	} else {
		for _, k := range r.CollectList {
			err = db.Create(&k).Error
			if err != nil {
				return err, "赞赏码添加失败"
			}
		}
	}
	return err, "增加成功"
}
func updateTag(id uint, tag []uint) (err error, msg string) {
	//先删除再添加
	db := global.GVA_DB
	//强删除
	err = db.Where("article_id = ?", id).Unscoped().Delete(&model.SysArticleTag{}).Error
	if err != nil {
		return err, "删除失败"
	}
	//添加
	if len(tag) > 0 {
		for _, k := range tag {
			var tagArticle model.SysArticleTag
			tagArticle.ArticleID = id
			tagArticle.TagID = k
			err = db.Create(&tagArticle).Error
			if err != nil {
				return err, "创建文章和tag关系失败"
			}
		}
	}
	return err, "关联文章和tag关系成功"
}

//查询文章详情
func GetArticleDetail(id string) (err error, article model.SysArticle, msg string) {
	article = model.SysArticle{}
	db := global.GVA_DB
	err = db.Where("id=? ", id).Find(&article).Error
	if err != nil {
		return err, article, "查询详情失败"
	}
	//查询tag
	var tagArticle []model.SysArticleTag
	err = db.Where("article_id = ?", article.ID).Find(&tagArticle).Error
	if err != nil {
		return err, article, "查询tag失败"
	}
	fmt.Println(tagArticle)
	for i, k := range tagArticle {
		fmt.Println(i, k)
		article.TagArray = append(article.TagArray, k.TagID)
	}
	return err, article, "查询详情成功"
}

//新增公告
func AddNotice(r model.BlogNotice) (err error, msg string) {
	db := global.GVA_DB
	if r.Show == "1" {
		updateNotice()
	}
	if r.ID != 0 {
		err = db.Model(&model.BlogNotice{}).Update(&r).Error
	} else {
		err = db.Create(&r).Error
		if err != nil {
			return err, "新增失败"
		}
	}
	if r.Show == "1" {
		err = db.Model(&model.SysConfig{}).Update("blog_notice_id", r.ID).Error
		if err != nil {
			return err, "更新配置失败"
		}
	}
	blog.AddDynamic()
	return err, "成功"
}
func updateNotice() (err error, msg string) {
	var notice []model.BlogNotice
	db := global.GVA_DB
	err = db.Find(&notice).Error
	if err == nil {
		for i, k := range notice {
			if k.Show == "1" {
				notice[i].Show = "0"
			}
		}
		db.Model(&notice).UpdateColumn("show", "0")
	} else {
		return err, "更新数据失败"
	}
	blog.AddDynamic()
	return err, "更新成功"
}

//新增tag
func AddTag(c *gin.Context) (err error, tag model.SysTag, msg string) {
	t := c.Query("tag")
	db := global.GVA_DB
	tag.TagName = t
	var tagList []model.SysTag
	//需要查重
	db.Where("tag_name=?", t).Find(&tagList)
	if len(tagList) != 0 {
		return err, tag, "tag重复"
	}
	err = db.Create(&tag).Error
	if err != nil {
		return err, tag, "增加tag失败"
	}
	blog.AddDynamic()
	return err, tag, "增加tag成功"
}

//删除公告
func DeleteNotice(c *gin.Context) (msg string, err error) {
	db := global.GVA_DB
	id := c.Query("id")
	if id == "" {
		return "参数缺失", err
	}
	var notice model.BlogNotice
	err = db.Where("ID=?", id).Delete(&notice).Error
	if err != nil {
		return "删除公告失败", err
	}
	return "删除公告失败", err
}
