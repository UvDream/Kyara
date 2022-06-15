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
	if err := global.DB.Where("uuid = ?", articleOpts.UUID).Save(&articleContent).Error; err != nil {
		return nil, "修改文章失败", err
	}
	//存储redis
	msg, err = SetArticleRedis(articleContent)
	if err != nil {
		return nil, msg, err
	}
	return nil, "", nil
}
