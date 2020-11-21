package blog

import (
	"fmt"
	"server/global"
	"server/model"
)

func Comment(r model.BlogComment)  {
	db := global.GVA_DB
	err:=db.Create(&r).Error
	fmt.Println(err)
}