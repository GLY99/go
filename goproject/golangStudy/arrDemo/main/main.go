package main

import "fmt"

func test1(arr [3]int) {
	arr[0] = 0
}

func test2(arr *[3]int) {
	(*arr)[0] = 0
}

func main() {
	// 定义一个数组
	var arr [6]float64
	// 给数组每个元素赋值
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	arr[5] = 6

	// 数组的地址就是数组第一个元素的地址
	fmt.Printf("%p\n", &arr) //0xc000010420
	fmt.Println(&arr[0])     // 0xc000010420
	// 下一个元素的地址是前一个元素的地址加上它所占用的字节数
	fmt.Println(&arr[1]) // 0xc000010428

	// 数组的访问
	fmt.Println(arr[0]) // 1

	// 四种初始化数组的方式
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{4, 5, 6}
	var arr3 = [...]int{7, 8, 9}
	var arr4 = [3]string{1: "b", 0: "a", 2: "c"}
	fmt.Printf("%v, %v, %v, %v\n", arr1, arr2, arr3, arr4)

	// 数组遍历的两种方式
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for idx, val := range arr2 {
		fmt.Printf("%d, %d\n", idx, val)
	}

	// go数组是值传递
	arr5 := [...]int{1, 2, 3}
	test1(arr5)
	fmt.Printf("%T, %v\n", arr5, arr5) // [3]int, [1 2 3]
	test2(&arr5)
	fmt.Printf("%T, %v\n", arr5, arr5) // [3]int, [0 2 3]
}
