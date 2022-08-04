package article

import (
	"github.com/gin-gonic/gin"
	"server/code"
	article2 "server/model/article"
	"server/model/article/request"
	"server/model/common/response"
	"server/utils"
)

type ArticlesApi struct {
}

// CreateArticle 创建文章
func (article *ArticlesApi) CreateArticle(c *gin.Context) {
	var articleOpts request.ArticleRequest
	//校验必填信息
	err := c.ShouldBindJSON(&articleOpts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	articleContent, msg, err := articleService.CreateArticle(articleOpts)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(articleContent, msg, c)
}

// DeleteArticle 删除文章
func (article *ArticlesApi) DeleteArticle(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
	}
	msg, err := articleService.DeleteArticleService(id)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithMessage(msg, c)
}

// UpdateArticle 修改文章
func (article *ArticlesApi) UpdateArticle(c *gin.Context) {
	var articleOpts article2.Article
	//校验必填信息
	err := c.ShouldBindJSON(&articleOpts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if articleOpts.ID == "" {
		response.FailWithMessage("缺少必要参数uuid", c)
		return
	}
	articleContent, msg, err := articleService.UpdateArticleService(articleOpts)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(articleContent, msg, c)

}

// GetArticleList 查询文章
func (article *ArticlesApi) GetArticleList(c *gin.Context) {
	//x-token
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		response.FailWithMessage("token验证失败", c)
		return
	}
	var articleOpts request.ArticleListRequest
	err = c.ShouldBindQuery(&articleOpts)
	if err != nil {
		response.FailResponse(code.ErrorGetQueryParam, c)
		return
	}
	list, total, msg, err := articleService.GetArticleListService(articleOpts, claims.ID, c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"list":  list,
		"total": total,
	}, msg, c)
}

//GetArticleHistory 查询文章历史记录
func (article *ArticlesApi) GetArticleHistory(c *gin.Context) {
	uid := c.Query("id")
	if uid == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	list, msg, err := articleService.GetArticleHistoryService(uid)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(gin.H{
		"list": list,
	}, "查询成功", c)
}

//GetArticleDetail 查询文章详情
func (article *ArticlesApi) GetArticleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	articleContent, msg, err := articleService.GetArticleDetailService(id)
	if err != nil {
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(articleContent, msg, c)
}
