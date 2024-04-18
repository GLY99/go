package main

import (
	"fmt"
	"reflect"
)

// 使用反射遍历结构体字段，调用结构体的方法，获取结构体标签的值

// 定义Monster结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float64
	Sex   string
}

// 定义Monster结构体指针的方法
func (monster *Monster) Print() {
	fmt.Println("-----Print start-----")
	fmt.Println(monster)
	fmt.Println("-----Print end-----")
}

func (monster *Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (monster *Monster) Set(name string, age int, score float64, sex string) {
	monster.Name = name
	monster.Age = age
	monster.Score = score
	monster.Sex = sex
}

func TestFunc(t interface{}) {
	// 通过反射获取t的类型和值还有类别
	rType := reflect.TypeOf(t)
	rVal := reflect.ValueOf(t)
	kd := rVal.Kind() // struct
	if kd != reflect.Ptr {
		fmt.Println("t kind is not Ptr!")
		return
	}

	// 获取结构体字段数量
	num := rVal.Elem().NumField()             // 4
	fmt.Printf("struct has %d fields\n", num) //struct has 4 fields

	// 遍历结构体字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d, 值为=%v\n", i, rVal.Elem().Field(i))
		// 获取struct标签, 需要通过reflect.Type获取tag标签
		tagVal := rType.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Filed %d, tag=%v\n", i, tagVal)
		}
	}

	// 获取结构体指针方法数量, 方法和指针绑定所以不用Elem()
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod) // struct has 3 methods

	// 获取第二个方法，并且用Call调用, 方法排序按照函数名字的ASCLL码
	rVal.Method(1).Call(nil) // 这里调用Print

	// 调用第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := rVal.Method(0).Call(params)
	fmt.Printf("res=%d\n", res[0].Int()) // res=30

	// 调用第三个方法，修改struct属性
	params = make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf("huahua"))
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(100.9))
	params = append(params, reflect.ValueOf("nan"))
	rVal.Method(2).Call(params)
}

func main() {
	var monster Monster = Monster{
		Name:  "mimi",
		Age:   1900,
		Score: 20.9,
	}
	TestFunc(&monster)
	fmt.Println(monster) // {huahua 20 100.9 nan}
}
