package request

import "server/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
