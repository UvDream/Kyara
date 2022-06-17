package article

import (
	"server/global"
	"server/model/article"
)

type ToTagService struct{}

//CreateTagService 创建tag
func (t *ToTagService) CreateTagService(data article.Tag) (tag *article.Tag, msg string, err error) {
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
func (t *ToTagService) DeleteTagService() {

}

//UpdateTagService 更新tag
func (t *ToTagService) UpdateTagService() {

}

//GetTagsService 获取tag列表
func (t *ToTagService) GetTagsService() {

}
