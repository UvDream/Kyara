package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/article"
	"server/model/article/request"
	"server/utils"
)

func (a *ToArticleService) GetArticleListService(query request.ArticleListRequest, uuid string, c *gin.Context) (list *[]article.Article, total int64, msg string, err error) {

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
		db = db.Where("create_time >= ?", query.StartTime)
	}
	if query.EndTime != "" {
		db = db.Where("create_time <= ?", query.EndTime)
	}
	//TODO:查询用户文章,区分出管理员和普通用户
	//查询总数
	if err = db.Where("auth_id = ?", uuid).Count(&total).Error; err != nil {
		return nil, 0, "查询总数失败", err
	}
	//查询列表
	if err = db.Where("auth_id = ?", uuid).Preload("Tags").Preload("Categories").Preload("Auth").Order("create_time desc").Scopes(utils.Paginator(c)).Find(&articleList).Error; err != nil {
		return nil, 0, "查询列表失败", err
	}
	return &articleList, total, "查询成功", err
}

//GetArticleHistoryService 获取文章历史
func (a *ToArticleService) GetArticleHistoryService(id string) (list interface{}, msg string, err error) {
	ctx := context.Background()
	msg, err = global.Redis.Get(ctx, id).Result()
	if err != nil {
		return nil, "查询redis失败", err
	}
	at := make(map[string]article.Article)
	err = json.Unmarshal([]byte(msg), &at)
	if err != nil {
		return nil, "数据转换失败", err
	}
	return at, "查询成功", nil
}

//GetArticleDetailService 获取文章详情
func (a *ToArticleService) GetArticleDetailService(id string) (list article.Article, msg string, err error) {
	err = global.DB.Where("id=?", id).Preload("Tags").Preload("Categories").Preload("Auth").First(&list).Error
	if err != nil {
		return article.Article{}, "查询失败", err
	}
	return list, "查询成功", nil
}
