package main

import "fmt"

func main() {
	// new用来给值类型分配内存，返回一个该类型的指针
	num1 := new(int)
	// num1 type=*int, num1 val=0xc00000a0a8, num1 aadr=0xc00006c020
	fmt.Printf("num1 type=%T, num1 val=%v, num1 aadr=%v", num1, num1, &num1)
}
