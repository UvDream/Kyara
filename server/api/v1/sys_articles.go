package v1

import (
	"fmt"
	"server/global/response"
	resp "server/model/response"

	"server/model/request"
	"server/service"
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
	 err,msg,Total,errMsg:=service.ArticleList(R)
	 if err!=nil{
	 	fmt.Println("获取文章列表错误")
	 	response.FailWithMessage(errMsg,c)
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
// 文章分类
func ArticleClassify(c *gin.Context)  {
	err,msg:=service.ArticleClassify()
	fmt.Println(err,msg)
	if err!=nil{
		response.FailWithMessage("获取分类失败", c)
	}else{
		response.OkDetailed(msg, "获取文章分类成功", c)
	}
}
//热门文章
func HotArticle(c *gin.Context){
 err,list:=service.HotArticle()
 if err!=nil{
	 response.FailWithMessage("获取文章失败", c)
 }else{
	 response.OkDetailed(list, "获取文章成功", c)
 }
}
