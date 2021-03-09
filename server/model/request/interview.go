package request
type InterviewSearch struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	Level string `json:"level"` //等级
	Tag string `json:"tag"`
	LevelSort string `json:"level_sort"` //asc desc 正序逆序
	Classify string `json:"classify"` //分类
}
