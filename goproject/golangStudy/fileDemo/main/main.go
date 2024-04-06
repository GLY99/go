package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func closeFile(fd *os.File) {
	err := fd.Close()
	if err != nil {
		fmt.Printf("close file %v fail, err=%v\n", fd.Name(), err)
		return
	}
	fmt.Printf("close file %v succ\n", fd.Name())
}

func pathExist(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func copyFile(srcPath, dstPath string) (written int64, err error) {
	srcFd, err1 := os.Open(srcPath)
	if err1 != nil {
		return -1, err1
	}
	defer closeFile(srcFd)
	dstFd, err2 := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 6)
	if err2 != nil {
		return -1, err2
	}
	defer closeFile(dstFd)
	Reader := bufio.NewReader(srcFd)
	Writer := bufio.NewWriter(dstFd)
	written, err = io.Copy(Writer, Reader)
	return written, err
}

func main() {
	filePath := "C:/mycode/go/goproject/golangStudy/fileDemo/test.txt"
	fd, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file %s fail, err=%v\n", filePath, err)
		return
	}
	defer closeFile(fd)

	// 带缓冲的读取文件
	reader := bufio.NewReader(fd)
	err = nil
	str := ""
	for {
		if err == io.EOF { // 读取到文件末尾
			break
		} else if err != nil {
			fmt.Printf("读取文件失败, err=%v\n", err)
			break
		} else {
			str, err = reader.ReadString('\n')
			if !strings.HasSuffix(str, "\n") {
				str += "\n"
			}
			fmt.Printf("%s", str)
		}
	}

	// 一次性读取小文件
	byteSlice, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败, err=%v\n", err)
		return
	}
	fmt.Printf("%v\n", byteSlice)         // [104 101 108 108 111 32 119 111 114 108 100 13 10 104 101 108 108 111 32 103 111 232 175 173 232 168 128]
	fmt.Printf("%v\n", string(byteSlice)) // hello world\nhello go语言

	// 带缓冲的写入文件
	newFd, err := os.OpenFile(filePath, os.O_RDONLY|os.O_WRONLY, 6)
	if err != nil {
		fmt.Printf("open file %s fail, err=%v\n", filePath, err)
		return
	}
	defer closeFile(newFd)
	writer := bufio.NewWriter(newFd)
	for i := 0; i < 5; i++ {
		flag, err := writer.WriteString("hello 我爱学习\n") // flag接收的是写入的字节数
		if err != nil {
			fmt.Printf("write file fail\r\n")
			continue
		}
		fmt.Printf("本次写入了%d个字节到缓存\n", flag)
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush fail")
	} else {
		fmt.Println("flush succ")
	}

	// 一次性写入少量数据，这里会进行覆盖的写入
	str = "你在干神魔？"
	err = os.WriteFile(filePath, []byte(str), 6)
	if err != nil {
		fmt.Println("写入文件失败")
	}

	// 判断文件是否存在
	flag, err := pathExist("C:/")
	if flag && err == nil {
		fmt.Println("文件存在")
	} else if !flag && err == nil {
		fmt.Println("文件不存在")
	} else {
		fmt.Printf("无法判断文件是否存在, err=%v\n", err)
	}

	// copy文件
	dstPath := "C:/mycode/go/goproject/golangStudy/fileDemo/testcopy.txt"
	written, err := copyFile(filePath, dstPath)
	if err != nil {
		fmt.Printf("copy file fail, err=%v\n", err)
	} else {
		fmt.Printf("copy file succ, copy %d Byte\n", written)
	}
}
