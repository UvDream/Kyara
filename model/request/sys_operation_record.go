package request

import "blog-api/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
