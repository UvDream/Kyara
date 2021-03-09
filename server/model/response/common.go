package response

import "server/model"

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type BaiduResponse struct {
	Success interface{} `json:"success"`
	Remain  interface{} `json:"remain"`
}

type InterviewResponse struct {
	Msg string `json:"msg"`
	TotalCount int64 `json:"total_count"`
	Data []model.Interview `json:"data"`
}