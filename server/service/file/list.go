package file

import (
	"github.com/google/uuid"
	code2 "server/code"
	"server/global"
	"server/model/common/request"
	"server/model/file"
)

func (*FilesService) ListFileService(query request.PaginationRequest, uuid uuid.UUID) (list []file.File, total int64, code int, err error) {
	db := global.DB.Model(file.File{})
	if query.KeyWord != "" {
		db = db.Where("name like ?", "%"+query.KeyWord+"%").Or("path like ?", "%"+query.KeyWord+"%").Or("url like ?", "%"+query.KeyWord+"%").Or("key like ?", "%"+query.KeyWord+"%").Or("type like ?", "%"+query.KeyWord+"%").Or("position like ?", "%"+query.KeyWord+"%")
	}
	db = db.Where("auth_id = ?", uuid)
	if err = db.Count(&total).Error; err != nil {
		return list, total, code2.ErrorListFile, err
	}
	if err = db.Find(&list).Error; err != nil {
		return list, total, code2.ErrorListFile, err
	}
	return list, total, 0, err
}
