package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"strconv"
	"time"
)


//获取博客配置
func GetConfig() (err error, res response.SysConfigsResponse,msg string) {
	db := global.GVA_DB
	config := model.SysConfig{}
	err = db.Find(&config).Error
	if err!=nil{
		return err,res,"获取配置失败"
	}
	res.AuthorAvatar = config.AuthorAvatar
	//时间
	res.ActiveTime = config.ActiveTime
	res.BlogStartTime = config.BlogStartTime
	res.AuthorLink = config.AuthorLink
	res.AuthorMotto = config.AuthorMotto
	res.AuthorName = config.AuthorName
	res.BlogLogo = config.BlogLogo
	res.BlogName = config.BlogName
	//查询博客公告信息
	blogNotice:=model.BlogNotice{}
	err=db.Where("id=?",config.BlogNoticeID).Find(&blogNotice).Error
	//if err!=nil{
	//	return err,res,"获取博客公告失败"
	//}
	res.BlogNotice = blogNotice.Title
	res.BlogViewCount = config.BlogViewCount
	res.FilingMsg=config.FilingMsg
	//查询文章数量
	err=db.Table("sys_articles").Count(&res.ArticleCount).Error
	if err!=nil{
		return err,res,"获取文章数量失败"
	}
	return err, res,"获取配置成功"
}
//获取github仓库列表
func GetGithub()(githubList []response.GithubList,err error)  {
	blogConfig:= model.SysConfig{}
	db := global.GVA_DB
	err=db.Find(&blogConfig).Error
	r,err:=req.Get("https://api.github.com/users/"+blogConfig.GithubUserName+"/repos")
	if err==nil{
		r.ToJSON(&githubList)
	}
	return githubList,err
}
//访问博客
func ViewBlogCount(c *gin.Context)(err error,msg string)  {
	db := global.GVA_DB
	id := c.Query("id")
	if id!="blog"{
		article:=model.SysArticle{}
		err=db.Where("id=?",id).Find(&article).Error
		if err==nil{
			num, err:= strconv.ParseInt(article.ViewCount, 10, 64)
			num=num+1
			article.ViewCount=strconv.FormatInt(num, 10)
			err=db.Model(&article).Where("id=?",id).Update("view_count",article.ViewCount).Error
			if err!=nil{
				return err,"更新文章访问量失败"
			}
			return err,"更新文章访问量成功"
		}
	}
	if id=="blog"{
		//获取ip地址
		ipAddress:=utils.GetLocalIP()
		fmt.Println(ipAddress)
		//根据ip地址查询此ip是否存在库中
		var ipList model.BlogView
		err=db.Where("ip=?",ipAddress).First(&ipList).Error
		if err!=nil{
			//	不存在,就创建
			ipList.IP=ipAddress
			err=db.Create(&ipList).Error
			//更新
			err,msg=UpdateBlogView()
			if err!=nil{
				return err,msg
			}
		}else{
			//	存在
			//现在时间
			t2,_:=time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05") , time.Local)
			//之前更新时间
			t1,_:=time.ParseInLocation("2006-01-02 15:04:05", ipList.UpdatedAt.Format("2006-01-02 15:04:05") , time.Local)
			diff := (t2.Unix() - t1.Unix())/3600
			fmt.Println(diff)
			if diff>12{
				//需要更新时间
				err=db.Model(&ipList).Where("id = ?", ipList.ID).Update("updated_at", time.Now()).Error
				if err!=nil {
					return err,"更新新的访问时间失败"
				}
				err,msg=UpdateBlogView()
				if err!=nil{
					return err,msg
				}
			}
		}

		return err,"更新博客访问量成功"
	}
	return err,"更新访问量成功"
}
func UpdateBlogView() (err error,msg string) {
	db := global.GVA_DB
	config:=model.SysConfig{}
	err=db.Find(&config).Error
	num, err:= strconv.ParseInt(config.BlogViewCount, 10, 64)
	num=num+1
	config.BlogViewCount=strconv.FormatInt(num, 10)
	err=db.Model(&config).Update("blog_view_count",config.BlogViewCount).Error
	if err!=nil{
		return err,"更新博客访问量失败"
	}
	return err,"更新博客访问量成功"
}

//获取公告
func GetNotice(r request.ListStruct)(err error,list []model.BlogNotice,msg string,total int64)  {
	db := global.GVA_DB
	offset:=r.PageSize*(r.Page-1)
	err=db.Limit(r.PageSize).Offset(offset).Where("title LIKE ?", "%"+r.Search+"%").Find(&list).Error
	var config model.SysConfig
	err=db.Find(&config).Error
	if err!=nil{
		return err,list,"查询配置失败",total
	}
	for i,k:=range list{
		if k.ID==config.BlogNoticeID {
			list[i].Show="1"
		}else{
			list[i].Show="0"
		}
	}
	if err!=nil{
		return err,list,"获取公告失败",total
	}
	err = db.Where("title LIKE ?", "%"+r.Search+"%").Table("blog_notices").Count(&total).Error
	if err!=nil{
		return err,list,"获取公告总数失败",total
	}
	return  err,list,"获取公告成功",total
}
