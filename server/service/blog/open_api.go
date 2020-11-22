package blog

import (
	"fmt"
	"github.com/imroc/req"
	"server/model/request"
	"server/model/response"
)

func RandomAvatar(r request.AvatarResponse)(err error,img string,msg string)  {
	url:="https://api.uomg.com/api/rand.avatar?sort="+r.Sort+"&format="+r.Format
	res,err:=req.Get(url)
	fmt.Print(res,err)
	if err!=nil{
		return err,"","获取图片失败"
	}
	var avatar response.AvatarResponse
	res.ToJSON(&avatar)
	return  err,avatar.Imgurl,"获取图片成功"
}
