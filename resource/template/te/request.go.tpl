package request

import "blog-api/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}