package article

import (
	"server/global"
	"server/model/article"
)

type TagService struct{}

//CreateTagService 创建tag
func (t *TagService) CreateTagService(data article.Tag) (tag *article.Tag, msg string, err error) {
	db := global.DB
	//首先查询是否重复
	var tagItem article.Tag
	err = db.Where("name = ?", data.Name).First(&tagItem).Error
	if err == nil {
		return &tagItem, "tag已经存在", err
	}
	if err := db.Create(&data).Error; err != nil {
		return &data, "创建tag失败", err
	}
	return &data, "创建tag成功", err
}

//DeleteTagService 删除tag
func (t *TagService) DeleteTagService(id string) (msg string, err error) {
	//查询是否存在
	db := global.DB
	var tag article.Tag
	err = db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return "tag不存在", err
	}
	err = db.Where("id = ?", id).Delete(&tag).Error
	if err != nil {
		return "删除tag失败", err
	}
	return "删除tag成功", err
}

//UpdateTagService 更新tag
func (t *TagService) UpdateTagService(data article.Tag) (tag article.Tag, msg string, err error) {
	db := global.DB
	//	查询是否存在
	if err := db.Where("id = ?", data.ID).First(&article.Tag{}).Error; err != nil {
		return data, "tag不存在", err
	}
	if err := db.Where("id = ?", data.ID).Save(&data).Error; err != nil {
		return data, "tag更新失败", err
	}
	return data, "更新成功", err
}

//GetTagsService 获取tag列表
func (t *TagService) GetTagsService(keyword string) (tagList []article.Tag, msg string, err error) {
	db := global.DB
	db = db.Model(&article.Tag{})
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%").Or("slug LIKE ?", "%"+keyword+"%")
	}
	if err := db.Find(&tagList).Error; err != nil {
		return nil, "查询tag列表失败", err
	}
	return tagList, "查询tag列表成功", err
}
