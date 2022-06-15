package article

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/model/article/request"
	"server/model/common/response"
)

type ArticlesApi struct {
}

// CreateArticle 创建文章
func (article *ArticlesApi) CreateArticle(c *gin.Context) {
	var articleOpts request.ArticleRequest
	//校验必填信息
	_ = c.ShouldBindJSON(&articleOpts)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//}
	articleOpts.UUID = uuid.New()
	articleContent, msg, err := articleService.CreateArticle(articleOpts)
	if err != nil {
		response.FailWithMessage(msg, c)
	}
	response.OkWithDetailed(articleContent, msg, c)
}

// DeleteArticle 删除文章
func (article *ArticlesApi) DeleteArticle(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// UpdateArticle 修改文章
func (article *ArticlesApi) UpdateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GetArticleList 查询文章
func (article *ArticlesApi) GetArticleList(c *gin.Context) {
	var articleOpts request.ArticleListRequest
	//校验必填信息
	if err := c.ShouldBindJSON(&articleOpts); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	list, total, msg, err := articleService.GetArticleListService(articleOpts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
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
	}
	list, msg, err := articleService.GetArticleHistoryService(uid)
	if err != nil {
		response.FailWithMessage(msg, c)
	}
	response.OkWithDetailed(gin.H{
		"list": list,
	}, "查询成功", c)
}
