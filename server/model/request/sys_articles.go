package request

// 文章接口结构体
type ArticleListStruct struct {
	ArticeID         string `json:"articleId"`
	Search           string `json:"search"`
	Page             int    `json:"page" form:"page"`
	PageSize         int    `json:"pageSize" form:"pageSize"`
	ClassificationID string `json:"classification_id"` //分类
	Tag              string `json:"tag"`
}

// 获取图床token结构体
type ImagesStruct struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// 如优获取图床图片列表结构体
type ImagesListStruct struct {
	Page   string `json:"page"`
	Limit  string `json:"limit"`
	Folder string `json:"folder"`
}

//正常分页结构体
type ListStruct struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Search   string `json:"search"`
	ArticleID string `json:"article_id"`
}

// 百度收录结构体
type ToBaiDuRequest struct {
	Site     string `json:"site"`
	Token    string `json:"token"`
	URL      string `json:"url"`
	Argument string `json:"argument"`
}
