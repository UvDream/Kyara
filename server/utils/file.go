package utils

import (
	"os"
	"strings"
)

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

// SimplifyPath 简化路径
func SimplifyPath(path string) string {
	var stack []string
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}
