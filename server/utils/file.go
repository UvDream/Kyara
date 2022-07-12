package utils

import "os"

// IsDirExist 判断文件夹是否存在
func IsDirExist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

// CreateDir 创建文件夹
func CreateDir(dir string) (err error) {
	if !IsDirExist(dir) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
