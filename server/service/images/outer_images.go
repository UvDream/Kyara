package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/kirinlabs/HttpRequest"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
)

//获取token
func GetImagesToken(r request.ImagesStruct) (msg string, err error) {
	db := global.GVA_DB
	sys := model.SysConfig{}
	err = db.Find(&sys).Error
	if sys.ImgurType == "0" {
		fmt.Println("如优")
	} else if sys.ImgurType == "1" {
		fmt.Println("白熊图床")
		msg, err = getTokenBX(r)
	}
	return msg, err
}

//白熊图床获取token
func getTokenBX(r request.ImagesStruct) (data string, er error) {
	url := "https://pic.baixiongz.com/api/token"
	param := req.Param{
		"email":    r.UserName,
		"password": r.Password,
	}
	c, err := req.Post(url, param)
	var msg response.BxTokenResponse
	c.ToJSON(&msg)
	db := global.GVA_DB
	sys := model.SysConfig{}
	err = db.Model(&sys).Update(model.SysConfig{ImgurToken: msg.Data.Token, ImgurType: r.Type, ImgurUserName: r.UserName, ImgurPassword: r.Password}).Error
	if err != nil {
		return "更新token失败", err
	}
	return msg.Data.Token, err
}

//---------------------------------------获取图片列表-----------------------------------
func GetImagesList(r request.ImagesListStruct) (msg interface{}, err error) {
	//获取用户配置
	db := global.GVA_DB
	sys := model.SysConfig{}
	err = db.Find(&sys).Error
	if err != nil {
		return "获取错误", err
	}
	if sys.ImgurType == "0" {
		fmt.Println("如优")
		msg, err = getRyImgurList(r, sys.ImgurToken)
		return msg, err
	} else if sys.ImgurType == "1" {
		fmt.Println("白熊图床")
		msg, err = getBxImgurList(r, sys.ImgurToken)
		return msg, err
	}
	return "", err
}

//获取白熊图床图片列表
func getBxImgurList(r request.ImagesListStruct, token string) (data response.BxResponse, err error) {
	url := "https://pic.baixiongz.com/api/images"
	header := req.Header{
		"token": token,
	}
	param := req.Param{
		"page": r.Page,
		"rows": r.Limit,
	}
	c, err := req.Post(url, header, param)
	var msg response.BxResponse
	c.ToJSON(&msg)
	return msg, err
}
func getRyImgurList(r request.ImagesListStruct, token string) (list interface{}, err error) {
	//用token获取列表
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{"Content-Type": "application/json"})
	res, _ := req.Post("https://img.rruu.net/api/imageList", map[string]interface{}{
		"token":  token,
		"page":   r.Page,
		"limit":  r.Limit,
		"folder": r.Folder,
	})
	body, err := res.Body()
	if err != nil {
		return "获取如优图床错误", err
	}
	list = gojsonq.New().FromString(string(body)).Find("data")
	defer res.Close()
	return list, err
}

//文件上传
func UploadImage(c *gin.Context) (msg string, err error, data request.BxData) {
	db := global.GVA_DB
	sys := model.SysConfig{}
	err = db.Find(&sys).Error
	fmt.Println(err)
	if sys.ImgurType == "0" {
		fmt.Println("如优")
	} else if sys.ImgurType == "1" {
		//白熊图床
		msg, err, data = uploadBx(c, sys.ImgurToken)
		if err != nil {
			return "白熊图床上传失败", err, data
		}
		var imageList model.ExaFileUploadAndDownload
		imageList.Key = data.MD5
		imageList.Type = c.Query("type")
		imageList.Url = data.URL
		imageList.Name = data.Name
		err = db.Create(&imageList).Error
		if err != nil {
			return "关联图床到系统失败", err, data
		}
	}
	return msg, err, data
}

//白熊图床上传
func uploadBx(c *gin.Context, token string) (msg string, err error, list request.BxData) {
	file, path, err := c.Request.FormFile("image")
	//url:="https://pic.baixiongz.com/api/upload"
	//header := req.Header{
	//	"token": token,
	//}
	//r,err:=req.Post(url,header,req.FileUpload{
	//	File: file,
	//	FieldName: "image",
	//})
	//fmt.Println("查看错误")
	//fmt.Println(r,err)
	//fmt.Println("查看错误")
	//if err!=nil{
	//	return "上传失败", nil
	//}

	url := "https://pic.baixiongz.com/api/upload"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part1, errFile1 := writer.CreateFormFile("image", path.Filename)
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return "文件参数写入错误", err, list
	}
	err = writer.Close()
	if err != nil {
		return "文件写入失败", err, list
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "上传失败", err, list
	}
	req.Header.Add("token", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "请求上传失败", err, list
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "获取返回结果失败", err, list
	}
	var bxData request.BxRequest
	err = json.Unmarshal(body, &bxData)
	if bxData.Code != 200 {
		return "获取返回解析错误", err, list
	}
	list = bxData.Data
	err,msg=SaveImgUrl(list,c)
	if err!=nil {
		return msg, nil, list
	}
	return "上传成功", err, list
}
func SaveImgUrl(list request.BxData ,c *gin.Context)(err error ,msg string){
	db := global.GVA_DB
	var imgList model.ExaFileUploadAndDownload
	imgList.Url=list.URL
	imgList.Name=list.Name
	imgList.Tag=list.Mime
	imgList.Type=c.PostForm("type")
	err=db.Create(&imgList).Error
	if err!=nil {
		return err,"存储图片路径错误"
	}
	return err,"存储成功"
}
