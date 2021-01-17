package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/service/blog"
	"strconv"
)

//获取博客留言
func GetBlogComment(r request.ListStruct) (err error, comment []model.BlogComment, msg string, total int) {
	db := global.GVA_DB
	offset := r.PageSize * (r.Page - 1)
	err = db.Limit(r.PageSize).Offset(offset).Where("comment_content LIKE ?", "%"+r.Search+"%").Find(&comment).Error
	if err != nil {
		return err, comment, "查询留言失败", total
	}
	for i,k:=range comment{
		if k.UserID!=""{
			var user model.SysUser
			err=db.Where("ID=?",k.UserID).Find(&user).Error
			if err!=nil {
				return err, comment, "查询回复信息失败", total
			}
			comment[i].UserName=user.NickName
			comment[i].Email=user.Email
			comment[i].BlogURL=user.BlogURL
			comment[i].Avatar=user.HeaderImg
		}
	}
	err = db.Where("comment_content LIKE ?", "%"+r.Search+"%").Where("deleted_at is null").Table("blog_comments").Count(&total).Error
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

//回复留言
func RevertComment(c *gin.Context, r response.ReplyComment) (msg string, err error) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	var comment model.BlogComment
	comment.CommentContent = r.Comment
	comment.ParentID = r.ID
	comment.UserName = waitUse.NickName
	comment.UserID = strconv.Itoa(int(waitUse.ID ))
	comment.Status = "1"
	db := global.GVA_DB
	err = db.Create(&comment).Error
	if err != nil {
		return "回复留言失败", err
	}
	// 查询用户信息
	return "回复成功", err
}

//删除留言
func DeleteComment(c *gin.Context) (msg string, err error) {
	id := c.Query("id")
	if id == "" {
		return "参数缺失", err
	}
	db := global.GVA_DB
	var comment model.BlogComment
	err = db.Where("parent_id=?", id).Delete(&comment).Error
	if err != nil {
		return "删除子节点失败", err
	}
	err = db.Where("ID=?", id).Delete(&comment).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", err
}
