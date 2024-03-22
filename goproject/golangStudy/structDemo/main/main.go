package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyStruct struct {
	Num   int
	Ptr   *int
	Arr   [5]int
	Slice []int
	Map   map[string]string
}

type Person struct {
	Name string
	Age  int
}

type Point struct {
	x int8
	y int8
}

type Rect1 struct {
	LeftUp, RightDown Point
}

type Rect2 struct {
	LeftUp, RightDown *Point
}

func main() {
	var cat1 Cat
	cat1.Name = "miaomiao"
	cat1.Age = 17
	fmt.Println(cat1)

	// 结构体中属性都会有一个默认值, 对于引用类型定义完后默认值是nil, 需要make空间来使用
	var myStruct1 MyStruct
	fmt.Println(myStruct1.Num)   // 0
	fmt.Println(myStruct1.Ptr)   // <nil>
	fmt.Println(myStruct1.Arr)   // [0 0 0 0 0]
	fmt.Println(myStruct1.Slice) // []
	fmt.Println(myStruct1.Map)   // map[]
	if myStruct1.Slice == nil {
		fmt.Println("nil") // 这里输出
	}
	if myStruct1.Map == nil {
		fmt.Println("nil") // 这里输出
	}
	myStruct1.Slice = make([]int, 5, 10)
	myStruct1.Map = make(map[string]string, 5)
	num := 1
	myStruct1.Ptr = &num
	fmt.Println(myStruct1.Num)   // 0
	fmt.Println(myStruct1.Ptr)   // 0xc00000a118
	fmt.Println(myStruct1.Arr)   // [0 0 0 0 0]
	fmt.Println(myStruct1.Slice) // [0 0 0 0 0]
	fmt.Println(myStruct1.Map)   // map[]
	if myStruct1.Slice == nil {
		fmt.Println("nil") // 这里不会输出
	}
	if myStruct1.Map == nil {
		fmt.Println("nil") // 这里不会输出
	}

	// 不同结构体变量互相不会影响
	var cat2 Cat
	cat2.Name = "miaomiao"
	cat2.Age = 17
	var cat3 Cat = Cat{Name: "mimi", Age: 18}
	fmt.Println(cat2) // {miaomiao 17}
	fmt.Println(cat3) // {mimi 18}

	// 创建结构体指针
	cat4 := new(Cat)
	cat4.Name = "haha" // go中针对结构体指针访问字段做了简化，这里等于（*cat4）.Name
	cat4.Age = 18
	fmt.Println(*cat4)
	var cat5 *Cat = &Cat{Name: "yiyi", Age: 1}
	fmt.Println(*cat5)

	// 结构体内存分布
	p1 := Person{Name: "xiaoming", Age: 18}
	p2 := p1 // 值类型会值拷贝
	p2.Name = "tom"
	fmt.Printf("p1的值%v,p2的值%v\n", p1, p2) // p1的值{xiaoming 18},p2的值{tom 18}

	p3 := Person{Name: "xiaoming", Age: 18}
	var p4 *Person = &p3 // 引用会地址传递
	p4.Name = "tom"
	fmt.Printf("p3的值%v,p4的值%v\n", p3, *p4)                    // p3的值{tom 18},p4的值{tom 18}
	fmt.Printf("p3的地址是%p, p4的值是%p, p4的地址是%v\n", &p3, p4, &p4) // p3的地址是0xc000008108, p4的值是0xc000008108, p4的地址是0xc00006c028

	// 结构体所有字段在内存中是连续的,结构体的地址就是第一个字段的地址
	rect1 := Rect1{Point{1, 1}, Point{2, 2}}
	// rect1.LeftUp.x的地址是0xc000016220,rect1.LeftUp.y的地址是0xc000016228,rect1.RightDown.x的地址是0xc000016230, rect1.RightDown.y的地址是0xc000016238
	fmt.Printf("rect1.LeftUp.x的地址是%p,rect1.LeftUp.y的地址是%p,"+
		"rect1.RightDown.x的地址是%p, rect1.RightDown.y的地址是%p\n",
		&rect1.LeftUp.x, &rect1.LeftUp.y, &rect1.RightDown.x, &rect1.RightDown.y)
	// rect1.LeftUp的地址是0xc000016220, rect1.RightDown的地址是0xc000016230
	fmt.Printf("rect1.LeftUp的地址是%p, rect1.RightDown的地址是%p\n", &rect1.LeftUp, &rect1.RightDown)

	// 结构体中指针的地址是连续的，但是指向的地址不一定是连续的
	rect2 := Rect2{&Point{1, 1}, &Point{2, 2}}
	// rect2.LeftUp的地址是0xc000026070, rect2.RightDown的地址是0xc000026078
	fmt.Printf("rect2.LeftUp的地址是%p, rect2.RightDown的地址是%p\n", &rect2.LeftUp, &rect2.RightDown)
	// rect2.LeftUp指向的地址是0xc00000a160, rect2.RightDown指向的地址是0xc00000a170
	fmt.Printf("rect2.LeftUp指向的地址是%p, rect2.RightDown指向的地址是%p\n", rect2.LeftUp, rect2.RightDown)

	// 结构体的序列化和反序列化&&结构体的tag
	cat6 := Cat{Name: "xiaomei", Age: 10}
	jsonBytesSlice, err := json.Marshal(cat6)
	if err != nil {
		fmt.Println("Marshal cat6 fail!")
	}
	fmt.Printf("jsonStr: %s", string(jsonBytesSlice)) // jsonStr: {"name":"xiaomei","age":10}
}
