package main

import (
	"fmt"
	"reflect"
)

// 使用反射创建并且操作结构体
type User struct {
	Name string
	Age  int
}

func NewStruct() *User {
	var model *User
	// 获取model遍历的类型*User
	tType := reflect.TypeOf(model)
	fmt.Printf("tType Kind is %v\n", tType.Kind()) // tType Kind is ptr
	// 获得*User实际指向的User类型
	tElem := tType.Elem()
	fmt.Printf("tElem Kind is %v\n", tElem.Kind()) // tElem Kind is struct
	// New返回一个Value类型值，该值有一个指向类型为tElem的新申请的零值指针
	elem := reflect.New(tElem)
	fmt.Printf("elem Kind is %v\n", elem.Kind())               // elem Kind is ptr
	fmt.Printf("elem.Elem() Kind is %v\n", elem.Elem().Kind()) // elem.Elem() Kind is struct
	// 将reflect.Value类型转成指针，然后转成*User
	model = elem.Interface().(*User) // model是*User类型，它和elem指向的内存空间是一样的
	elem = elem.Elem()
	model.Name = "xiaohua"
	elem.FieldByName("Age").SetInt(18)
	return model
}

func main() {
	fmt.Println(NewStruct()) // &{xiaohua 18}
}
