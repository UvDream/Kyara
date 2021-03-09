package initialize

import (
	"server/global"
	"server/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
		model.SysOperationRecord{},
		model.SysArticle{},
		model.SysComment{},
		model.SysTag{},
		model.SysClassify{},
		model.SysConfig{},
		model.Rewards{},
		model.SysLike{},
		model.SysArticleTag{},
		model.BlogNotice{},
		model.BlogView{},
		model.BlogComment{},
		model.CollectionCode{},
		model.BlogDynamic{},
		model.Interview{},//面试题
		model.InterviewClassify{},//面试题分类
		model.InterviewTag{},//面试题tag
	)
	global.GVA_LOG.Debug("注册成功!")
}
