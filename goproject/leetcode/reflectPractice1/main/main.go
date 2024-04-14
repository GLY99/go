package main

import (
	"fmt"
	"reflect"
)

func reflectFunc(t interface{}) {
	rVal := reflect.ValueOf(t)
	fmt.Printf("rVal Kind=%v, Type=%v\n", rVal.Kind(), rVal.Type()) // rVal Kind=float64, Type=float64

	i := rVal.Interface()
	num, ok := i.(float64)
	if ok {
		fmt.Printf("num=%v, type=%T", num, num) // num=1.2, type=float64
	}
}

func main() {
	var num float64 = 1.2
	reflectFunc(num)
}
