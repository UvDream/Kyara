package utils

import "os"

// PathExists
//@description 检查文件夹是否存在
//@param path string
//@return bool
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
