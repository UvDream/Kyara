package request
type BxRequest struct {
	Code int64 `json:"code"`
	Msg string `json:"msg"`
	Data BxData `json:"data"`
}
type BxData struct{
	Name string `json:"name"`
	URL string `json:"url"`
	Size string `json:"size"`
	Mime string `json:"mime"`
	MD5 string `json:"md_5"`
}

