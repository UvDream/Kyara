package request

// PaginationRequest 分页通用请求
type PaginationRequest struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}
