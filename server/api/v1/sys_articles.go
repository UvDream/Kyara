package v1

import (
	"fmt"
	"server/global/response"
	resp "server/model/response"
	"server/service/article"

	"github.com/gin-gonic/gin"
	"server/model/request"
)

// @title    Article
// @description   article , 文章列表
// @auth                     （2020/04/05  20:22）
// @param     u               model.ArticleList
// @return    err             error
// @return    userInter       *SysArticleListResponse
func ArticleList(c *gin.Context) {
	var R request.ArticleListStruct
	_ = c.ShouldBindJSON(&R)
	err, msg, Total, errMsg := article.ArticleList(R)
	if err != nil {
		fmt.Println("获取文章列表错误")
		response.FailWithMessage(errMsg, c)
	} else {
		response.OkDetailed(resp.SysArticleListResponse{Msg: msg, TotalCount: Total}, "获取成功", c)
	}
}

//获取文章详情
func GetArticleDetail(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	error, message := article.CheckPassword(id, password)
	if !error {
		response.FailWithMessage(message, c)
		return
	}
	if id == "" {
		response.FailWithMessage("请携带文章id", c)
		return
	}
	err, msg := article.GetArticleDetail(c)
	if err == nil {
		response.OkDetailed(msg, "获取详情成功", c)
	}
}

//验证密码是否正确
func CheckPassword(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	err, message := article.CheckPassword(id, password)
	fmt.Println(err)
	if err {
		response.OkDetailed("", message, c)
	} else {
		response.FailWithMessage(message, c)
	}
}

// 文章分类
func ArticleClassify(c *gin.Context) {
	err, msg := article.ArticleClassify()
	fmt.Println(err, msg)
	if err != nil {
		response.FailWithMessage("获取分类失败", c)
	} else {
		response.OkDetailed(msg, "获取文章分类成功", c)
	}
}

//热门文章
func HotArticle(c *gin.Context) {
	err, list := article.HotArticle()
	if err != nil {
		response.FailWithMessage("获取文章失败", c)
	} else {
		response.OkDetailed(list, "获取文章成功", c)
	}
}

//获取所有tag
func AllTag(c *gin.Context) {
	err, list, msg := article.AllTag()
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(list, msg, c)
	}
}
//获取博客配置
func GetConfig(c *gin.Context)  {
	err,data,msg:= article.GetConfig()
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(data,msg, c)
	}
}
//获取github配置
func GetGithub(c *gin.Context)  {
	msg,err:= article.GetGithub()
	if err != nil {
		response.FailWithMessage("获取github仓库失败", c)
	} else {
		response.OkDetailed(msg, "获取github仓库成功", c)
	}
	
}
//博客访问量统计
func ViewBlogCount(c *gin.Context)  {
	//id := c.Query("id")
	//global.GVA_LOG.Info(c.ClientIP())
	err,msg:= article.ViewBlogCount(c)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkWithMessage(msg,c)
	}
}
//获取公告
func GetNotice(c *gin.Context)  {
	var R request.ListStruct
	_ = c.ShouldBindJSON(&R)
	err,list,msg,total:= article.GetNotice(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(resp.BlogNotice{List: list,TotalCount: total}, msg, c)
	}
}