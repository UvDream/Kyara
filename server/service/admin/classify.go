package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/response"
)

func EditClassify(r response.EditClassifyResponse)(err error,msg string,classify model.SysClassify)  {
	fmt.Println(r)
	db := global.GVA_DB

	classify.TypeName=r.TypeName
	classify.Icon=r.Icon
	classify.ParentID=r.ParentID
	//新增
	if r.ID==0 {
		err=db.Where("type_name = ?",r.TypeName).Find(&classify).Error
		if err==nil {
			return err,"名称重复",classify
		}

		err=db.Create(&classify).Error
		if err!=nil {
			return err,"新增失败",classify
		}
	}else{
		classify.ID=r.ID
		err=db.Where("ID=?",r.ID).Save(&classify).Error
		if err!=nil {
			return err,"更新失败",classify
		}
	}
	return err,"success",classify
}