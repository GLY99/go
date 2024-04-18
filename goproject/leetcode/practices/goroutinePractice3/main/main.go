package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func writeFile(filePath string, writeFinash chan bool) {
	defer func() {
		writeFinash <- true
	}()
	fd, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 6)
	if err != nil {
		fmt.Printf("open file %s fail\n", filePath)
		return
	}
	numStr := ""
	for i := 1; i <= 1000; i++ {
		numStr += strconv.Itoa(rand.Intn(10000))
		if i != 1000 {
			numStr += ","
		}
	}
	writer := bufio.NewWriter(fd)
	_, err = writer.WriteString(numStr)
	if err != nil {
		fmt.Printf("write data to file %s fail\n", filePath)
	} else {
		fmt.Printf("write data to file %s succ\n", filePath)
	}
}

func sortData(readFilePath, writeFilePath string, sortFinash chan bool) {
	defer func() {
		sortFinash <- true
	}()
	rfd, err1 := os.OpenFile(readFilePath, os.O_RDONLY, 6)
	if err1 != nil {
		fmt.Printf("open file %s fail\n", readFilePath)
		return
	}
	numStr := ""
	reader := bufio.NewReader(rfd)
	for {
		v, err := reader.ReadString('\n')
		if v == "" && err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			fmt.Printf("reader data fail from file %s, err=%v\n", readFilePath, err)
			return
		} else {
			numStr += v
		}
	}
	numStrSlice := strings.Split(numStr, ",")
	numSlice := make([]int, 1000)
	for idx, v := range numStrSlice {
		num, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if err != nil {
			fmt.Printf("ParseInt fail, val is %s, err=%v\n", v, err)
			return
		}
		numSlice[idx] = int(num)
	}
	sort.Ints(numSlice)

	wfd, err := os.OpenFile(writeFilePath, os.O_WRONLY|os.O_TRUNC, 6)
	if err != nil {
		fmt.Printf("open file %s fail\n", writeFilePath)
		return
	}
	numStr = ""
	for idx, num := range numSlice {
		numStr += strconv.Itoa(num)
		if idx != 1000 {
			numStr += ","
		}
	}
	writer := bufio.NewWriter(wfd)
	_, err = writer.WriteString(numStr)
	if err != nil {
		fmt.Printf("write data to file %s fail\n", writeFilePath)
		return
	}
}

func waitFinash(flagChannel chan bool) {
	count := 0
	for {
		if count == 5 {
			break
		}
		v := <-flagChannel
		if v {
			count++
		}
	}
}

func main() {
	fileBasePath, err := os.Getwd() // C:\mycode\go\goproject\leetcode
	if err != nil {
		os.Exit(1)
	}
	flagChannel := make(chan bool, 5)
	defer close(flagChannel)
	for i := 1; i <= 5; i++ {
		filepath := fileBasePath + "/goroutinePractice3/test" + strconv.Itoa(i) + ".txt"
		go writeFile(filepath, flagChannel)
	}

	waitFinash(flagChannel)

	for i := 1; i <= 5; i++ {
		readfilepath := fileBasePath + "/goroutinePractice3/test" + strconv.Itoa(i) + ".txt"
		writefilepath := fileBasePath + "/goroutinePractice3/result" + strconv.Itoa(i) + ".txt"
		go sortData(readfilepath, writefilepath, flagChannel)
	}

	waitFinash(flagChannel)
}
