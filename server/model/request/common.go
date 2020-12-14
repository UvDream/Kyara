package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
type DynamicRequest struct {
	List     [][]string `json:"list" form:"list"`
	Today    string     `json:"today" form:"today"`
	ThatDay  string     `json:"thatDay" form:"thatDay"`
	MaxCount int        `json:"maxCount" form:"maxCount"`
}
