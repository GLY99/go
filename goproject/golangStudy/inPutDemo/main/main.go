package main

import (
	"fmt"
)

func func1() {
	var name string
	var age int8
	var sal float32
	var isPass bool

	fmt.Printf("%s\n", "input name")
	flag, _ := fmt.Scanln(&name)
	if flag == 1 {
		fmt.Printf("name is %v\n", name)
	} else {
		fmt.Printf("scan fail %d", flag)
	}

	fmt.Printf("%s\n", "input age")
	flag, _ = fmt.Scanln(&age)
	if flag == 1 {
		fmt.Printf("age is %v\n", age)
	} else {
		fmt.Printf("scan fail %d", flag)
	}

	fmt.Printf("%s\n", "input sal")
	flag, _ = fmt.Scanln(&sal)
	if flag == 1 {
		fmt.Printf("sal is %v\n", sal)
	} else {
		fmt.Printf("scan fail %d", flag)
	}

	fmt.Printf("%s\n", "input isPass")
	flag, _ = fmt.Scanln(&isPass)
	if flag == 1 {
		fmt.Printf("isPass is %v\n", isPass)
	} else {
		fmt.Printf("scan fail %d", flag)
	}
}

func func2() {
	var name string
	var age int8
	var sal float32
	var isPass bool
	fmt.Printf("请输入name age sal isPass\n")
	flag, _ := fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	if flag == 4 {
		fmt.Printf("name %s, age %d, sal %f, isPass %t", name, age, sal, isPass)
	} else {
		fmt.Printf("scan fail %d", flag)
	}
}

func main() {
	func1()
	func2()
}
