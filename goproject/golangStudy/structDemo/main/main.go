package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

type MyStruct struct {
	Num   int
	Ptr   *int
	Arr   [5]int
	Slice []int
	Map   map[string]string
}

func main() {
	var cat1 Cat
	cat1.Name = "miaomiao"
	cat1.Age = 17
	fmt.Println(cat1)

	// 结构体中属性都会有一个默认值, 对于引用类型定义完后默认值是nil, 需要make空间来使用
	var myStruct1 MyStruct
	fmt.Println(myStruct1.Num)   // 0
	fmt.Println(myStruct1.Ptr)   // <nil>
	fmt.Println(myStruct1.Arr)   // [0 0 0 0 0]
	fmt.Println(myStruct1.Slice) // []
	fmt.Println(myStruct1.Map)   // map[]
	if myStruct1.Slice == nil {
		fmt.Println("nil") // 这里输出
	}
	if myStruct1.Map == nil {
		fmt.Println("nil") // 这里输出
	}
	myStruct1.Slice = make([]int, 5, 10)
	myStruct1.Map = make(map[string]string, 5)
	num := 1
	myStruct1.Ptr = &num
	fmt.Println(myStruct1.Num)   // 0
	fmt.Println(myStruct1.Ptr)   // 0xc00000a118
	fmt.Println(myStruct1.Arr)   // [0 0 0 0 0]
	fmt.Println(myStruct1.Slice) // [0 0 0 0 0]
	fmt.Println(myStruct1.Map)   // map[]
	if myStruct1.Slice == nil {
		fmt.Println("nil") // 这里不会输出
	}
	if myStruct1.Map == nil {
		fmt.Println("nil") // 这里不会输出
	}

	// 不同结构体变量互相不会影响
	var cat2 Cat
	cat2.Name = "miaomiao"
	cat2.Age = 17
	var cat3 Cat
	cat3.Name = "mimi"
	cat3.Age = 18
	fmt.Println(cat2) // {miaomiao 17}
	fmt.Println(cat3) // {mimi 18}
}
