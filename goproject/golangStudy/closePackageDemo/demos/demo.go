package demos

import "strings"

func MakeSuffix(suffix string) func(string) string {
	// 通过闭包实现，判断一个文件名是否以指定文件后缀结尾
	return func(fileName string) string {
		if !strings.HasSuffix(fileName, suffix) {
			return fileName + suffix
		}
		return fileName
	}
}
