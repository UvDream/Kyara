package response

import "server/model"

//白熊图床相关返回结构体
type BxResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data bxMsg  `json:"data"`
}
type bxMsg struct {
	Total       int      `json:"total"`
	CurrentPage int      `json:"current_page"`
	Data        []bxData `json:"data"`
}
type bxData struct {
	ID int `json:"id"`
	//地址
	URL string `json:"url"`
	//上传时间
	UploadDate string `json:"upload_date"`
	//名称
	Name string `json:"name"`
	//大小
	Size  string `json:"size"`
	Token string `json:"token"`
}

//白熊图床token
type BxTokenResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data bxData `json:"data"`
}

type CommentMsg struct {
	Total int64               `json:"total"`
	Data  []model.BlogComment `json:"data"`
}
