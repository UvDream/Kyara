package service

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"github.com/kirinlabs/HttpRequest"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"net/http"
	"strings"
)
type res struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
//获取token
func GetImagesToken( r request.ImagesStruct)(interface{})  {
	fmt.Println(r)
	req:=HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{"Content-Type": "application/json"})
	fmt.Println(req)
	res,_:=req.Post("https://pic.baixiongz.com/api/token",map[string]interface{}{
		"email":r.UserName,
		"password":r.Password,
	})
	body, err := res.Body()
	if err!=nil {
		return "获取token失败"
	}
	token:=gojsonq.New().FromString(string(body)).Find("data.token")
	db := global.GVA_DB
	sys:=model.SysConfig{}
	err=db.Model(&sys).Update(model.SysConfig{ImgurToken:token .(string),ImgurType:r.Type,ImgurUserName: r.UserName,ImgurPassword: r.Password}).Error
	if err!=nil {
		return "更新配置失败"
	}
	return token
}
//获取图片列表
func GetImagesList(r request.ImagesListStruct)(interface{})  {
	//获取用户配置
	db := global.GVA_DB
	sys:=model.SysConfig{}
	err:=db.Find(&sys).Error
	if err!=nil{
		return "获取错误"
	}
	if sys.ImgurType == "0"{
		fmt.Println("如优")
		return getRyImgurList(r,sys.ImgurToken)
	}else if sys.ImgurType=="1"{
		fmt.Println("白熊图床")
		return getBxImgurList(r,sys.ImgurToken)
	}
	return ""
}


//获取白熊图床图片列表
func getBxImgurList(r request.ImagesListStruct,token string) (data response.BxResponse)  {
	url:="https://pic.baixiongz.com/api/images"
	payload:=strings.NewReader("page="+r.Page+"&rows="+r.Limit)
	req,_:=http.NewRequest("POST",url,payload)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("token",token)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var msg response.BxResponse
	temp := []byte(body)
	json.Unmarshal(temp, &msg)
	fmt.Println(msg)
	//list:=gojsonq.New().FromString(string(body)).Find("data")
	return msg
}
func getRyImgurList(r request.ImagesListStruct,token string)(interface{})  {
	//用token获取列表
	req:=HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{"Content-Type": "application/json"})
	res,_:=req.Post("https://img.rruu.net/api/imageList",map[string]interface{}{
		"token":token,
		"page":r.Page,
		"limit":r.Limit,
		"folder":r.Folder,
	})
	body, err := res.Body()
	if err!=nil {
		return "获取如优图床错误"
	}
	list:=gojsonq.New().FromString(string(body)).Find("data")
	return list
}