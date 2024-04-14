package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest1(t interface{}) {
	// rVal是指针
	rVal := reflect.ValueOf(t)
	fmt.Printf("rVal Kind=%v, Type=%v\n", rVal.Kind(), rVal.Type()) // rVal Kind=ptr, Type=*int

	// 通过reflect.Setxxx修改变量值
	// rVal.SetInt(20) // panic: reflect: reflect.Value.SetInt using unaddressable value

	// 地址类型需要获取到地址的值才能修改
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflectTest1(&num) // num=20
	fmt.Printf("num=%d\n", num)
	// var std Student = Student{Name: "Tom", Age: 20}
	// reflectTest1(&std) // 这里需要传递地址才能通过反射修改
}
