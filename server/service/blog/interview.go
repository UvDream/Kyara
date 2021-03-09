package blog

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/request"
)
//新增面试题分类
func AddInterviewClassify(r model.InterviewClassify)(err error,msg string) {
	fmt.Println(r)
	db := global.GVA_DB
	//先校验名称
	err=db.Where("classify_name=?",r.ClassifyName).Find(&r).Error
	if err==nil {
		return err,"该名称已经存在"
	}
	err=db.Create(&r).Error
	if err!=nil {
		return err,"创建分类失败"
	}
	return err,"创建分类成功"
}
//获取面试题分类
func GetInterviewClassify() (err error, msg string, data []model.InterviewClassify) {
	db := global.GVA_DB
	err = db.Find(&data).Error
	if err != nil {
		return err, "获取分类失败", data
	}
	return err, "获取分类成功", data
}

//获取面试题
func GetInterview(r request.InterviewSearch) (err error, msg string, interview []model.Interview, totalCount int64) {
	db := global.GVA_DB
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	//获取面试题总数
	err = db.Find(&interview).Count(&totalCount).Error
	if err != nil {
		return err, "获取题库", interview, 0
	}
	//获取面试题
	err = db.Where("level LIKE ?", "%"+r.Level+"%").Where("classify = ?", r.Classify).Order("level " + r.LevelSort).Find(&interview).Limit(limit).Offset(offset).Error

	if err != nil {
		return err, "获取题库失败", interview, totalCount
	}
	return err, "获取题库成功", interview, totalCount
}

func GetInterviewDetail(id string)(err error,msg string,data model.Interview)  {
	db := global.GVA_DB
	err=db.Where("classify_id=?",id).Find(&data).Error
	if err!=nil {
		return err,"获取详情失败",data
	}
	//还需要获取tag
	var tag []model.InterviewTag
	err=db.Where("interview_id=?",data.ID).Find(&tag).Error
	if err!=nil {
		return err,"获取tag失败",data
	}

	for _,k:=range tag{
		data.Tag=append(data.Tag,k.TagID)
		var tagDetail model.SysTag
		err=db.Where("id=?",k.TagID).Find(&tagDetail).Error
		if err!=nil {
			return err,"获取tag中文名失败",data
		}
		data.TagArr=append(data.TagArr,tagDetail.TagName)
	}
	return err,"获取成功",data
}

//新增面试题
func AddInterview(r model.Interview)(err error ,msg string)  {
	db := global.GVA_DB
	if r.Level=="" {
		r.Level="1"
	}
	err=db.Create(&r).Error
	if err!=nil {
		return  err,"创建失败"
	}
	for _,k:=range r.Tag{
		var tag model.InterviewTag
		tag.TagID=k
		tag.InterviewID=r.ID
		err=db.Where("interview_id=? AND tag_id=?",tag.TagID,tag.InterviewID).Error
		if err==nil {
			err=db.Create(&tag).Error
			if err!=nil{
				return err,"创建试题和tag之间关系失败"
			}
		}

	}
	return  err,"创建试题成功"
}
