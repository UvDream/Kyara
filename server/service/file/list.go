package file

import (
	code2 "server/code"
	"server/global"
	"server/models"
	"server/models/file"
)

func (*FilesService) ListFileService(query models.PaginationRequest, uuid string) (list []file.File, total int64, code int, err error) {
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	db := global.DB.Model(file.File{})
	if query.KeyWord != "" {
		db = db.Where("name like ?", "%"+query.KeyWord+"%").Or("path like ?", "%"+query.KeyWord+"%").Or("url like ?", "%"+query.KeyWord+"%").Or("key like ?", "%"+query.KeyWord+"%").Or("type like ?", "%"+query.KeyWord+"%").Or("position like ?", "%"+query.KeyWord+"%")
	}
	db = db.Where("auth_id = ?", uuid)
	if err = db.Count(&total).Error; err != nil {
		return list, total, code2.ErrorListFile, err
	}
	if err = db.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return list, total, code2.ErrorListFile, err
	}
	return list, total, code2.SUCCESS, err
}
