package file

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"server/code"
	"server/global"
	"server/utils"
	"strings"
	"time"
)

type LocalService struct{}

func (s *LocalService) UploadFile(file *multipart.FileHeader) (string, string, int, error) {
	//	读取文件后缀
	ext := path.Ext(file.Filename)
	//拼接文件名
	name := strings.TrimSuffix(file.Filename, ext)
	fileName := name + "_" + time.Now().Format("20060102150405") + ext
	filePosition := global.Config.Local.Path
	filePath := filePosition + "/" + fileName
	//文件路径
	err := utils.CreateDir(filePosition)
	if err != nil {
		return "", "", code.ErrorCreateDir, err
	}
	f, err := file.Open()
	if err != nil {
		return "", "", code.ErrorOpenFile, err
	}
	defer f.Close()
	out, err := os.Create(filePath)
	if err != nil {
		return "", "", code.ErrorSaveFile, err
	}
	defer out.Close()
	_, err = io.Copy(out, f) //拷贝文件
	if err != nil {
		return "", "", code.ErrorSaveFile, err
	}
	return utils.SimplifyPath(filePath), fileName, code.SaveFileSuccess, nil
}

// DeleteFile 删除文件
func (s *LocalService) DeleteFile(key string) error {
	p := global.Config.Local.Path + "/" + key
	if strings.Contains(p, global.Config.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
