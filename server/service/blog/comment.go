package blog

import (
	"server/global"
	"server/model"
)

func Comment(r model.BlogComment)(err error ,msg string)  {
	db := global.GVA_DB
	if r.ID==0{
		err=db.Create(&r).Error
		if err!=nil {
			return err,"留言入库失败"
		}
	}

	return err,"留言成功"
}