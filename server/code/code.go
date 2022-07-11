package code

import "server/global"

const (
	SUCCESS            = 200
	ErrorImageNotFound = 1001 //图片不存在
	UploadImageSuccess = 2001 //上传图片成功
)

func Text(code int) string {
	lang := global.Config.System.Language
	if lang == "en-us" {
		return enUSText[code]
	}
	if lang == "zh-cn" {
		return zhCNText[code]
	}
	return zhCNText[code]
}
