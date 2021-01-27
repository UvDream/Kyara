package request

import "time"
// 博客分类
type Actives struct{
	Time time.Time `json:"time"`
	List []ActivesItem `json:"list"`
}
type ActivesItem struct{
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
}