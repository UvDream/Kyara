package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/kirinlabs/HttpRequest"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
)
type res struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
//获取token
func GetImagesToken( r request.ImagesStruct)(msg string ,err error) {
	db := global.GVA_DB
	sys:=model.SysConfig{}
	err=db.Find(&sys).Error
	if sys.ImgurType == "0"{
		fmt.Println("如优")
	}else if sys.ImgurType=="1"{
		fmt.Println("白熊图床")
		msg,err=getTokenBX(r)
	}
	return msg,err
}
//白熊图床获取token
func getTokenBX(r request.ImagesStruct)(data string ,er error){
	url:="https://pic.baixiongz.com/api/token"
	param := req.Param{
		"email":  r.UserName,
		"password": r.Password,
	}
	c,err:=req.Post(url,param)
	var msg response.BxTokenResponse
	c.ToJSON(&msg)
	db := global.GVA_DB
	sys:=model.SysConfig{}
	err=db.Model(&sys).Update(model.SysConfig{ImgurToken:msg.Data.Token,ImgurType:r.Type,ImgurUserName: r.UserName,ImgurPassword: r.Password}).Error
	if err!=nil {
		return "更新token失败",err
	}
	return msg.Data.Token,err
}
//---------------------------------------获取图片列表-----------------------------------
func GetImagesList(r request.ImagesListStruct)(msg interface{},err error)  {
	//获取用户配置
	db := global.GVA_DB
	sys:=model.SysConfig{}
	err=db.Find(&sys).Error
	if err!=nil{
		return "获取错误",err
	}
	if sys.ImgurType == "0"{
		fmt.Println("如优")
		msg,err=getRyImgurList(r,sys.ImgurToken)
		return msg,err
	}else if sys.ImgurType=="1"{
		fmt.Println("白熊图床")
		 msg,err=getBxImgurList(r,sys.ImgurToken)
		return msg,err
	}
	return "",err
}


//获取白熊图床图片列表
func getBxImgurList(r request.ImagesListStruct,token string) (data response.BxResponse ,err error)  {
	url:="https://pic.baixiongz.com/api/images"
	header := req.Header{
		"token": token,
	}
	param := req.Param{
		"page":  r.Page,
		"rows": r.Limit,
	}
	c, err := req.Post(url, header, param)
	var  msg response.BxResponse
	c.ToJSON(&msg)
	return msg,err
}
func getRyImgurList(r request.ImagesListStruct,token string)(list interface{},err error)  {
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
		return "获取如优图床错误",err
	}
	list=gojsonq.New().FromString(string(body)).Find("data")
	return list,err
}

//文件上传
func UploadImage(c *gin.Context)(msg string,err error)  {
	db := global.GVA_DB
	sys:=model.SysConfig{}
	err=db.Find(&sys).Error
	fmt.Println(err)
	if sys.ImgurType == "0"{
		fmt.Println("如优")
	}else if sys.ImgurType=="1"{
		msg,err=uploadBx(c,sys.ImgurToken)
	}
	return msg,err
}
//白熊图床上传
func uploadBx(c *gin.Context,token string) (msg string,err error) {
	file,a,s:= c.Request.FormFile("image")
	fmt.Println(file,a,s)
	url:="https://pic.baixiongz.com/api/upload"
	authHeader := req.Header{
		"Accept":"multipart/form-data",
		"token": token,
	}
	r,err:=req.Post(url, authHeader,req.FileUpload{
		File:      file,
		FieldName: "image",
	})
	fmt.Println(r,err)
	if err!=nil{
		return "上传失败", nil
	}
	return "上传成功",err
}