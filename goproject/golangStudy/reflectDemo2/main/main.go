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
	// 获取t变量的值,返回reflect.Value
	rVal := reflect.ValueOf(t)

	// 通过reflect.Value获取变量类别
	fmt.Printf("变量类别是%v, 类型是%v\n", rVal.Kind(), rVal.Type()) // 变量类别是struct, 类型是main.Student
}

func main() {
	var std Student = Student{Name: "Tom", Age: 20}
	reflectTest1(std)
}
