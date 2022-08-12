package article

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm/clause"
	"server/global"
	"server/models/article"
	"time"
)

type ToArticleService struct{}

//CreateArticle 创建文章
func (*ToArticleService) CreateArticle(articleOpts article.Article) (at *article.Article, msg string, err error) {
	var articleContent article.Article
	articleContent = SetArticleContent(articleContent, articleOpts)
	//存储数据库
	if err := global.DB.Preload(clause.Associations).Create(&articleContent).Error; err != nil {
		return nil, "创建文章失败", err
	}
	//创建关联关系
	//category
	msg, err = CreateCategoryArticle(articleContent.ID, articleOpts.CategoriesID)
	if err != nil {
		return nil, msg, err
	}
	//tag
	msg, err = CreateTagArticle(articleContent.ID, articleOpts.TagsID)
	if err != nil {
		return nil, msg, err
	}
	//存储redis
	msg, err = SetArticleRedis(articleContent)
	if err != nil {
		return nil, msg, err
	}
	return &articleContent, "创建文章成功", nil
}

//CreateCategoryArticle 创建文章 category关联关系
func CreateCategoryArticle(articleUUID string, categoryID []uint) (msg string, err error) {
	db := global.DB
	for _, i := range categoryID {
		categoryArticle := article.CategoryArticle{}
		categoryArticle.ArticleID = articleUUID
		categoryArticle.CategoryID = i
		//	先查询是否存在
		if err := db.Where("article_id = ? and category_id = ?", articleUUID, i).First(&categoryArticle).Error; err == nil {
			return "category article关系已经存在", err
		}
		if err := db.Create(&categoryArticle).Error; err != nil {
			return "创建category关联关系失败", err
		}
	}
	return "创建文章分类关系成功", nil
}

//CreateTagArticle 创建文章 tag关联关系
func CreateTagArticle(articleID string, tags []uint) (msg string, err error) {
	db := global.DB
	for _, i := range tags {
		tagArticle := article.TagArticle{}
		tagArticle.ArticleID = articleID
		tagArticle.TagID = i
		//	先查询是否存在
		if err := db.Where("article_id = ? and tag_id = ?", articleID, i).First(&tagArticle).Error; err == nil {
			return "tag article关系已经存在", err
		}
		if err := db.Create(&tagArticle).Error; err != nil {
			return "创建关联关系失败", err
		}
	}
	return "tag article关系创建成功", nil
}

//SetArticleContent 设置文章内容
func SetArticleContent(articleContent article.Article, articleOpts article.Article) article.Article {
	articleContent.Title = articleOpts.Title
	articleContent.Status = articleOpts.Status
	articleContent.Slug = articleOpts.Slug
	articleContent.EditorType = articleOpts.EditorType
	articleContent.MetaKeyWords = articleOpts.MetaKeyWords
	articleContent.MetaDescription = articleOpts.MetaDescription
	articleContent.Summary = articleOpts.Summary
	articleContent.Thumbnail = articleOpts.Thumbnail
	articleContent.DisableComments = articleOpts.DisableComments
	articleContent.Password = articleOpts.Password
	articleContent.WordCount = articleOpts.WordCount
	articleContent.MdContent = articleOpts.MdContent
	articleContent.HtmlContent = articleOpts.HtmlContent
	articleContent.CommentCount = articleOpts.CommentCount
	articleContent.TagsID = articleOpts.TagsID
	articleContent.CategoriesID = articleOpts.CategoriesID
	articleContent.IsTop = articleOpts.IsTop
	articleContent.AuthID = articleOpts.AuthID
	return articleContent
}

// SetArticleRedis 存储redis
func SetArticleRedis(articleContent article.Article) (msg string, err error) {
	ctx := context.Background()
	id := articleContent.ID
	now := time.Now().Format("2006-01-02 15:04:05.0000")
	//	查出文章redis的数量
	at := make(map[string]article.Article)
	//先查询是否存在
	if global.Redis.Exists(ctx, id).Val() == 1 {
		//	存在则更新
		fmt.Println("存在")
		//	获取redis的数据
		val, err := global.Redis.Get(ctx, id).Result()
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
	_, err = global.Redis.Set(ctx, id, atJson, 0).Result()
	if err != nil {
		return "redis存储失败", err
	}
	return "redis存储成功", nil
}
