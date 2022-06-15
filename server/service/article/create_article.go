package article

import (
	"context"
	"encoding/json"
	"fmt"
	"server/global"
	"server/model/article"
	"server/model/article/request"
	"time"
)

type ToArticleService struct{}

//CreateArticle 创建文章
func (a *ToArticleService) CreateArticle(articleOpts request.ArticleRequest) (at *article.Article, msg string, err error) {
	var articleContent article.Article
	articleContent.UUID = articleOpts.UUID
	articleContent.Title = articleOpts.Title
	articleContent.Status = articleOpts.Status
	articleContent.Slug = articleOpts.Slug
	articleContent.EditorType = articleOpts.EditorType
	articleContent.MetaKeyWords = articleOpts.MetaKeyWords
	articleContent.MetaDescription = articleOpts.MetaDescription
	articleContent.Summary = articleOpts.Summary
	articleContent.Thumbnail = articleOpts.Thumbnail
	articleContent.Visits = articleOpts.Visits
	articleContent.DisableComments = articleOpts.DisableComments
	articleContent.Password = articleOpts.Password
	articleContent.Likes = articleOpts.Likes
	articleContent.WordCount = articleOpts.WordCount
	articleContent.MdContent = articleOpts.MdContent
	articleContent.HtmlContent = articleOpts.HtmlContent
	articleContent.CommentCount = articleOpts.CommentCount
	articleContent.TagsID = articleOpts.TagsID
	articleContent.CategoryID = articleOpts.CategoryID
	articleContent.IsTop = articleOpts.IsTop
	articleContent.AuthID = articleOpts.AuthID
	//	存储数据库
	if err := global.DB.Create(&articleContent).Error; err != nil {
		return nil, "创建文章失败", err
	}
	// 存储redis
	msg, err = SetArticleRedis(articleContent)
	if err != nil {
		return nil, msg, err
	}
	return &articleContent, "创建文章成功", nil
}

func SetArticleRedis(articleContent article.Article) (msg string, err error) {
	ctx := context.Background()
	uuid := articleContent.UUID.String()
	now := time.Now()
	//	查出文章redis的数量
	at := make(map[time.Time]article.Article)
	//先查询是否存在
	if global.Redis.Exists(ctx, uuid).Val() == 1 {
		//	存在则更新
		fmt.Println("存在")
		//	获取redis的数据
		val, err := global.Redis.Get(ctx, uuid).Result()
		if err != nil {
			return "获取redis数据失败", err
		}
		//	转换成map
		err = json.Unmarshal([]byte(val), &at)
		if err != nil {
			return "老数据转换map失败", err
		}
	}
	at[now] = articleContent
	atJson, err := json.Marshal(at)
	_, err = global.Redis.Set(ctx, uuid, atJson, 0).Result()
	if err != nil {
		return "redis存储失败", err
	}
	return "redis存储成功", nil
}
