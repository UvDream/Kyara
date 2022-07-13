package article

import (
	"server/global"
	"server/model/article"
	"server/model/article/request"
)

func (a *ToArticleService) UpdateArticleService(articleOpts request.ArticleRequest) (ac *article.Article, msg string, err error) {
	//修改文章
	var articleContent article.Article
	articleContent = SetArticleContent(articleContent, articleOpts)
	//	存储数据库
	if err := global.DB.Model(article.Article{}).Where("uuid = ?", articleContent.UUID).Updates(&articleContent).Error; err != nil {
		return nil, "修改文章失败", err
	}
	//存储redis
	msg, err = SetArticleRedis(articleContent)
	if err != nil {
		return nil, "修改文章存储redis失败", err
	}
	return &articleContent, "修改文章成功", nil
}
