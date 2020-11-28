package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type BaiduResponse struct {
	Success interface{} `json:"success"`
	Remain interface{} `json:"remain"`
}
