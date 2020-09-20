package response
type BxResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data bxMsg `json:"data"`
}
type bxMsg struct {
	Total int `json:"total"`
	CurrentPage int `json:"current_page"`
	Data []bxData `json:"data"`
}
type bxData struct {
	id int `json:"id"`
	URL string `json:"url"`
	UploadTime string `json:"upload_time"`
	UploadDate string `json:"upload_date"`
	Name string `json:"name"`
}