package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model"
)

//新增文章
func AddArticle(r model.SysArticle,c *gin.Context) (err error, msg string,data model.SysArticle) {
	db := global.GVA_DB
	//claims, _ := c.Get("claims")
	//waitUse := claims.(*request.CustomClaims)
	//r.UserID=waitUse.UUID
	//处理置顶问题
	if r.CoverType =="2"{
		r.Top="1"
	}
	if r.ImgURL==""{
		r.ImgURL="https://gitee.com/UvDream/images/raw/master/images/20200826195538.png"
	}
	if r.UserID==""{
		r.UserID="ce0d6685-c15f-4126-a5b4-890bc9d2356d"
	}

	if r.ArticleID == 0 {
		//增加文章
		err := db.Create(&r).Error
		fmt.Println(r)
		fmt.Println(err)
		if err != nil {
			return err, "文章创建失败",r
		}else{
			return err,"文章创建成功",r
		}
	} else {
		//查询文章是否存在
		err := db.Where("ID=?", r.ArticleID).Find(&model.SysArticle{}).Error
		if err != nil {
			return err, "文章不存在",r
		}
		err = db.Model(&model.SysArticle{}).Where("ID=?", r.ArticleID).Update(&r).Error
		if err != nil {
			return err, "更新失败",r
		} else {
			return err, "更新文章成功",r
		}
	}
	return nil, "",r
}
//查询文章详情
func GetArticleDetail(id string)(err error,article model.SysArticle)  {
	article =model.SysArticle{}
	db := global.GVA_DB
	err = db.Where("id=? ", id).Find(&article).Error
	return err,article
}

//新增公告
func AddNotice(r model.BlogNotice) (err error,msg string) {
	db := global.GVA_DB

	if r.Show=="1"{
		updateNotice()
	}
	if r.ID!=0{
		err=db.Model(&model.BlogNotice{}).Update(&r).Error
	}else{
		err = db.Create(&r).Error
		if err!=nil{
			return err,"新增失败"
		}
	}
	if r.Show=="1"{
		err=db.Model(&model.SysConfig{}).Update("blog_notice_id",r.ID).Error
		if err!=nil {
			return  err,"更新配置失败"
		}
	}
	return err,"成功"
}
func updateNotice()(err error,msg string)  {
	var notice []model.BlogNotice
	db := global.GVA_DB
	err=db.Find(&notice).Error
	if err==nil {
		for i,k:=range notice{
			if k.Show=="1" {
				notice[i].Show="0"
			}
		}
		db.Model(&notice).UpdateColumn("show", "0")
	}else{
		return err,"更新数据失败"
	}
	return err,"更新成功"
}
//新增tag
func AddTag(c *gin.Context)(err error,tag model.SysTag,msg string)  {
	t:=c.Query("tag")
	db := global.GVA_DB
	tag.TagName=t
	//需要查重
	err=db.Where("tag_name=?",tag).Find(&tag).Error
	if err!=nil {
		return err,tag,"tag重复"
	}
	err=db.Create(&tag).Error
	if err!=nil {
		return err,tag,"增加tag失败"
	}
	return err,tag,"增加tag成功"
}