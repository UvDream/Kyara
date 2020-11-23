package blog

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/request"
)

func Comment(r model.BlogComment)(err error ,msg string)  {
	db := global.GVA_DB
	r.Status="0"
	if r.ID==0{
		err=db.Create(&r).Error
		if err!=nil {
			return err,"留言入库失败"
		}
	}

	return err,"留言成功"
}
func GetComment(r request.ListStruct)(err error,msg string, blogComment []model.BlogComment,totalCount int64)  {
	db := global.GVA_DB
	offset:=r.PageSize*(r.Page-1)

	err=db.Where("status=? AND parent_id=?","1","").Limit(r.PageSize).Offset(offset).Find(&blogComment).Count(&totalCount).Error
	if err!=nil{
		return err,"查询失败",blogComment,0
	}
	fmt.Println(blogComment)
	for k,i:=range blogComment{
		blogComment[k].Children=findChildren(i.ID)
	}
	return err,"查询留言成功",blogComment,totalCount
}
func findChildren(parentId uint)(child []model.BlogComment)  {
	db := global.GVA_DB
	err:=db.Where("parent_id=? AND status=?",parentId,"1").Find(&child).Error
	fmt.Println(err)
	for k, i := range child{
		child[k].Children = findChildren(i.ID)
	}
	return child
}
