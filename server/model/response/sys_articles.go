package response

import "server/model"

type SysArticleListResponse struct{
	Msg []model.SysArticle `json:"msg"`
	TotalCount int64 `json:"totalCount"`
}
