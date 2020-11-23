package blog

import (
	"fmt"
	"server/global"
	"server/model"
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
func GetComment()(err error,msg string, blogComment []model.BlogComment)  {
	db := global.GVA_DB
	err=db.Where("status=? AND parent_id=?","1","").Find(&blogComment).Error
	if err!=nil{
		return err,"查询失败",blogComment
	}
	fmt.Println(blogComment)
	for k,i:=range blogComment{
		blogComment[k].Children=findChildren(i.ID)
	}
	return err,"查询留言成功",blogComment
}
func findChildren(parentId uint)(child []model.BlogComment)  {
	db := global.GVA_DB
	err:=db.Where("parent_id=?",parentId).Find(&child).Error
	fmt.Println(err)
	for k, i := range child{
		child[k].Children = findChildren(i.ID)
	}
	return child
}
