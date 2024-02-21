package demos

import "fmt"

func Func1() {
	fmt.Printf("请输入小写字母a-z:\n")
	var c byte
	fmt.Scanf("%c\n", &c)
	switch c {
	case 'a':
		fmt.Printf("A\n")
	case 'b':
		fmt.Printf("B\n")
	case 'c':
		fmt.Printf("C\n")
	case 'd':
		fmt.Printf("D\n")
	default:
		fmt.Printf("other\n")
	}
}

func Func2() {
	fmt.Printf("请输入成绩:\n")
	var score float64
	fmt.Scanf("%f\n", &score)
	switch {
	case score >= 60 && score <= 100:
		fmt.Printf("及格\n")
	case score < 60:
		fmt.Printf("不及格\n")
	}
}

func Func3() {
	fmt.Printf("请输入月份:\n")
	var month int8
	fmt.Scanf("%d\n", &month)
	switch month {
	case 3, 4, 5:
		fmt.Printf("春季\n")
	case 6, 7, 8:
		fmt.Printf("夏季\n")
	case 9, 10, 11:
		fmt.Printf("秋季\n")
	case 12, 1, 2:
		fmt.Printf("冬季\n")
	default:
		fmt.Printf("错误的月份\n")
	}
}
