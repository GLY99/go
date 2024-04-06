package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type collect struct {
	chCount    int
	numCount   int
	spaceCount int
	otherCount int
}

func collectChar(filePath string) (*collect, error) {
	charCount := &collect{}
	fd, err := os.OpenFile(filePath, os.O_RDONLY, 6)
	if err != nil {
		return charCount, err
	}
	defer func() {
		err := fd.Close()
		if err != nil {
			fmt.Printf("close file %s fail, err=%v\n", filePath, err)
		}
	}()
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if str == "" && err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			return charCount, err
		} else {
			for _, char := range str {
				switch {
				case char >= 'a' && char <= 'z':
					fallthrough
				case char >= 'A' && char <= 'Z':
					charCount.chCount++
				case char == ' ' || char == '\t':
					charCount.spaceCount++
				case char >= '0' && char <= '9':
					charCount.numCount++
				default:
					charCount.otherCount++
				}
			}
		}
	}
	return charCount, nil
}

func main() {
	filePath := "C:/mycode/go/goproject/leetcode/filePractice1/main/test.txt"
	charCount, err := collectChar(filePath)
	if err == nil {
		fmt.Printf("字母=%d,数字=%d,空格=%d,其它字符=%d", charCount.chCount, charCount.numCount, charCount.spaceCount, charCount.otherCount)
	} else {
		fmt.Printf("err=%v\n", err)
	}
}
