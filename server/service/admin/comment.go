package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/service/blog"
)

//获取博客留言
func GetBlogComment(r request.ListStruct) (err error, comment []model.BlogComment, msg string, total int) {
	db := global.GVA_DB
	offset := r.PageSize * (r.Page - 1)
	err = db.Limit(r.PageSize).Offset(offset).Where("comment_content LIKE ?", "%"+r.Search+"%").Find(&comment).Error
	if err != nil {
		return err, comment, "查询留言失败", total
	}
	err = db.Where("comment_content LIKE ?", "%"+r.Search+"%").Find(&comment).Count(&total).Error
	if err != nil {
		return err, comment, "查询总数失败", total
	}
	return err, comment, "查询留言成功", total
}

//审核博客留言
func CheckBlogComment(r response.CheckComment) (err error, msg string) {
	db := global.GVA_DB
	var comment model.BlogComment
	err = db.Where("ID=?", r.ID).Find(&comment).Error
	fmt.Println(err, comment)
	//审核
	status := "1"
	if r.Check {
		status = "1"
	} else {
		status = "0"
	}
	err = db.Model(&comment).Where("ID=?", r.ID).Update("status", status).Error
	if err != nil {
		return err, "审核更新数据库失败"
	}
	blog.AddDynamic()
	return err, "审核完成"
}
