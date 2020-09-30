package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
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
