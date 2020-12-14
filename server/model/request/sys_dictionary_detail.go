package request

import "server/model"

type SysDictionaryDetailSearch struct {
	model.SysDictionaryDetail
	PageInfo
}
