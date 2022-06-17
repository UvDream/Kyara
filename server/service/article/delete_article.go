package article

import (
	"context"
	"server/global"
	"server/model/article"
)

func (a *ToArticleService) DeleteArticleService(id string) (msg string, err error) {
	db := global.DB
	//查询文章是否存在
	err = db.Where("uuid = ?", id).First(&article.Article{}).Error
	if err != nil {
		return "文章不存在", err
	}
	//删除文章
	err = db.Where("uuid = ?", id).Delete(&article.Article{}).Error
	if err != nil {
		return "删除文章失败", err
	}
	//删除历史文章
	ctx := context.Background()
	if global.Redis.Exists(ctx, id).Val() == 1 {
		err = global.Redis.Del(ctx, id).Err()
		if err != nil {
			return "删除历史文章失败", err
		}
	}
	return "删除文章成功", err
}
