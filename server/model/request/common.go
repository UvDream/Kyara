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
type IconfontRequest struct {
	Code int64        `json:"code"`
	Data IconfontData `json:"data"`
}
type IconfontData struct {
	Icons []IconList `json:"icons"`
}
type IconList struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ProjectID int64  `json:"project_id"`
	ShowSvg   string `json:"show_svg"`
	FontClass string `json:"font_class"`
}
