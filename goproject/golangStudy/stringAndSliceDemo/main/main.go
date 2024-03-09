package main

import "fmt"

func main() {
	// string底层是一个byte数组，因此可以切片
	str := "helloworld"
	slice1 := str[5:]
	fmt.Println(slice1) // world

	// string本身不可变，如果通过str[0]修改会报错，但是可以通过切片修改
	slice2 := []byte(str)
	slice2[5] = 'W'
	str = string(slice2)
	fmt.Println(slice2, str) // [104 101 108 108 111 87 111 114 108 100] helloWorld

	// []byte()是按照一个字节处理，但是一个汉字三个字节，处理会乱码, 但是可以用[]rune
	slice3 := []rune(str)
	slice3[4] = '欧'
	str = string(slice3)
	fmt.Println(slice3, str) // [104 101 108 108 27431 87 111 114 108 100] hell欧World
}
