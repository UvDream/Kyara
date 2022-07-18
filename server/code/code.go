package code

import "server/global"

const (
	SUCCESS                 = 200
	ErrorImageNotFound      = 1001  //图片不存在
	UploadImageSuccess      = 2001  //上传图片成功
	ErrorCreateDir          = 1002  //创建文件夹失败
	SaveFileSuccess         = 2002  //保存文件成功
	ErrorSaveFile           = 1003  //保存文件失败
	ErrorSaveFileData       = 1004  //保存文件信息到数据库失败
	ErrorUploadQiNiu        = 1005  //七牛云上传失败
	ErrorOpenFile           = 1006  //打开文件失败
	ErrorUploadQiNiuSuccess = 3001  //七牛云上传成功
	ErrorMissingId          = 4001  //缺少id
	ErrorFileNotFound       = 404   //文件不存在
	ErrorDeleteFile         = 5001  //删除文件失败
	ErrorDeleteFileData     = 5002  //删除文件数据库失败
	ErrorListFile           = 6001  //获取文件列表失败
	ErrorCreateTheme        = 7001  //创建主题失败
	ErrorGetTheme           = 8001  //获取主题失败
	ErrorDeleteTheme        = 9001  //删除主题失败
	ErrorThemeNotFound      = 10001 //主题不存在
	ErrorUpdateTheme        = 10002 //主题更新失败

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
