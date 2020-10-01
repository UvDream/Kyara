package request

import "server/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}