package main

import (
	"fmt"
	"math"
)

func func1() {
	var age int8
	fmt.Printf("请输入年龄\n")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Printf("你的年龄大于18\n")
	} else {
		fmt.Printf("你的年龄未满18岁\n")
	}
}

func func2() {
	var num1 int32 = 10
	var num2 int32 = 30
	if num1+num2 > 20 {
		fmt.Printf("hello world\n")
	}
}

func func3() {
	var num1 float32 = 12.5
	var num2 float32 = 13.5
	if num1 > 10 && num2 < 20 {
		fmt.Printf("hello world\n")
	}
}

func func4() {
	var num1 int32 = 10
	var num2 int32 = 30
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Printf("yes\n")
	} else {
		fmt.Printf("no\n")
	}
}

func func5() {
	var year int32 = 2014
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("run nian\n")
	} else {
		fmt.Printf("no run nian\n")
	}
}

func func6() {
	var score int8
	fmt.Printf("请输入成绩:\n")
	n, err := fmt.Scanln(&score)
	if err != nil {
		fmt.Printf("读取成绩错误\n")
		return
	}
	if n != 1 {
		fmt.Printf("只能输入一个成绩\n")
		return
	}
	if score == 100 {
		fmt.Printf("BMW\n")
	} else if score <= 99 && score > 80 {
		fmt.Printf("iphone7plus\n")
	} else if score <= 80 && score >= 60 {
		fmt.Printf("iPad\n")
	} else {
		fmt.Printf("挨打\n")
	}
}

func func7() {
	var a, b, c float64
	fmt.Printf("请依次输入abc的值,用‘,’分割\n")
	fmt.Scanf("%f,%f,%f\n", &a, &b, &c)
	var flag float64 = b*b - 4*a*c
	var x1 float64
	if flag > 0 {
		var x2 float64
		x1 = (-b + math.Sqrt(flag)) / 2 * a
		x2 = (-b - math.Sqrt(flag)) / 2 * a
		fmt.Printf("x1=%v, x2=%v\n", x1, x2)
	} else if flag == 0 {
		x1 = (-b + math.Sqrt(flag)) / 2 * a
		fmt.Printf("x1=%v\n", x1)
	} else {
		fmt.Printf("无解\n")
	}
}

func func8() {
	var height int16
	var memory float64
	var handsome bool
	fmt.Printf("请输入升高(cm),财富(千万),是否很帅(false/true)\n")
	fmt.Scanf("%d,%f,%t\n", &height, &memory, &handsome)
	if height > 180 && memory > 1 && handsome {
		fmt.Printf("ok1\n")
	} else if height > 180 || memory > 1 || handsome {
		fmt.Printf("ok2\n")
	} else {
		fmt.Printf("no\n")
	}
}

func func9() {
	var score float64
	fmt.Printf("请输入成绩:\n")
	fmt.Scanf("%f\n", &score)
	if score < 10 {
		fmt.Printf("恭喜你进入决赛,请输入性别:\n")
		var gender string
		fmt.Scanf("%s\n", &gender)
		if gender == "男" {
			fmt.Printf("请按时参加男子决赛\n")
		} else if gender == "女" {
			fmt.Printf("请按时参加女子决赛\n")
		} else {
			fmt.Printf("性别错误\n")
		}
	} else {
		fmt.Printf("很遗憾你没有进入决赛\n")
	}
}

func main() {
	func1()
	func2()
	func3()
	func4()
	func5()
	func6()
	func7()
	func8()
	func9()
}
