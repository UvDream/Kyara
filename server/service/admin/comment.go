package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/request"
)

func GetBlogComment(r request.ListStruct)(err error,comment []model.BlogComment,msg string,total int)  {
	fmt.Print(r)
	db := global.GVA_DB
	offset:=r.PageSize*(r.Page-1)
	err=db.Limit(r.PageSize).Offset(offset).Where("comment_content LIKE ?","%"+r.Search+"%").Find(&comment).Error
	if err!=nil {
		return err,comment,"查询留言失败",total
	}
	err=db.Where("comment_content LIKE ?","%"+r.Search+"%").Find(&comment).Count(&total).Error
	if err!=nil{
		return  err,comment,"查询总数失败",total
	}
	return err,comment,"查询留言成功",total
}
