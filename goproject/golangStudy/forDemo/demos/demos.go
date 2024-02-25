package demos

import (
	"fmt"
	"math/rand"
	_ "time"
)

func Func1() {
	// for 循环基本语法, 这里i的作用域只在for循环中
	for i := 0; i < 10; i++ {
		fmt.Printf("hellow world\n")
	}

	// for循环的写法2, 这里j的作用域在这个函数中
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("hellow world\n")
	}
	fmt.Printf("%d\n", j)

	// for 循环写法3，类似while
	var n int64 = 0
	for n < 10 {
		fmt.Printf("hellow world\n")
		n++
	}

	// 无限for循环
	var m int64 = 0
	for {
		if m < 10 {
			fmt.Printf("hellow world\n")
			m++
		} else {
			break
		}
	}
}

func Func2() {
	// 遍历字符串
	var str string = "hello world!北京"
	// 如果字符串中有中文，采用下面的方式遍历会乱码
	// 下面的这种方式一次取一个字节，但是go采用utf8编码，中文占用三个字节
	// 解决办法：1.用切片，2.用range
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
		if i == len(str)-1 {
			fmt.Printf("\n")
		}
	}

	for idx, val := range str {
		fmt.Printf("%d->%c\n", idx, val)
	}

	var str2 []rune = []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
		if i == len(str2)-1 {
			fmt.Printf("\n")
		}
	}

}

func Func3() {
	var sum int32 = 0
	var count int8 = 0
	var i int32 = 1
	for ; i <= 100; i++ {
		if i%9 == 0 {
			sum = sum + i
			count++
		}
	}
	fmt.Printf("count=%d, sum=%d\n", count, sum)
}

func Func4() {
	limit := 6
	for i := 0; i <= limit; i++ {
		fmt.Printf("%d + %d = %d\n", i, limit-i, limit)
	}
}

func Func5() {
	// go语言的while do while
	var i int8 = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Printf("hello world\n")
		i++
	}

	i = 0
	for {
		fmt.Printf("hello world\n")
		i++
		if i >= 10 {
			break
		}
	}

}

func Func6() {
	// 空心金字塔
	// f = 3
	//   *    i=1; all -> f + i - 1; " " f-i
	//  ***   i=2; all -> f + i - 1; " " f-i
	// *****  i=3; all -> f + i - 1; " " f-i
	f := 6
	for i := 1; i <= f; i++ {
		for j := 1; j <= f+i-1; j++ {
			if j <= f-i {
				fmt.Printf(" ")
			} else {
				if j == f-i+1 || j == f+i-1 || i == f {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}

	// 用while
	i := 1
	j := 1
	for {
		if i > f {
			break
		}
		j = 1
		for {
			if j > f+i-1 {
				break
			}
			if j <= f-i {
				fmt.Printf(" ")
			} else {
				if j == f-i+1 || j == f+i-1 || i == f {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			j++
		}
		fmt.Printf("\n")
		i++
	}
}

func Func7() {
	// 99乘法表
	// 1 * 1 = 1
	// 1 * 2 = 2; 2 * 2 = 4
	// 1 * 3 = 3; 2 * 3 = 6; 3 * 3 = 9
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
		}
		fmt.Printf("\n")
	}
}

func Func8() {
	var classNums int8
	fmt.Printf("请输入班级数量：\n")
	fmt.Scanf("%d\n", &classNums)
	for i := 1; i <= int(classNums); i++ {
		var sumScores float64 = 0.0
		var studentNums int8
		fmt.Printf("请输入第%d个班级的人数：\n", i)
		fmt.Scanf("%d\n", &studentNums)
		for j := 1; j <= int(studentNums); j++ {
			var score float64
			fmt.Printf("请输入第%d个班级第%d个学生的总成绩：\n", i, j)
			fmt.Scanf("%f\n", &score)
			sumScores = sumScores + score
		}
		fmt.Printf("第%d班共%d个学生，平均成绩%f\n", i, studentNums, sumScores/float64(studentNums))
	}
}

func Func9() {
	var count uint = 0
	for {
		num := rand.Intn(11)
		count = count + 1
		fmt.Printf("%d\n", num)
		if num == 9 {
			break
		}
	}
	fmt.Printf("一共使用了%d次\n", count)
}

func Func10() {
lable1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				break lable1
			}
			fmt.Printf("%d\n", j)
		}
	}
}

func Func11() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum = sum + i
		if sum == 3 {
			fmt.Printf("%d\n", i)
			break
		}
	}
}

func Func12() {
	var password string
	var count uint8 = 3
	for i := count; i > 0; i-- {
		fmt.Printf("请输入密码(还有%d次机会)：\n", i)
		fmt.Scanf("%s\n", &password)
		if password == "888" {
			fmt.Printf("密码正确")
			break
		}
	}
}

func Func13() {
lable1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				// continue
				continue lable1 // continue外层循环相当于break本次循环
			}
			fmt.Printf("%d\n", j)
		}
	}
	// 01010101
}

func Func14() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d\n", i)
	}
}

func Func15() {
	positive_nums := 0
	negative_nums := 0
	for {
		fmt.Printf("请输入正数或者负数\n")
		var num float64
		fmt.Scanf("%f\n", &num)
		if num == 0 {
			break
		}
		if num < 0 {
			negative_nums++
		} else {
			positive_nums++
		}
	}
	fmt.Printf("正数%d个，负数%d个", positive_nums, negative_nums)
}

func Func16() {
	var count uint64 = 0
	var money float64 = 100000
	for {
		if money < 1000 {
			break
		}
		if money > 50000 {
			money = money - money*0.05
			count++
			continue
		}
		if money <= 50000 {
			money = money - 1000
			count++
		}
	}
	fmt.Printf("%d", count)
}

func Func17() {
	for i := 0; i <= 2; i++ {
		if i == 1 {
			goto lable1 //跳到lable1所在的行继续执行，后面的循环都不会执行了。如果这里lable1放在上面会死循环
		}
	}

lable1:
	for i := 0; i < 5; i++ {
		fmt.Printf("hello world\n")
	}
}
