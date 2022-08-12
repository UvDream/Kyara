package request

// PaginationRequest 分页通用请求
type PaginationRequest struct {
	KeyWord  string `json:"key_word" form:"key_word"` // 关键字
	Page     int    `json:"page" binding:"required" form:"page"`
	PageSize int    `json:"page_size" binding:"required" form:"page_size"`
}
