package model

import (
	"github.com/jinzhu/gorm"
	"server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server
}

// 面试题库结构体
type Interview struct {
	gorm.Model
	//题目
	Title string `json:"title" gorm:"comment:'题目';type:text"`
	//tag
	Tag []uint `json:"tag" gorm:"-"`
	//tag 集合
	TagArr []string `json:"tag_arr" gorm:"-"`
	//	答案
	AnswerMd string `json:"answer_md" gorm:"comment:'答案markdown格式';type:longblob"`
	// 难以程度
	Level string `json:"level" gorm:"comment:'难易程度';default:'1'"`
	//分类
	ClassifyID string `json:"classify_id" gorm:"comment:'分类id'"`
}

//面试题分类
type InterviewClassify struct {
	gorm.Model
	ClassifyName string `json:"classify_name" gorm:"comment:'分类名称'"`
	ClassifyIcon string `json:"classify_icon" gorm:"comment:'分类图标'"`
	Count      uint   `json:"count" gorm:"-"`
}

//面试题tag集合
type InterviewTag struct {
	ID          uint `gorm:"primary_key"`
	InterviewID uint `json:"interview_id"`
	TagID       uint `json:"tag_id"`
}
