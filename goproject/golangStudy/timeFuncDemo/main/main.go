package main

import (
	"fmt"
	"strconv"
	"time"
)

func costTime() {
	str := ""
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
	endTime := time.Now()
	fmt.Printf("Start run at %v, end at %v, cost %ds\n",
		startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"),
		endTime.Unix()-startTime.Unix())
}

func main() {
	nowTime := time.Now()
	fmt.Printf("now time=%v\n", nowTime)

	fmt.Printf("%v\n", nowTime.Year())
	fmt.Printf("%v\n", nowTime.Month())      // March
	fmt.Printf("%v\n", int(nowTime.Month())) // 3
	fmt.Printf("%v\n", nowTime.Day())
	fmt.Printf("%v\n", nowTime.Hour())
	fmt.Printf("%v\n", nowTime.Minute())
	fmt.Printf("%v\n", nowTime.Second())

	// 格式化日期
	nowTimeStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d",
		nowTime.Year(), nowTime.Month(), nowTime.Day(),
		nowTime.Hour(), nowTime.Minute(), nowTime.Minute())
	fmt.Printf("%s\n", nowTimeStr)

	nowTimeStr = nowTime.Format("2006-01-02 15:04:05") // 2006-01-02 15:04:05 这里的每个数字是固定的
	fmt.Printf("%s\n", nowTimeStr)
	nowTimeStr = nowTime.Format("01/02 15:04:05") // 01-02 15:04:05 这里的每个数字是固定的
	fmt.Printf("%s\n", nowTimeStr)

	time.Sleep(5 * time.Second) // 休眠1s 这里默认是按照ns休眠，所以要带上定义好的时间常量

	n := nowTime.Unix() // 获取秒级别的世界戳
	fmt.Printf("%d\n", n)
	n = nowTime.UnixNano() // 获取纳秒级别的世界戳
	fmt.Printf("%d\n", n)

	costTime()
}
