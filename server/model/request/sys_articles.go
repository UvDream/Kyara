package request

// 文章接口结构体
type ArticleListStruct struct{
	ArticeID string `json:"articleId"`
	Search string `json:"search"`
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
// 获取图床token结构体
type ImagesStruct struct{
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Type string `json:"type"`
}
// 如优获取图床图片列表结构体
type ImagesListStruct struct {
	Page string `json:"page"`
	Limit string `json:"limit"`
	Folder string `json:"folder"`
}