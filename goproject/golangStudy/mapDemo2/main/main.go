package main

import "fmt"

func test(myMap map[int]string) {
	fmt.Printf("%p\n", myMap) // 0xc0000200f0
	myMap[1] = "Z"
}

func test2(slice []int) {
	// 切片只有指向数组的指针是引用类型，size cap是值类型
	fmt.Printf("%p\n", slice) // 0xc00000a0e0
	slice = append(slice, 3)
	fmt.Println(slice) // [1 2 3]
}

func main() {
	var myMap = make(map[int]string)
	myMap[1] = "A"
	fmt.Printf("%p\n", myMap) // 0xc0000200f0
	test(myMap)
	fmt.Println(myMap) // map[1:Z]

	slice := []int{1, 2}
	fmt.Printf("%p\n", slice) // 0xc00000a0e0
	test2(slice)
	fmt.Println(slice) // [1 2]
}
