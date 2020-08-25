package response

import "gin-vue-admin/model"

type SysArticleListResponse struct{
	Data []model.SysArticle `json:"data"`
	TotalCount int64 `json:"totalCount"`
}
