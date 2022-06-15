package article

import (
	"context"
	"server/global"
	"server/model/article"
)

func (a *ToArticleService) DeleteArticleService(id string) (msg string, err error) {
	//删除文章
	err = global.DB.Where("uuid = ?", id).Delete(&article.Article{}).Error
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
