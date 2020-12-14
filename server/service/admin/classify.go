package service

import (
	"server/global"
	"server/model"
	"server/model/response"
	"server/service/blog"

	"github.com/gin-gonic/gin"
)

func EditClassify(r response.EditClassifyResponse) (err error, msg string, classify model.SysClassify) {
	db := global.GVA_DB
	blog.AddDynamic()
	classify.TypeName = r.TypeName
	classify.Icon = r.Icon
	classify.ParentID = r.ParentID
	//新增
	if r.ID == 0 {
		err = db.Where("type_name = ?", r.TypeName).Find(&classify).Error
		if err == nil {
			return err, "名称重复", classify
		}

		err = db.Create(&classify).Error
		if err != nil {
			return err, "新增失败", classify
		}
	} else {
		classify.ID = r.ID
		err = db.Where("ID=?", r.ID).Save(&classify).Error
		if err != nil {
			return err, "更新失败", classify
		}
	}
	return err, "success", classify
}
func DeleteClassify(c *gin.Context) (err error, msg string) {
	db := global.GVA_DB
	blog.AddDynamic()
	var classify model.SysClassify
	id := c.Query("id")
	err = db.Where("parent_id = ?", id).Find(&classify).Error
	if err == nil {
		return err, "存在子目录"
	}
	err = db.Where("ID = ?", id).Find(&classify).Error
	if err != nil {
		return err, "不存在该目录"
	}
	//先查是否有元素存在下级
	err = db.Where("ID = ?", id).Unscoped().Delete(&model.SysClassify{}).Error
	if err != nil {
		return err, "删除失败"
	}
	return err, "删除成功"
}

func GetIconfontClassify() {

}
