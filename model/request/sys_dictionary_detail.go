package request

import "blog-api/model"

type SysDictionaryDetailSearch struct{
    model.SysDictionaryDetail
    PageInfo
}