package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	num1, num2 int
}

func (cal *Cal) GetSum(name string) {
	fmt.Printf("%s 完成了加法 %d + %d = %d", name, cal.num1, cal.num2, cal.num1+cal.num2)
}

func ReflectFunc(t interface{}) {
	// rType := reflect.TypeOf(t)
	rVal := reflect.ValueOf(t)
	fieldNum := rVal.Elem().NumField()
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("idx is %d field val is %v\n", i, rVal.Elem().Field(i))
		// idx is 0 field val is 1
		// idx is 1 field val is 1
	}
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("xiaohua")
	rVal.Method(0).Call(params) // xiaohua 完成了加法 1 + 1 = 2
}

func main() {
	cal := Cal{1, 1}
	ReflectFunc(&cal)
}
