package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str = "hello北"
	fmt.Printf("%d\n", len(str)) // 8, len按照字节计算长度，汉字占用三个字节

	// 字符串遍历存在中文需要转成[]rune
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
		if i == len(r)-1 {
			fmt.Printf("\n")
		}
	}

	// 字符串转整数，用strconv.Atoi(string) int, err
	str = "123"
	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("转换错误！\n")
	} else {
		fmt.Printf("%d\n", n)
	}

	// 整数转字符串
	str = strconv.Itoa(n)
	fmt.Printf("%s\n", str)

	// 字符串转byte切片
	str = "hello"
	b := []byte(str)
	fmt.Printf("%v\n", b) // [104 101 108 108 111]

	// byte切片转字符串
	b = []byte{104, 101, 108, 108, 111}
	str = string(b)
	fmt.Printf("%s\n", str)

	// 10进制转10 2 8 16进制的字符串表示
	str = strconv.FormatInt(123, 2)
	fmt.Printf("%s\n", str)

	// 查找子串是否在字符串中
	fmt.Printf("%t\n", strings.Contains("abcedef", "ab"))

	// 统计一个字符串中有几个不重叠子串
	fmt.Printf("%d\n", strings.Count("aaaaaaaa", "aa")) // 4

	// 不区分大小写进行字符串比较，==是区分大小写的
	fmt.Printf("%t\n", "abc" == "ABC")                  // false
	fmt.Printf("%t\n", strings.EqualFold("abc", "ABC")) // true

	// 返回子串在字符串中第一次出现的位置
	fmt.Printf("%d\n", strings.Index("abcedef", "abc"))

	// 返回子串在字符串中最后一次出现的位置
	fmt.Printf("%d\n", strings.LastIndex("abcdefabc", "abc"))

	// 字符串替换, -1表示全部替换
	fmt.Printf("%s\n", strings.Replace("abceabceda", "a", "A", -1))

	// 字符串按照指定字符分割
	fmt.Printf("%v\n", strings.Split("abce|abce|aaaa", "|")) // [abce abce aaaa]

	// 字符大小写转换
	fmt.Printf("%s\n", strings.ToLower("abcdE"))
	fmt.Printf("%s\n", strings.ToUpper("ABCDe"))

	// 将字符串左右两边的空格删除
	fmt.Printf("%s\n", strings.TrimSpace("  aaa  "))

	// 将字符串左右两边的指定字符删除
	fmt.Printf("%s\n", strings.Trim("abbbaa", "a"))

	// 将字符串左边或者右边的指定字符删除
	fmt.Printf("%s\n", strings.TrimLeft("abbbaa", "a"))
	fmt.Printf("%s\n", strings.TrimRight("abbbaa", "a"))

	// 判断字符串是否以某些字符开头或者结尾
	fmt.Printf("%t\n", strings.HasPrefix("aaccc", "aac"))
	fmt.Printf("%t\n", strings.HasSuffix("aaccc", "c"))
}
