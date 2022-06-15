package request

// PaginationRequest 分页通用请求
type PaginationRequest struct {
	KeyWord  string `json:"key_word"`
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
}
