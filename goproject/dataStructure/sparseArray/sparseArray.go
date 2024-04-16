package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	row, col, val int
}

const arrRow = 20
const arrCol = 20

func main() {
	// 创建原始数组
	var chessArr [arrRow][arrCol]int
	chessArr[1][9] = 1
	chessArr[10][10] = 2

	// 输出原始数组
	for _, v := range chessArr {
		fmt.Println(v)
	}

	// 原始数组转成稀疏数组
	var sparseArr []node
	sparseArr = append(sparseArr, node{arrRow, arrCol, 0})
	for i, v := range chessArr {
		for j, v1 := range v {
			if v1 != 0 {
				valNode := node{i, j, v1}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println(sparseArr)

	// 稀疏数组存盘
	modelPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd() fail, err=%v\n", err)
		return
	}
	dataPath := modelPath + "/sparseArray/sparseArray.data"
	fd, err := os.OpenFile(dataPath, os.O_RDWR, 6)
	if err != nil {
		fmt.Printf("open file %s fail, err=%v\n", dataPath, err)
		return
	}
	defer fd.Close()
	writer := bufio.NewWriter(fd)
	for _, v := range sparseArr {
		str := fmt.Sprintf("%d,%d,%d", v.row, v.col, v.val)
		_, err := writer.WriteString(str + "\n")
		if err != nil {
			fmt.Printf("write data %s fail, err=%v\n", str, err)
			continue
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Printf("writer.Flush() fail, err=%v\n", err)
		return
	}

	// 稀疏数组转原始数组
	var newChessArr [arrRow][arrCol]int
	readfd, err := os.OpenFile(dataPath, os.O_RDWR, 6)
	if err != nil {
		fmt.Printf("open file %s fail, err=%v\n", dataPath, err)
		return
	}
	defer readfd.Close()
	reader := bufio.NewReader(readfd)
	for {
		str, err := reader.ReadString('\n')
		if str == "" && err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			fmt.Printf("read data fail, err=%v\n", err)
			continue
		} else {
			str := strings.Trim(str, "\r\n")
			strSlice := strings.Split(str, ",")
			row, err1 := strconv.ParseInt(strSlice[0], 10, 64)
			col, err2 := strconv.ParseInt(strSlice[1], 10, 64)
			val, err3 := strconv.ParseInt(strSlice[2], 10, 64)
			if err1 != nil || err2 != nil || err3 != nil {
				fmt.Printf("ParseInt fail with %v, %v, %v\n", strSlice[0], strSlice[1], strSlice[2])
			}
			if row == arrRow || col == arrCol {
				continue
			}
			newChessArr[row][col] = int(val)
		}
	}
	for _, v := range newChessArr {
		fmt.Println(v)
	}
}
