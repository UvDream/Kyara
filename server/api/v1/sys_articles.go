package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	resp "gin-vue-admin/model/response"

	"gin-vue-admin/model/request"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)
// @title    Article
// @description   article , 文章列表
// @auth                     （2020/04/05  20:22）
// @param     u               model.ArticleList
// @return    err             error
// @return    userInter       *SysArticleListResponse
func ArticleList(c *gin.Context)  {
	var R request.ArticleListStruct
	_ = c.ShouldBindJSON(&R)
	 err,msg,Total:=service.ArticleList(R)
	 if err!=nil{
	 }else {
		 response.OkDetailed(resp.SysArticleListResponse{Msg:msg,TotalCount:Total},"获取成功",c)
	 }
}
//获取文章详情
func GetArticleDetail(c *gin.Context)  {
	id:=c.Query("id")
	password:=c.Query("password")
	error,message:=service.CheckPassword(id,password)
	if !error{
		response.FailWithMessage(message, c)
		return
	}
	if id=="" {
		response.FailWithMessage("请携带文章id", c)
		return
	}
	err,msg:=service.GetArticleDetail(c)
	if err==nil{
		response.OkDetailed(msg, "获取详情成功", c)
	}
}
//验证密码是否正确
func CheckPassword(c *gin.Context)  {
	id:=c.Query("id")
	password:=c.Query("password")
	err,message:=service.CheckPassword(id,password)
	fmt.Println(err)
	if err{
		response.OkDetailed("",message,c)
	}else{
		response.FailWithMessage(message, c)
	}
}
