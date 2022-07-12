package code

import "server/global"

const (
	SUCCESS            = 200
	ErrorImageNotFound = 1001 //图片不存在
	UploadImageSuccess = 2001 //上传图片成功
	ErrorCreateDir     = 1002 //创建文件夹失败
	SaveFileSuccess    = 2002 //保存文件成功
	ErrorSaveFile      = 1003 //保存文件失败
	ErrorSaveFileData  = 1004 //保存文件信息到数据库失败
	ErrorUploadQiNiu   = 1005 //七牛云上传失败
	ErrorOpenFile      = 1006 //打开文件失败
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
