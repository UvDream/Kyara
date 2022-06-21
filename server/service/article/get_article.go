package article

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"server/global"
	"server/model/article"
	"server/model/article/request"
	"time"
)

func (a *ToArticleService) GetArticleListService(query request.ArticleListRequest) (list *[]article.Article, total int64, msg string, err error) {
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	var articleList []article.Article
	db := global.DB.Model(article.Article{})
	//keyword
	if query.KeyWord != "" {
		db = db.Where("title LIKE ?", "%"+query.KeyWord+"%").Or("md_content LIKE ?", "%"+query.KeyWord+"%").Or("html_content LIKE ?", "%"+query.KeyWord+"%")
	}
	//status
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.StartTime != "" {
		db = db.Where("created_at >= ?", query.StartTime)
	}
	if query.EndTime != "" {
		db = db.Where("created_at <= ?", query.EndTime)
	}
	//查询总数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, "查询总数失败", err
	}
	//查询列表
	if err = db.Preload("Tags").Limit(limit).Offset(offset).Find(&articleList).Error; err != nil {
		return nil, 0, "查询列表失败", err
	}
	//查询标签
	//for k, i := range articleList {
	//	articleList[k].Tags, articleList[k].TagsID, msg, err = getTags(i.UUID)
	//	if err != nil {
	//		return nil, 0, "查询标签失败", err
	//	}
	//}
	return &articleList, total, "查询成功", err
}

//GetArticleHistoryService 获取文章历史
func (a *ToArticleService) GetArticleHistoryService(id string) (list interface{}, msg string, err error) {
	ctx := context.Background()
	msg, err = global.Redis.Get(ctx, id).Result()
	if err != nil {
		return nil, "查询失败", err
	}
	at := make(map[time.Time]article.Article)
	err = json.Unmarshal([]byte(msg), &at)
	if err != nil {
		return nil, "数据转换失败", err
	}
	return at, "查询成功", nil
}

func getTags(articleID uuid.UUID) (tags []article.Tag, tagsID []uint, msg string, err error) {
	db := global.DB
	var tagArticle []article.TagArticle
	if err := db.Where("article_id = ?", articleID).Find(&tagArticle).Error; err != nil {
		return nil, nil, "查询标签失败", err
	}
	for _, i := range tagArticle {
		tagsID = append(tagsID, i.TagID)
		//	根据tag_id查询tag
		if err := db.Where("id = ?", i.TagID).Find(&tags).Error; err != nil {
			return nil, nil, "查询标签内容失败", err
		}
	}
	return tags, tagsID, "查询成功", nil
}
