package request

import "blog-api/model"

type SysDictionarySearch struct{
    model.SysDictionary
    PageInfo
}