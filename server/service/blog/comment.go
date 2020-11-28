package blog

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"net/http"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"strings"
	"time"
)

func Comment(r model.BlogComment)(err error ,msg string)  {
	db := global.GVA_DB
	r.Status="0"
	if r.ID==0{
		err=db.Create(&r).Error
		if err!=nil {
			return err,"留言入库失败"
		}
	}

	return err,"留言成功"
}
func GetComment(r request.ListStruct)(err error,msg string, blogComment []model.BlogComment,totalCount int64)  {
	db := global.GVA_DB
	offset:=r.PageSize*(r.Page-1)

	err=db.Where("status=? AND parent_id=?","1","").Limit(r.PageSize).Offset(offset).Find(&blogComment).Count(&totalCount).Error
	if err!=nil{
		return err,"查询失败",blogComment,0
	}
	fmt.Println(blogComment)
	for k,i:=range blogComment{
		blogComment[k].Children=findChildren(i.ID)
	}
	return err,"查询留言成功",blogComment,totalCount
}
func findChildren(parentId uint)(child []model.BlogComment)  {
	db := global.GVA_DB
	err:=db.Where("parent_id=? AND status=?",parentId,"1").Find(&child).Error
	fmt.Println(err)
	for k, i := range child{
		child[k].Children = findChildren(i.ID)
	}
	return child
}
//百度收录
func ToBaidu(r request.ToBaiDuRequest)(err error,msg string,baiduResponse response.BaiduResponse)  {
	var config model.SysConfig
	db := global.GVA_DB
	err=db.Find(&config).Error
	if  err!=nil{
		return err,"获取配置失败",baiduResponse
	}
	url:="http://data.zz.baidu.com/urls?site="+config.DomainName+"&token="+config.BaiDuToken
	contentType:="text/plain"
	data:=config.DomainName+r.Argument
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, contentType, strings.NewReader(data))
	resp.Header.Add("User-Agent", "curl/7.12.1")
	resp.Header.Add("Host", "data.zz.baidu.com")
	resp.Header.Add("Content - Type", "text / plain")
	resp.Header.Add("Content-Length", "83")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	success:=gojsonq.New().FromString(string(result)).Find("success")
	remain:=gojsonq.New().FromString(string(result)).Find("remain")
	baiduResponse.Success=success
	baiduResponse.Remain=remain
	return err,"百度收录成",baiduResponse
}
