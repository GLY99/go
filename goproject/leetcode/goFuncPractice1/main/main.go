package main

import (
	"fmt"
	"time"
)

func getFebruaryDays(year int) int {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 29
	} else {
		return 28
	}
}

func getMonthDays(year int, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		return getFebruaryDays(year)
	default:
		return -1
	}
}

func getYearDays(year int) int {
	var days int = 0
	for i := 1; i <= 12; i++ {
		monthDays := getMonthDays(year, i)
		days += monthDays
	}
	return days
}

func determineFish(times *time.Time) bool {
	// 假设从1999 01 01开始三天打鱼两天筛网，判断函数输入的时间是打鱼还是筛网
	year := times.Year()
	month := int(times.Month())
	day := times.Day()
	allDays := 0
	for i := 1999; i < year; i++ {
		yearDays := getYearDays(i)
		allDays += yearDays
	}
	for i := 1; i < month; i++ {
		monthDays := getMonthDays(year, i)
		allDays += monthDays
	}
	allDays += day
	if allDays%5 == 0 || allDays%5 == 4 {
		return false
	} else {
		return true
	}
}

func main() {
	times := time.Now()
	fmt.Println(determineFish(&times))
}
