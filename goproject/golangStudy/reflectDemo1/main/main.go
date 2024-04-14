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
	// 获取t变量的实际类型,返回reflect.Type类型
	rType := reflect.TypeOf(t)
	fmt.Printf("rType:%v\n", rType) // rType:int 这个rType不是基本数据中的int类型，是reflect包中的Type类型

	// 获取t变量的值,返回reflect.Value
	rVal := reflect.ValueOf(t)
	fmt.Printf("rVal:%v\n", rVal) // rVal:100 这个rVal不是真的100，是reflect包中的Value类型

	// reflect.Value类型不能直接参与计算，因为他是结构体类型
	num1 := 2 + rVal.Int()
	fmt.Printf("num1=%d\n", num1) // num1=102

	// 将refelct.Value类型变量转成interface{}类型的变量
	i := rVal.Interface()

	// 通过断言将interface类型转成需要的基本数据类型中的int
	num2 := i.(int)
	fmt.Printf("num2=%d, num2类型是%T\n", num2, num2) // num2=100, num2类型是int
}

func reflectTest2(t interface{}) {
	// 获取t变量的实际类型,返回reflect.Type类型
	rType := reflect.TypeOf(t)
	// 这个rType不是Student类型，是reflect包中的Type类型
	fmt.Printf("rType:%v\n", rType) // rType:main.Student

	// 获取t变量的值,返回reflect.Value
	rVal := reflect.ValueOf(t)
	// 这个rVal不是真的100，是reflect包中的Value类型
	fmt.Printf("rVal:%v, rVal类型是%T\n", rVal, rVal) // rVal:{Tom 20}, rVal类型是reflect.Value

	// 将refelct.Value类型变量转成interface{}类型的变量
	i := rVal.Interface()
	// i类型是main.Student 这里的接口类型虽然是Student但是不能对Student的属性操作，还是需要类型断言
	fmt.Printf("i:%v, i类型是%T\n", i, i) // i:{Tom 20}, i类型是main.Student

	// 通过断言将interface类型转成需要的自定义类型Student, 两种断言都可以
	// student1, ok := i.(Student)
	// if ok {
	// 	fmt.Printf("student1:%v, student1类型是%T\n", student1, student1)               // student1:{Tom 20}, student1类型是main.Student
	// 	fmt.Printf("student1.Name=%s, student1.Age=%d", student1.Name, student1.Age) // student1.Name=Tom, student1.Age=20
	// 	student1.Age = 30
	// }

	switch v := i.(type) {
	case Student:
		fmt.Printf("student1:%v, student1类型是%T\n", v, v)                 // student1:{Tom 20}, student1类型是main.Student
		fmt.Printf("student1.Name=%s, student1.Age=%d\n", v.Name, v.Age) // student1.Name=Tom, student1.Age=20
		v.Age = 30
	case *Student:
		fmt.Printf("student1:%v, student1类型是%T\n", v, v)                 // student1:{Tom 20}, student1类型是main.Student
		fmt.Printf("student1.Name=%s, student1.Age=%d\n", v.Name, v.Age) // student1.Name=Tom, student1.Age=20
		v.Age = 30
	default:
		fmt.Printf("没有找到的类型\n")
	}
	fmt.Printf("after type assert, %v\n", i) // 如果t是指针变量，这里输出&{Tom 30}，如果否则输出{Tom, 20}
}

func main() {

	// 对基本数据类型进行反射基本操作
	var num int = 100
	reflectTest1(num)

	// 对结构体类型进行反射基本操作
	var std Student = Student{Name: "Tom", Age: 20}
	reflectTest2(std)
	fmt.Println(std) // {Tom 20}
	reflectTest2(&std)
	fmt.Println(std) // {Tom 30}
}
