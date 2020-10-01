package request

import "server/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
