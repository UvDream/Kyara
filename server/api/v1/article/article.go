package article

import "github.com/gin-gonic/gin"

type ArticlesApi struct {
}

// CreateArticle 创建文章
func (article *ArticlesApi) CreateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// DeleteArticle 删除文章
func (article *ArticlesApi) DeleteArticle(c *gin.Context) {
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
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
