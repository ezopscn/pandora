package utils

import (
	"os"
)

// 判断文件是否存在
func FileExist(path string) bool {
	stat, err := os.Stat(path)
	if os.IsExist(err) && !stat.IsDir() {
		return true
	}
	return false
}

// 判断目录是否存在
func DirExist(path string) bool {
	stat, err := os.Stat(path)
	if os.IsExist(err) && stat.IsDir() {
		return true
	}
	return false
}
