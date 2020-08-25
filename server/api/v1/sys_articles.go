package v1

import (
	"gin-vue-admin/global/response"
	resp "gin-vue-admin/model/response"

	"gin-vue-admin/model/request"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)
//文章列表
func ArticleList(c *gin.Context)  {
	var R request.ArticleListStruct
	_ = c.ShouldBindJSON(&R)
	 err,msg,Total:=service.ArticleList(R)
	 if err!=nil{
	 }else {
		 response.OkDetailed(resp.SysArticleListResponse{Data:msg,TotalCount:Total},"获取成功",c)
	 }
}
