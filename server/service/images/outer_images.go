package service

import (
	"fmt"
	"gin-vue-admin/model/request"
	"github.com/kirinlabs/HttpRequest"
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
		"email":r.Email,
		"password":r.Password,
	})
	body, err := res.Body()

	if res.StatusCode()==200 {
		fmt.Println(res,err,body)
		fmt.Println(res.Json(body))
	}
	return string(body)
}
