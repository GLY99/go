package main

import "fmt"

type T interface {
	Add()
}

type Num1 struct {
	n        int
	numSlice []int
}

type Num2 struct {
	n        int
	numSlice []int
}

func (num Num1) Add() {
	fmt.Printf("num1 add %v\n", num)
	num.n++
	num.numSlice = append(num.numSlice, 1)
}

func (num *Num2) Add() {
	fmt.Printf("num2 add %v\n", num)
	num.n++
	num.numSlice = append(num.numSlice, 1)
}

func test1(t T) { // 这里实际上其实就是Num1 Num2类型
	// num1.Add() &num1.Add() &num2.Add()
	// 因为num1的Add()是和Num类型绑定的，所以num1传地址或者值都可以, 实际过去的都是值
	// num2的Add()是和*Num2类型绑定的，所以num2只能传地址。
	fmt.Printf("test %v\n", t)
	t.Add()
}

func test2(t T) {
	// 接口时引用类型，里面存储的地址
	fmt.Printf("%p\n", t.Add) // 0x905d00
	test3(t)                  // 这里传递了一份地址
}

func test3(t T) {
	fmt.Printf("%p\n", t.Add) // 0x905d00
}

func main() {
	num1 := Num1{1, []int{}}
	num2 := Num2{1, []int{}}
	// Num1实现的接口方法都绑定在了结构体类型，所以传地址还是值都可以
	test1(num1)  // test {1 []} -> num1 add {1 []}
	test1(&num1) // test &{1 []} -> num1 add {1 []}
	// Num2实现的接口方法都绑定在了结构体指针类型，所以只能传地址
	test1(&num2)      // test &{1 []} -> num2 add &{1 []}
	fmt.Println(num1) // {1 []} 绑定到结构体类型，无论传值还是地址，实际取的都是值
	fmt.Println(num2) // {2 [1]}

	// 证明接口属于引用类型
	fmt.Printf("%p\n", num1.Add) // 0x905ca0
	test2(num1)                  // copy了一份num1中接口的方法
}
