package blog

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/thedevsaddam/gojsonq/v2"
)

func Comment(r model.BlogComment) (err error, msg string) {
	db := global.GVA_DB
	r.Status = "0"
	if r.ID == 0 {
		err = db.Create(&r).Error
		if err != nil {
			return err, "留言入库失败"
		}
	}

	return err, "留言成功"
}
func GetComment(r request.ListStruct) (err error, msg string, blogComment []model.BlogComment, totalCount int64) {
	db := global.GVA_DB
	offset := r.PageSize * (r.Page - 1)

	err = db.Where("status=? AND parent_id=?", "1", "").Limit(r.PageSize).Offset(offset).Find(&blogComment).Count(&totalCount).Error
	if err != nil {
		return err, "查询失败", blogComment, 0
	}
	fmt.Println(blogComment)
	for k, i := range blogComment {
		blogComment[k].Children = findChildren(i.ID)
	}
	return err, "查询留言成功", blogComment, totalCount
}
func findChildren(parentId uint) (child []model.BlogComment) {
	db := global.GVA_DB
	err := db.Where("parent_id=? AND status=?", parentId, "1").Find(&child).Error
	fmt.Println(err)
	for k, i := range child {
		child[k].Children = findChildren(i.ID)
	}
	return child
}

//百度收录
func ToBaidu(r request.ToBaiDuRequest) (err error, msg string, baiduResponse response.BaiduResponse) {
	var config model.SysConfig
	db := global.GVA_DB
	err = db.Find(&config).Error
	if err != nil {
		return err, "获取配置失败", baiduResponse
	}
	url := "http://data.zz.baidu.com/urls?site=" + config.DomainName + "&token=" + config.BaiDuToken
	contentType := "text/plain"
	data := config.DomainName + r.Argument
	//判断是否收录
	err, msg = BaiduInclude(data)
	if err != nil {
		return err, "百度已经收录", baiduResponse
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, contentType, strings.NewReader(data))
	resp.Header.Add("User-Agent", "curl/7.12.1")
	resp.Header.Add("Host", "data.zz.baidu.com")
	resp.Header.Add("Content - Type", "text / plain")
	resp.Header.Add("Content-Length", "83")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	success := gojsonq.New().FromString(string(result)).Find("success")
	remain := gojsonq.New().FromString(string(result)).Find("remain")
	baiduResponse.Success = success
	baiduResponse.Remain = remain
	return err, "百度收录成功", baiduResponse
}

type baiduCollect struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//判断百度是否收录
func BaiduInclude(url string) (err error, msg string) {
	param := req.Param{
		"url": url,
	}
	r, err := req.Get("https://api.uomg.com/api/collect.baidu", param)
	if err != nil {
		return err, "请求百度是否收录失败"
	}
	var res baiduCollect
	r.ToJSON(&res)
	if res.Code == 1 {
		return errors.New("EOF"), "该网址已经收录"
	}
	return nil, "网址未被收录"
}

// GetClassifyStatistics 获取分类动态
func GetClassifyStatistics() (error, []int, []response.Classify) {
	var classify []model.SysClassify
	db := global.GVA_DB
	var countArr []int
	var classifyArr []response.Classify
	err := db.Find(&classify).Error
	if err != nil {
		return err, countArr, classifyArr
	}

	for _, k := range classify {
		var count int
		var classify response.Classify
		classify.Name = k.TypeName
		err := db.Where("classification_id = ?", k.ID).Table("sys_articles").Count(&count).Error
		if err != nil {
			return err, countArr, classifyArr
		}
		classify.Count = count
		classifyArr = append(classifyArr, classify)
		countArr = append(countArr, count)
	}
	max, _ := utils.GetMaxNumber(countArr)
	fmt.Println(max)
	for i, _ := range classifyArr {
		classifyArr[i].Max = max
	}
	return err, countArr, classifyArr
}

// AddDynamic 添加动态
func AddDynamic() {
	db := global.GVA_DB
	//今天
	now := time.Now().Format("2006-01-02")
	var dynamic model.BlogDynamic
	err := db.Where("date = ?", now).Find(&dynamic).Error
	fmt.Println(err)
	if err != nil {
		// 不存在
		dynamic.Date = now
		dynamic.Count = "1"
		err := db.Create(&dynamic).Error
		if err != nil {
			fmt.Println("创建失败")
		}
	} else {
		count, _ := strconv.Atoi(dynamic.Count)
		count++
		fmt.Println(count)
		err := db.Model(&dynamic).Where("ID = ?", dynamic.ID).Update(model.BlogDynamic{Count: strconv.Itoa(count)}).Error
		if err != nil {
			fmt.Println("更新失败")
		}
	}
}

// GetBlogDynamic 获取动态
func GetBlogDynamic() ([][]string, string, string, int, error) {
	db := global.GVA_DB
	//今天
	now := time.Now()
	//一年前的一天
	thatDay := now.AddDate(-1, 0, 0)
	var data [][]string
	var maxArr []int
	for t := thatDay; t.Before(now); t = t.AddDate(0, 0, 1) {
		var a []string

		var dynamic model.BlogDynamic
		err := db.Where("date = ?", t.Format("2006-01-02")).Find(&dynamic).Error
		if err != nil {
			a = append(a, t.Format("2006-01-02"))
			a = append(a, "0")
		} else {
			a = append(a, t.Format("2006-01-02"))
			a = append(a, dynamic.Count)
			i, _ := strconv.Atoi(dynamic.Count)
			maxArr = append(maxArr, i)
		}
		data = append(data, a)
	}
	max, _ := utils.GetMaxNumber(maxArr)
	AddDynamic()
	return data, now.Format("2006-01-02"), thatDay.Format("2006-01-02"), max, nil
}
